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
func (l *UpdateRedeemCodeBatchLogic) UpdateRedeemCodeBatch(r *proto.UpdateRedeemCodeBatchRequest) (*emptypb.Empty, error) {
	// 参数验证
	if r.Id <= 0 || r.Title == "" {
		return &emptypb.Empty{}, errorx.ToGrpcError(errorx.ErrParam)
	}

	// 解析日期
	startedAt, err := helper.ParseDatetime(r.StartedAt)
	if err != nil {
		l.Logger.Errorw("解析开始时间失败", logx.Field("err", err), logx.Field("startedAt", r.StartedAt))
		return &emptypb.Empty{}, errorx.ToGrpcError(errorx.ErrParam)
	}
	endedAt, err := helper.ParseDatetime(r.EndedAt)
	if err != nil {
		l.Logger.Errorw("解析结束时间失败", logx.Field("err", err), logx.Field("endedAt", r.EndedAt))
		return &emptypb.Empty{}, errorx.ToGrpcError(errorx.ErrParam)
	}

	// 验证时间范围
	if endedAt.Before(*startedAt) {
		l.Logger.Errorw("结束时间早于开始时间", logx.Field("startedAt", r.StartedAt), logx.Field("endedAt", r.EndedAt))
		return &emptypb.Empty{}, errorx.ToGrpcError(errorx.ErrParam)
	}

	// 验证批次是否存在
	exist, err := model.GetRedeemCodeBatchByID(l.svcCtx.DB, r.Id)
	if err != nil {
		l.Logger.Errorw("获取兑换码批次失败", logx.Field("err", err), logx.Field("id", r.Id))
		return &emptypb.Empty{}, errorx.ToGrpcError(errorx.ErrInternal)
	}

	updateRedeemCodeBatch := model.RedeemCodeBatch{
		ID:          r.Id,
		Title:       r.Title,
		Description: r.Description,
		// UsageLimit:  r.UsageLimit, // 不允许修改使用限制
		// TotalCount:  r.TotalCount, // 不允许修改总数
		StartedAt: *startedAt,
		EndedAt:   *endedAt,
		Status:    r.Status,
	}

	// 更新批次信息
	result := l.svcCtx.DB.Model(&model.RedeemCodeBatch{}).Where("id=?", r.Id).Updates(&updateRedeemCodeBatch)
	if result.Error != nil {
		l.Logger.Errorw("更新兑换码批次失败", logx.Field("err", result.Error), logx.Field("id", r.Id))
		return &emptypb.Empty{}, errorx.ToGrpcError(errorx.ErrInternal)
	}

	// 如果结束时间发生变化，更新关联的兑换码过期时间
	if exist.EndedAt != *endedAt {
		if err := l.svcCtx.DB.Model(&model.RedeemCode{}).Where("redeem_code_batch_id=?", r.Id).Update("expiration_at", *endedAt).Error; err != nil {
			l.Logger.Errorw("更新兑换码过期时间失败", logx.Field("err", err), logx.Field("batchID", r.Id))
		}
	}

	return &emptypb.Empty{}, nil
}
