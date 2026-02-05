package startup

import (
	"context"
	"time"

	"github.com/carmasearch/carma-server/api/health"
	"github.com/carmasearch/carma-server/api/root"
	"github.com/carmasearch/carma-server/api/vehicle"
	vehicleDomain "github.com/carmasearch/carma-server/api/vehicle/domain"
	coreMW "github.com/carmasearch/carma-server/arch/middleware"
	"github.com/carmasearch/carma-server/arch/network"
	"github.com/carmasearch/carma-server/arch/redis"
	"github.com/carmasearch/carma-server/internal/config"
	"gorm.io/gorm"
)

type Module network.Module[module]

type module struct {
	Context        context.Context
	RootService    root.Service
	HealthService  health.Service
	VehicleService vehicleDomain.Service
	// List of service is here
	// UserService
	// AuthService and rest...
}

func (m *module) GetInstance() *module {
	return m
}

func (m *module) Controllers() []network.Controller {
	return []network.Controller{
		health.NewController(m.AuthenticationProvider(), m.AuthorizationProvider(), m.HealthService),
		root.NewController(m.AuthenticationProvider(), m.AuthorizationProvider(), m.RootService),
		// list of controller register will be here...
		vehicle.NewController(m.AuthenticationProvider(), m.AuthorizationProvider(), m.VehicleService),
	}
}

func (m *module) RootMiddlewares() []network.RootMiddleware {
	return []network.RootMiddleware{
		coreMW.NewErrorCatcher(), // NOTE this should be the first middleware to catche all errors
		coreMW.NewNotFound(),
	}
}

func (m *module) AuthenticationProvider() network.AuthenticationProvider {
	return &network.MockAuthenticationProvider{}
}

func (m *module) AuthorizationProvider() network.AuthorizationProvider {
	return &network.MockAuthorizationProvider{}
}

func NewModule(ctx context.Context, env *config.Config, db *gorm.DB, store redis.Store) Module {
	return &module{
		Context:       ctx,
		RootService:   root.NewService(),
		HealthService: health.NewService(env.Server.Environment, time.Now()),
		// userservice
		// auth service
		// list of serivce
		VehicleService: vehicle.NewService(vehicle.NewRepository(db)),
	}
}
