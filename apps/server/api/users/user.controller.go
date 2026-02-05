package users

import (
	"github.com/carmasearch/carma-server/arch/network"
	"github.com/gin-gonic/gin"
)

type userController struct {
	network.BaseController
	Service Service
}

func NewController(
	authProvider network.AuthenticationProvider,
	authorizeProvider network.AuthorizationProvider,
	service Service,
) network.Controller {
	return &userController{
		BaseController: network.NewBaseController("/api/v1", authProvider, authorizeProvider),
		Service:        service,
	}
}

func (c *userController) store(ctx *gin.Context) {}
func (c *userController) login(ctx *gin.Context) {}

func (c *userController) MountRoutes(group *gin.RouterGroup) {
	group.POST("register", c.store)
	group.POST("login", c.login)
}
