package handler

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/devloeprsabbir/go-elasticsearch/elastic"
	"github.com/devloeprsabbir/go-elasticsearch/models"
	"github.com/devloeprsabbir/go-elasticsearch/ranking"
	"github.com/devloeprsabbir/go-elasticsearch/repository"
	"github.com/elastic/go-elasticsearch/v9"
	"github.com/gin-gonic/gin"
)

// VehicleHandler handles all vehicle-related HTTP routes.
type VehicleHandler struct {
	repo *repository.VehicleRepository
	es   *elasticsearch.Client
}

func NewVehicleHandler(repo *repository.VehicleRepository, es *elasticsearch.Client) *VehicleHandler {
	return &VehicleHandler{repo: repo, es: es}
}

// SearchVehicles godoc
//
// POST /api/v1/vehicles/search
//
// Request body:
//
//	{
//	  "listing_url": "https://...",
//	  "page": 1,
//	  "page_size": 15
//	}
//
// The handler:
//  1. Looks up the reference vehicle by listing_url (tries Elasticsearch first, PG fallback)
//  2. Runs a weighted similarity search across 15+ fields in Elasticsearch
//  3. Applies the ranking engine (price prediction + multi-signal scoring)
//  4. Returns ranked results with score breakdown and price prediction metadata
func (h *VehicleHandler) SearchVehicles(c *gin.Context) {
	var req models.VehicleSearchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Defaults
	if req.Page < 1 {
		req.Page = 1
	}
	if req.PageSize < 1 || req.PageSize > 50 {
		req.PageSize = 15
	}

	ctx := c.Request.Context()

	// Step 1: Resolve the reference vehicle from the listing URL
	refVehicle, err := resolveReferenceVehicle(ctx, h, req.ListingURL)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "reference vehicle not found",
			"details": err.Error(),
		})
		return
	}

	// Step 2: Search for similar vehicles in Elasticsearch (returns scored results)
	scoredResults, total, err := elastic.SearchSimilarVehicles(ctx, refVehicle, req.Page, req.PageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "search failed: " + err.Error()})
		return
	}

	// Step 3: Convert to ranking.Vehicle — normalise ES score to [0,1] range
	var maxESScore float64
	for _, sv := range scoredResults {
		if sv.ESScore > maxESScore {
			maxESScore = sv.ESScore
		}
	}
	if maxESScore == 0 {
		maxESScore = 1 // avoid div-by-zero
	}

	rankVehicles := make([]ranking.Vehicle, 0, len(scoredResults))
	for _, sv := range scoredResults {
		rv := ranking.FromModel(sv.Vehicle, sv.ESScore/maxESScore)
		rankVehicles = append(rankVehicles, rv)
	}

	// Step 4: Build reference vehicle for ranking
	refRankVehicle := ranking.FromModel(*refVehicle, 1.0)

	// Step 5: Compute market average price from the candidate set
	marketAvg := ranking.MarketAvgPrice(rankVehicles)
	if marketAvg == 0 && refRankVehicle.Price > 0 {
		// Fall back to reference price as a proxy if ES returned no prices
		marketAvg = refRankVehicle.Price
	}

	// Step 6: Run the ranking engine
	cfg := ranking.DefaultRankingConfig()
	currentYear := time.Now().Year()
	rankResults := ranking.RankVehicles(refRankVehicle, rankVehicles, marketAvg, cfg, currentYear)

	// Step 7: Build the enriched response
	// Build a quick lookup: ranking.Vehicle.ID → models.ScoredVehicle
	vehicleByID := make(map[string]models.Vehicle, len(scoredResults))
	for _, sv := range scoredResults {
		vehicleByID[sv.Vehicle.UniqueID] = sv.Vehicle
	}

	enrichedResults := make([]models.RankedVehicleResult, 0, len(rankResults))
	for i, rr := range rankResults {
		mv, ok := vehicleByID[rr.ID]
		if !ok {
			continue // safety guard
		}

		predictedPrice := rr.Metadata["predicted_price"]
		listingPrice := rr.Price
		priceDelta := listingPrice - predictedPrice

		enrichedResults = append(enrichedResults, models.RankedVehicleResult{
			Rank:       i + 1,
			FinalScore: rr.FinalScore,
			ScoreBreakdown: models.ScoreBreakdown{
				PriceScore:      rr.Metadata["price_score"],
				MileageScore:    rr.Metadata["mileage_score"],
				YearScore:       rr.Metadata["year_score"],
				SimilarityScore: rr.Metadata["similarity_score"],
				PopularityScore: rr.Metadata["popularity_score"],
			},
			PricePrediction: models.PricePrediction{
				PredictedFairPrice: predictedPrice,
				ListingPrice:       listingPrice,
				PriceDelta:         priceDelta,
				DealQuality:        ranking.DealQuality(listingPrice, predictedPrice),
			},
			Vehicle: mv,
		})
	}

	c.JSON(http.StatusOK, models.EnrichedSearchResponse{
		Total:          total,
		Page:           req.Page,
		PageSize:       req.PageSize,
		MarketAvgPrice: marketAvg,
		QueryVehicle:   refVehicle,
		Results:        enrichedResults,
	})
}

// resolveReferenceVehicle tries to get the reference vehicle.
// It first tries Elasticsearch (fast), then falls back to PostgreSQL (complete).
func resolveReferenceVehicle(ctx context.Context, h *VehicleHandler, listingURL string) (*models.Vehicle, error) {
	// Try Elasticsearch first
	v, err := elastic.GetVehicleByListingURL(ctx, listingURL)
	if err == nil && v != nil {
		return v, nil
	}
	log.Printf("listing_url not found in Elasticsearch (%v), falling back to PostgreSQL", err)

	// Fallback to PostgreSQL
	return h.repo.GetByListingURL(ctx, listingURL)
}

// RunBackgroundSync starts the indexing process asynchronously.
// It is safe to call on application startup or via an API endpoint.
func (h *VehicleHandler) RunBackgroundSync() {
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Minute)
		defer cancel()

		count, err := h.repo.CountVehicles(ctx)
		if err != nil {
			log.Printf("BackgroundSync: count error: %v", err)
			return
		}
		log.Printf("BackgroundSync: starting indexing of %d vehicles", count)

		err = h.repo.FetchAllForIndexing(ctx, 500, func(batch []models.Vehicle) error {
			return elastic.BulkIndexVehicles(batch, 500)
		})
		if err != nil {
			log.Printf("BackgroundSync: error during indexing: %v", err)
		} else {
			log.Printf("BackgroundSync: completed successfully (%d vehicles)", count)
		}
	}()
}

// TriggerIndexing godoc
// POST /api/v1/vehicles/index
//
// Triggers a background re-indexing of all vehicle records from PostgreSQL → Elasticsearch.
// Returns immediately with 202 Accepted; indexing runs asynchronously.
// Progress is logged to the server log.
func (h *VehicleHandler) TriggerIndexing(c *gin.Context) {
	h.RunBackgroundSync()

	c.JSON(http.StatusAccepted, models.IndexTriggerResponse{
		Message:   "Indexing started in background. Check server logs for progress.",
		Scheduled: true,
	})
}

// GetVehicleStatus godoc
// GET /api/v1/vehicles/status
//
// Returns current count of indexed vehicles in Elasticsearch vs PostgreSQL.
func (h *VehicleHandler) GetVehicleStatus(c *gin.Context) {
	ctx := c.Request.Context()

	pgCount, err := h.repo.CountVehicles(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "pg count failed: " + err.Error()})
		return
	}

	var esCount int64
	esCountRes, esErr := elastic.EsClient.Count(
		elastic.EsClient.Count.WithIndex("vehicles"),
	)
	if esErr == nil && !esCountRes.IsError() {
		body, readErr := io.ReadAll(esCountRes.Body)
		esCountRes.Body.Close()
		if readErr == nil {
			var cr map[string]interface{}
			if json.Unmarshal(body, &cr) == nil {
				if v, ok := cr["count"].(float64); ok {
					esCount = int64(v)
				}
			}
		}
	} else if esCountRes != nil {
		esCountRes.Body.Close()
	}

	var indexedPct float64
	if pgCount > 0 {
		indexedPct = float64(esCount) / float64(pgCount) * 100
	}

	c.JSON(http.StatusOK, gin.H{
		"postgres_count":  pgCount,
		"elastic_count":   esCount,
		"indexed_percent": indexedPct,
	})
}
