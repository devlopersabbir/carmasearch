package vehicle

import (
	"net/http"
	"strconv"

	"github.com/carmasearch/carma-server/api/vehicle/core"
	"github.com/carmasearch/carma-server/api/vehicle/domain"
	"github.com/carmasearch/carma-server/arch/network"
	"github.com/gin-gonic/gin"
)

type vehicleController struct {
	network.BaseController
	service domain.Service
}

func NewController(
	authProvider network.AuthenticationProvider,
	authorizeProvider network.AuthorizationProvider,
	service domain.Service,
) network.Controller {
	return &vehicleController{
		BaseController: network.NewBaseController("/api/v1", authProvider, authorizeProvider),
		service:        service,
	}
}

func (c *vehicleController) create(ctx *gin.Context) {
	var vehicle core.Vehicle
	if err := ctx.ShouldBindJSON(&vehicle); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.CreateVehicle(&vehicle); err != nil {
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

//	func (c *vehicleController) MountRoutes(group *gin.RouterGroup) {
//		// vehicleGroup := group.Group("vehicles")
//		// {
//		// 	vehicleGroup.POST("", c.create)
//		// 	vehicleGroup.GET("/:id", c.get)
//		// 	vehicleGroup.GET("", c.list)
//		// }
//	}
func (c *vehicleController) MountRoutes(group *gin.RouterGroup) {}
