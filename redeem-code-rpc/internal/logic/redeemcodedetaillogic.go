package logic

import (
	"context"

	"github.com/1348453525/user-redeem-code-gozero/redeem-code-rpc/internal/svc"
	"github.com/1348453525/user-redeem-code-gozero/redeem-code-rpc/redeemcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type RedeemCodeDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRedeemCodeDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RedeemCodeDetailLogic {
	return &RedeemCodeDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取兑换码详情
func (l *RedeemCodeDetailLogic) RedeemCodeDetail(in *redeemcode.IDRequest) (*redeemcode.RedeemCodeResponse, error) {
	// todo: add your logic here and delete this line

	return &redeemcode.RedeemCodeResponse{}, nil
}
