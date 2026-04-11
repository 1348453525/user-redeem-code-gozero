package logic

import (
	"context"

	"github.com/1348453525/user-redeem-code-gozero/redeem-code-rpc/internal/svc"
	"github.com/1348453525/user-redeem-code-gozero/redeem-code-rpc/redeemcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type RedeemCodeBatchDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRedeemCodeBatchDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RedeemCodeBatchDetailLogic {
	return &RedeemCodeBatchDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取兑换码批次详情
func (l *RedeemCodeBatchDetailLogic) RedeemCodeBatchDetail(in *redeemcode.IDRequest) (*redeemcode.RedeemCodeBatchResponse, error) {
	// todo: add your logic here and delete this line

	return &redeemcode.RedeemCodeBatchResponse{}, nil
}
