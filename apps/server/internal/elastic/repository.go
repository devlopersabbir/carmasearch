package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"strconv"

	"github.com/carmasearch/carma-server/api/vehicle/core"
	es "github.com/carmasearch/carma-server/internal/database"
	esCore "github.com/carmasearch/carma-server/internal/elastic/core"
)

func Search2(c context.Context, body *core.Vehicle) ([]uint64, int64, error) {
	return nil, 0, nil
}

func Search(ctx context.Context, req *esCore.VehicleSearchQuery) ([]uint64, int64, error) {
	log.Println("search request::::::::::::::::", &req)
	if es.ESClient == nil {
		return nil, 0, nil
	}

	mustArr := []interface{}{}
	filterArr := []interface{}{}
	shouldArr := []interface{}{}

	// =========================
	// Full Text Search
	// =========================
	if req.Query != "" {
		mustArr = append(mustArr, map[string]interface{}{
			"multi_match": map[string]interface{}{
				"query":  req.Query,
				"fields": []string{"title^3", "make^2", "model^2", "seller_name"},
				"type":   "best_fields",
			},
		})
	}

	// =========================
	// Exact Match Filters
	// =========================
	if req.VehicleID != nil {
		filterArr = append(filterArr, map[string]interface{}{
			"term": map[string]interface{}{
				"vehicle_id.keyword": *req.VehicleID,
			},
		})
	}

	if req.Make != nil {
		filterArr = append(filterArr, map[string]interface{}{
			"term": map[string]interface{}{
				"make.keyword": *req.Make,
			},
		})
	}

	if req.Model != nil {
		filterArr = append(filterArr, map[string]interface{}{
			"term": map[string]interface{}{
				"model.keyword": *req.Model,
			},
		})
	}

	if req.SellerName != nil {
		filterArr = append(filterArr, map[string]interface{}{
			"term": map[string]interface{}{
				"seller_name.keyword": *req.SellerName,
			},
		})
	}

	// =========================
	// Multi Select Fields
	// =========================
	if len(req.FuelType) > 0 {
		filterArr = append(filterArr, map[string]interface{}{
			"terms": map[string]interface{}{
				"fuel_type.keyword": req.FuelType,
			},
		})
	}

	if len(req.Transmission) > 0 {
		filterArr = append(filterArr, map[string]interface{}{
			"terms": map[string]interface{}{
				"transmission.keyword": req.Transmission,
			},
		})
	}

	if len(req.ExteriorColors) > 0 {
		filterArr = append(filterArr, map[string]interface{}{
			"terms": map[string]interface{}{
				"exterior_colors.keyword": req.ExteriorColors,
			},
		})
	}

	if len(req.InteriorColors) > 0 {
		filterArr = append(filterArr, map[string]interface{}{
			"terms": map[string]interface{}{
				"interior_colors.keyword": req.InteriorColors,
			},
		})
	}

	if len(req.InteriorMaterials) > 0 {
		filterArr = append(filterArr, map[string]interface{}{
			"terms": map[string]interface{}{
				"interior_materials.keyword": req.InteriorMaterials,
			},
		})
	}

	// =========================
	// Range Filters
	// =========================

	addRangeFilter := func(field string, from interface{}, to interface{}) {
		rangeQuery := map[string]interface{}{}

		if from != nil {
			rangeQuery["gte"] = from
		}
		if to != nil {
			rangeQuery["lte"] = to
		}

		if len(rangeQuery) > 0 {
			filterArr = append(filterArr, map[string]interface{}{
				"range": map[string]interface{}{
					field: rangeQuery,
				},
			})
		}
	}

	addRangeFilter("price", req.PriceFrom, req.PriceTo)
	addRangeFilter("registration_year", req.RegistrationFrom, req.RegistrationTo)
	addRangeFilter("mileage", req.MileageFrom, req.MileageTo)
	addRangeFilter("co2_emission", req.CO2EmissionFrom, req.CO2EmissionTo)

	// =========================
	// Boolean Feature Filters
	// =========================
	addBoolFilter := func(field string, value *bool) {
		if value != nil && *value {
			filterArr = append(filterArr, map[string]interface{}{
				"term": map[string]interface{}{
					field: true,
				},
			})
		}
	}

	addBoolFilter("abs", req.ABS)
	addBoolFilter("esp", req.ESP)
	addBoolFilter("traction_control", req.TractionControl)
	addBoolFilter("emergency_brake_assist", req.EmergencyBrakeAssist)
	addBoolFilter("blind_spot_assist", req.BlindSpotAssist)
	addBoolFilter("lane_assist", req.LaneAssist)
	addBoolFilter("traffic_sign_recognition", req.TrafficSignRecognition)
	addBoolFilter("isofix", req.ISOFIX)

	addBoolFilter("heated_steering_wheel", req.HeatedSteeringWheel)
	addBoolFilter("start_stop_system", req.StartStopSystem)
	addBoolFilter("heated_seats", req.HeatedSeats)
	addBoolFilter("electric_seats", req.ElectricSeats)
	addBoolFilter("sport_seats", req.SportSeats)

	addBoolFilter("fog_lights", req.FogLights)
	addBoolFilter("adaptive_headlights", req.AdaptiveHeadlights)
	addBoolFilter("rain_sensor", req.RainSensor)

	addBoolFilter("radio", req.Radio)
	addBoolFilter("navigation_system", req.NavigationSystem)
	addBoolFilter("voice_control", req.VoiceControl)
	addBoolFilter("bluetooth", req.Bluetooth)
	addBoolFilter("usb", req.USB)
	addBoolFilter("apple_carplay", req.AppleCarPlay)
	addBoolFilter("android_auto", req.AndroidAuto)

	// =========================
	// Final Query
	// =========================

	var query interface{}

	if len(mustArr) == 0 && len(filterArr) == 0 {
		query = map[string]interface{}{
			"match_all": map[string]interface{}{},
		}
	} else {
		query = map[string]interface{}{
			"bool": map[string]interface{}{
				"must":   mustArr,
				"filter": filterArr,
				"should": shouldArr,
			},
		}
	}
	log.Println("query::::::::::::::::", query)
	sortArr := []interface{}{
		map[string]interface{}{
			req.SortBy: map[string]interface{}{
				"order": req.SortOrder,
			},
		},
	}

	queryMap := map[string]interface{}{
		"query": query,
		"from":  (req.Page - 1) * req.PageSize,
		"size":  req.PageSize,
		"sort":  sortArr,
	}

	queryBytes, _ := json.Marshal(queryMap)

	res, err := es.ESClient.Search(
		es.ESClient.Search.WithContext(ctx),
		es.ESClient.Search.WithIndex(es.ESIndexName),
		es.ESClient.Search.WithBody(bytes.NewReader(queryBytes)),
		es.ESClient.Search.WithTrackTotalHits(true),
	)

	if err != nil {
		return nil, 0, err
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Println("search error:", res.String())
		return nil, 0, nil
	}

	var r map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, 0, err
	}

	hitsMap, ok := r["hits"].(map[string]interface{})
	if !ok {
		return nil, 0, nil
	}

	hits, ok := hitsMap["hits"].([]interface{})
	if !ok {
		return nil, 0, nil
	}

	var total int64
	if totalMap, ok := hitsMap["total"].(map[string]interface{}); ok {
		if val, ok := totalMap["value"].(float64); ok {
			total = int64(val)
		}
	}

	var ids []uint64
	for _, hit := range hits {
		if hitMap, ok := hit.(map[string]interface{}); ok {
			if idStr, ok := hitMap["_id"].(string); ok {
				idUint, _ := strconv.ParseUint(idStr, 10, 64)
				ids = append(ids, idUint)
			}
		}
	}

	return ids, total, nil
}
func IndexVehicle(s *core.Vehicle) {
	jsonBody, err := json.Marshal(s)
	if err != nil {
		log.Println("marshal error:", err)
		return
	}

	res, err := es.ESClient.Index(
		es.ESIndexName,
		bytes.NewReader(jsonBody),
		es.ESClient.Index.WithDocumentID(strconv.Itoa(int(s.ID))),
		es.ESClient.Index.WithRefresh("true"),
	)
	if err != nil {
		log.Println("index request failed:", err)
		return
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Println("indexing error:", res.String())
		return
	}

	log.Println("indexed todo successfully:", s.ID)
}
