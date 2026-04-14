// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package redeemcode

import (
	"context"

	"github.com/1348453525/user-redeem-code-gozero/api/internal/svc"
	"github.com/1348453525/user-redeem-code-gozero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UseRedeemCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUseRedeemCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UseRedeemCodeLogic {
	return &UseRedeemCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UseRedeemCodeLogic) UseRedeemCode(req *types.UseRedeemCodeRequest) error {
	// todo: add your logic here and delete this line

	return nil
}
