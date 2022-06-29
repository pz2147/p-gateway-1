package svc

import (
	"github.com/pz2147/p-gateway-1/internal/config"
)

type ServiceContext struct {
	Config config.Config
	GwHttpX *GatewayHttpX
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		GwHttpX: NewGatewayHttpX(c),
	}
}
