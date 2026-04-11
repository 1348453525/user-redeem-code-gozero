package logic

import (
	"context"

	"github.com/1348453525/user-redeem-code-gozero/redeem-code-rpc/internal/svc"
	"github.com/1348453525/user-redeem-code-gozero/redeem-code-rpc/redeemcode"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRedeemCodeBatchLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRedeemCodeBatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRedeemCodeBatchLogic {
	return &UpdateRedeemCodeBatchLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新兑换码批次
func (l *UpdateRedeemCodeBatchLogic) UpdateRedeemCodeBatch(in *redeemcode.UpdateRedeemCodeBatchRequest) (*emptypb.Empty, error) {
	// todo: add your logic here and delete this line

	return &emptypb.Empty{}, nil
}
