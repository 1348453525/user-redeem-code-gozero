package svc

import "github.com/1348453525/user-redeem-code-gozero/redeem-code-rpc/internal/config"

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
