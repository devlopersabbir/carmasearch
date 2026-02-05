package users

import "github.com/carmasearch/carma-server/arch/network"

type Service interface {
	store() (string, error)
	login() (string, error)
}

type service struct {
	network.BaseService
}

func NewService() Service {
	return &service{
		BaseService: network.NewBaseService(),
	}
}

func (s *service) store() (string, error) {
	return "", nil
}

func (s *service) login() (string, error) {
	return "", nil
}
