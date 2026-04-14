package logic

import (
	"context"

	"github.com/1348453525/user-redeem-code-gozero/model"
	"github.com/1348453525/user-redeem-code-gozero/pkg/errorx"
	"github.com/1348453525/user-redeem-code-gozero/redeem-code-rpc/internal/svc"
	proto "github.com/1348453525/user-redeem-code-gozero/redeem-code-rpc/redeemcode"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteRedeemCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteRedeemCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRedeemCodeLogic {
	return &DeleteRedeemCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除兑换码
func (l *DeleteRedeemCodeLogic) DeleteRedeemCode(r *proto.IDRequest) (*emptypb.Empty, error) {
	if err := l.svcCtx.DB.Model(&model.RedeemCode{}).Where("id=?", r.Id).Update("is_del", 1).Error; err != nil {
		l.Logger.Errorw("删除兑换码失败", logx.Field("err", err), logx.Field("id", r.Id))
		return &emptypb.Empty{}, errorx.ToGrpcError(errorx.ErrInternal)
	}
	return &emptypb.Empty{}, nil
}
