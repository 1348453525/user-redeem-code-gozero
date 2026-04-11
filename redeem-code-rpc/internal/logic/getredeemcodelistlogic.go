package logic

import (
	"context"

	"github.com/1348453525/user-redeem-code-gozero/redeem-code-rpc/internal/svc"
	"github.com/1348453525/user-redeem-code-gozero/redeem-code-rpc/redeemcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRedeemCodeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRedeemCodeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRedeemCodeListLogic {
	return &GetRedeemCodeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取兑换码列表
func (l *GetRedeemCodeListLogic) GetRedeemCodeList(in *redeemcode.GetListRequest) (*redeemcode.GetRedeemCodeListResponse, error) {
	// todo: add your logic here and delete this line

	return &redeemcode.GetRedeemCodeListResponse{}, nil
}
