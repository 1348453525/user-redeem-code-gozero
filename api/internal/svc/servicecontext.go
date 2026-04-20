// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package svc

import (
	"github.com/1348453525/user-redeem-code-gozero/api/internal/config"
	"github.com/1348453525/user-redeem-code-gozero/redeem-code-rpc/redeemcodeclient"
	userclient "github.com/1348453525/user-redeem-code-gozero/user-rpc/client/user"
	"github.com/zeromicro/go-zero/zrpc"
	_ "github.com/zeromicro/zero-contrib/zrpc/registry/consul"
)

type ServiceContext struct {
	Config        config.Config
	UserRpc       userclient.User
	RedeemCodeRpc redeemcodeclient.RedeemCode
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UserRpc:       userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		RedeemCodeRpc: redeemcodeclient.NewRedeemCode(zrpc.MustNewClient(c.RedeemCodeRpc)),
	}
}
