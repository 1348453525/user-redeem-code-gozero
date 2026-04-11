package logic

import (
	"context"

	"github.com/1348453525/user-redeem-code-gozero/redeem-code-rpc/internal/svc"
	"github.com/1348453525/user-redeem-code-gozero/redeem-code-rpc/redeemcode"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRedeemCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRedeemCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRedeemCodeLogic {
	return &UpdateRedeemCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新兑换码
func (l *UpdateRedeemCodeLogic) UpdateRedeemCode(in *redeemcode.UpdateRedeemCodeRequest) (*emptypb.Empty, error) {
	// todo: add your logic here and delete this line

	return &emptypb.Empty{}, nil
}
