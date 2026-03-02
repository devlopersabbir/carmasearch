package vehicle

import (
	"net/http"
	"strconv"

	"github.com/carmasearch/carma-server/api/vehicle/core"
	"github.com/carmasearch/carma-server/api/vehicle/domain"
	"github.com/carmasearch/carma-server/arch/network"
	esCore "github.com/carmasearch/carma-server/internal/elastic/core"
	esDomain "github.com/carmasearch/carma-server/internal/elastic/domain"
	"github.com/carmasearch/carma-server/internal/scraper"
	"github.com/gin-gonic/gin"
)

type vehicleController struct {
	network.BaseController
	service   domain.Service
	esService esDomain.VehicleCompareService
}

func NewController(
	authProvider network.AuthenticationProvider,
	authorizeProvider network.AuthorizationProvider,
	service domain.Service,
	esService esDomain.VehicleCompareService,
) network.Controller {
	return &vehicleController{
		BaseController: network.NewBaseController("/api", authProvider, authorizeProvider),
		service:        service,
		esService:      esService,
	}
}

func (c *vehicleController) create(ctx *gin.Context) {
	var vehicle core.Vehicle
	if err := ctx.ShouldBindJSON(&vehicle); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.CreateVehicle(ctx.Request.Context(), &vehicle); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, vehicle)
}

func (c *vehicleController) get(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	vehicle, err := c.service.GetVehicle(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Vehicle not found"})
		return
	}

	ctx.JSON(http.StatusOK, vehicle)
}

func (c *vehicleController) list(ctx *gin.Context) {
	limitStr := ctx.DefaultQuery("limit", "10")
	offsetStr := ctx.DefaultQuery("offset", "0")

	limit, _ := strconv.Atoi(limitStr)
	offset, _ := strconv.Atoi(offsetStr)

	vehicles, count, err := c.service.ListVehicles(limit, offset)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  vehicles,
		"count": count,
	})
}

func (cn *vehicleController) search(c *gin.Context) {
	var body esCore.CompareRequest
	var query esCore.CompareRequestQuery

	// Bind JSON body
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Bind query params
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Normalize pagination
	if query.Page < 1 {
		query.Page = 1
	}
	if query.PageSize < 1 || query.PageSize > 100 {
		query.PageSize = 20
	}

	// 1. Call Scraper using body.Url
	if err := scraper.Scrape(body.Url); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 2. Build VehicleSearchAndCompare
	searchAndQuery := buildVehicleSearchAndCompare(&body, &query)

	// 3. Call service
	total, vehicles, err := cn.service.SearchAndCompare(
		c.Request.Context(),
		searchAndQuery,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result := make([]core.Vehicle, 0, len(vehicles))
	for _, v := range vehicles {
		if v != nil {
			result = append(result, *v)
		}
	}

	c.JSON(http.StatusOK, esCore.VehicleSearchQueryResponse{
		Total:    uint64(total),
		Page:     query.Page,
		Pagesize: query.PageSize,
		Vehicles: result,
	})
}

func (cn *vehicleController) search2(c *gin.Context) {

}

func (cn *vehicleController) MountRoutes(group *gin.RouterGroup) {
	v1 := group.Group("/v1/vehicles")
	{
		v1.POST("", cn.create)
		v1.GET("/:id", cn.get)
		v1.GET("", cn.list)
		v1.POST("/search", cn.search)
	}
	v2 := group.Group("/v2/vehicles")
	{
		v2.GET("", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Hello World from v2",
			})
		})
		v2.POST("/search", cn.search2)
	}
}
