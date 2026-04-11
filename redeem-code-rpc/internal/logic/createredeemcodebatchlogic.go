package logic

import (
	"context"

	"github.com/1348453525/user-redeem-code-gozero/redeem-code-rpc/internal/svc"
	"github.com/1348453525/user-redeem-code-gozero/redeem-code-rpc/redeemcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRedeemCodeBatchLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateRedeemCodeBatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRedeemCodeBatchLogic {
	return &CreateRedeemCodeBatchLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建兑换码批次
func (l *CreateRedeemCodeBatchLogic) CreateRedeemCodeBatch(in *redeemcode.CreateRedeemCodeBatchRequest) (*redeemcode.RedeemCodeBatchResponse, error) {
	// todo: add your logic here and delete this line

	return &redeemcode.RedeemCodeBatchResponse{}, nil
}
