// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package redeemcodebatch

import (
	"context"

	"github.com/1348453525/user-redeem-code-gozero/api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteRedeemCodeBatchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteRedeemCodeBatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRedeemCodeBatchLogic {
	return &DeleteRedeemCodeBatchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteRedeemCodeBatchLogic) DeleteRedeemCodeBatch() error {
	// todo: add your logic here and delete this line

	return nil
}
