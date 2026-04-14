// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package svc

import (
	"github.com/1348453525/user-redeem-code-gozero/api/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
