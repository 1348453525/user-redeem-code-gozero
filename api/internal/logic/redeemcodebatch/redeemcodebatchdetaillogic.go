// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package redeemcodebatch

import (
	"context"

	"github.com/1348453525/user-redeem-code-gozero/api/internal/svc"
	"github.com/1348453525/user-redeem-code-gozero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RedeemCodeBatchDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRedeemCodeBatchDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RedeemCodeBatchDetailLogic {
	return &RedeemCodeBatchDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RedeemCodeBatchDetailLogic) RedeemCodeBatchDetail() (resp *types.RedeemCodeBatchResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
