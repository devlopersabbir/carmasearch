package vehicle

import "github.com/carmasearch/carma-server/arch/network"

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
