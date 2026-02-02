package root

import (
	"github.com/carmasearch/carma-server/arch/network"
	"github.com/gin-gonic/gin"
)

type controller struct {
	network.BaseController
	Service Service
}

func NewController(
	authProvider network.AuthenticationProvider,
	authorizeProvider network.AuthorizationProvider,
	service Service,
) network.Controller {
	return &controller{
		BaseController: network.NewBaseController("", authProvider, authorizeProvider),
		Service:        service,
	}
}

func (c *controller) MountRoutes(group *gin.RouterGroup) {
	group.GET("", c.rootHandler)
}
func (c *controller) rootHandler(ctx *gin.Context) {
	c.Send(ctx).SuccessMsgResponse("Welcome to Carmasearch API Server")
}
