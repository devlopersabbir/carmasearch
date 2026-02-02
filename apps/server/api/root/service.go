package root

import "github.com/carmasearch/carma-server/arch/network"

type Service interface {
	RootHandler() (string, error)
}

type service struct {
	network.BaseService
}

func NewService() Service {
	return &service{
		BaseService: network.NewBaseService(),
	}
}
func (s *service) RootHandler() (string, error) {
	return "Hello World", nil
}
