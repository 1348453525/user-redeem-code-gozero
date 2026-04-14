// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package redeemcode

import (
	"context"

	"github.com/1348453525/user-redeem-code-gozero/api/internal/svc"
	"github.com/1348453525/user-redeem-code-gozero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RedeemCodeDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRedeemCodeDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RedeemCodeDetailLogic {
	return &RedeemCodeDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RedeemCodeDetailLogic) RedeemCodeDetail() (resp *types.RedeemCodeResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
