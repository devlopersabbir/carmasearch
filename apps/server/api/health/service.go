package health

import (
	"fmt"
	"runtime"
	"time"

	"github.com/carmasearch/carma-server/api/health/dto"
	"github.com/carmasearch/carma-server/arch/network"
)

type Service interface {
	GetApplicationHealth() (*dto.InfoHealth, error)
}
type service struct {
	network.BaseService
	env       string
	startTime time.Time
}

func NewService(env string, startTime time.Time) Service {
	return &service{
		BaseService: network.NewBaseService(),
		env:         env,
		startTime:   startTime,
	}
}

func (s *service) GetApplicationHealth() (*dto.InfoHealth, error) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	return &dto.InfoHealth{
		Application: "carma-server",
		Version:     "1.0.0",
		Environment: s.env,
		Uptime:      time.Since(s.startTime).String(),
		Memory:      fmt.Sprintf("%d MB", m.Alloc/1024/1024),
		CPU:         fmt.Sprintf("%d", runtime.NumCPU()),
	}, nil
}
