package vehicle

import (
	"log"
	"net/http"
	"strconv"

	"github.com/carmasearch/carma-server/api/vehicle/core"
	"github.com/carmasearch/carma-server/api/vehicle/domain"
	"github.com/carmasearch/carma-server/arch/network"
	esCore "github.com/carmasearch/carma-server/internal/elastic/core"
	esDomain "github.com/carmasearch/carma-server/internal/elastic/domain"
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
		BaseController: network.NewBaseController("/api/v1", authProvider, authorizeProvider),
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

	if err := c.service.CreateVehicle(ctx, &vehicle); err != nil {
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

func (c *vehicleController) search(ctx *gin.Context) {
	// get all query params without defined key
	// exmaple /search?version=1&color=red&price=100000
	// returns map[string][]string
	rowQueries := ctx.Request.URL.Query()
	filters := make(map[string]interface{})

	for key, values := range rowQueries {
		if len(values) == 1 {
			filters[key] = values[0]
		} else {
			// support multi-value filters like ?color=red&color=blue
			filters[key] = values
		}
	}
	vehicles, err := c.service.SearchVehicles(filters)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, vehicles)
}

func (c *vehicleController) compare(ctx *gin.Context) {
	var input esCore.VehicleSearchQuery

	// Scraper sends full vehicle JSON here
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println("there are no error")
	vehicles, err := c.esService.CompareVehicle(&input)
	log.Println("vehicless......")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":  vehicles,
		"count": len(vehicles),
	})
}

func (c *vehicleController) MountRoutes(group *gin.RouterGroup) {
	vehicleGroup := group.Group("vehicles")
	{
		vehicleGroup.POST("", c.create)
		vehicleGroup.GET("/:id", c.get)
		vehicleGroup.GET("", c.list)
		vehicleGroup.GET("/search", c.search)
		vehicleGroup.POST("/compare", c.compare)
	}
}
