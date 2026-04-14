// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package redeemcodebatch

import (
	"context"

	"github.com/1348453525/user-redeem-code-gozero/api/internal/svc"
	"github.com/1348453525/user-redeem-code-gozero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRedeemCodeBatchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateRedeemCodeBatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRedeemCodeBatchLogic {
	return &CreateRedeemCodeBatchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateRedeemCodeBatchLogic) CreateRedeemCodeBatch(req *types.CreateRedeemCodeBatchRequest) (resp *types.RedeemCodeBatchResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
