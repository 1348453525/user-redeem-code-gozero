package logic

import (
	"context"

	"github.com/1348453525/user-redeem-code-gozero/model"
	"github.com/1348453525/user-redeem-code-gozero/pkg/errorx"
	"github.com/1348453525/user-redeem-code-gozero/pkg/helper"
	"github.com/1348453525/user-redeem-code-gozero/redeem-code-rpc/internal/svc"
	proto "github.com/1348453525/user-redeem-code-gozero/redeem-code-rpc/redeemcode"
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
func (l *UpdateRedeemCodeLogic) UpdateRedeemCode(r *proto.UpdateRedeemCodeRequest) (*emptypb.Empty, error) {
	redeemCodeModel := model.RedeemCode{
		ID: r.Id,
	}
	if r.Title != "" {
		redeemCodeModel.Title = r.Title
	}
	if r.ExpirationAt != "" {
		expirationAt, err := helper.ParseDatetime(r.ExpirationAt)
		if err != nil {
			l.Logger.Errorw("解析过期时间失败", logx.Field("err", err), logx.Field("ExpirationAt", r.ExpirationAt))
			return &emptypb.Empty{}, errorx.ToGrpcError(errorx.ErrParam)
		}
		redeemCodeModel.ExpirationAt = *expirationAt
	}
	if r.IsDel != 0 {
		redeemCodeModel.IsDel = r.IsDel
	}
	if err := l.svcCtx.DB.Model(&model.RedeemCode{}).Where("id=?", r.Id).Updates(&redeemCodeModel).Error; err != nil {
		l.Logger.Errorw("更新兑换码失败", logx.Field("err", err), logx.Field("id", r.Id))
		return &emptypb.Empty{}, errorx.ToGrpcError(errorx.ErrInternal)
	}
	return &emptypb.Empty{}, nil
}
