package logic

import (
	"context"

	"github.com/1348453525/user-redeem-code-gozero/redeem-code-rpc/internal/svc"
	"github.com/1348453525/user-redeem-code-gozero/redeem-code-rpc/redeemcode"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteRedeemCodeBatchLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteRedeemCodeBatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRedeemCodeBatchLogic {
	return &DeleteRedeemCodeBatchLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除兑换码批次
func (l *DeleteRedeemCodeBatchLogic) DeleteRedeemCodeBatch(in *redeemcode.IDRequest) (*emptypb.Empty, error) {
	// todo: add your logic here and delete this line

	return &emptypb.Empty{}, nil
}
