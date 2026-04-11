package logic

import (
	"context"

	"github.com/1348453525/user-redeem-code-gozero/redeem-code-rpc/internal/svc"
	"github.com/1348453525/user-redeem-code-gozero/redeem-code-rpc/redeemcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRedeemCodeBatchListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRedeemCodeBatchListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRedeemCodeBatchListLogic {
	return &GetRedeemCodeBatchListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取兑换码批次列表
func (l *GetRedeemCodeBatchListLogic) GetRedeemCodeBatchList(in *redeemcode.GetListRequest) (*redeemcode.GetRedeemCodeBatchListResponse, error) {
	// todo: add your logic here and delete this line

	return &redeemcode.GetRedeemCodeBatchListResponse{}, nil
}
