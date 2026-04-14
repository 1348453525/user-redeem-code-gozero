package logic

import (
	"context"

	"github.com/1348453525/user-redeem-code-gozero/model"
	"github.com/1348453525/user-redeem-code-gozero/pkg/errorx"
	"github.com/1348453525/user-redeem-code-gozero/redeem-code-rpc/internal/svc"
	proto "github.com/1348453525/user-redeem-code-gozero/redeem-code-rpc/redeemcode"
	"google.golang.org/protobuf/types/known/timestamppb"

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
func (l *RedeemCodeBatchDetailLogic) RedeemCodeBatchDetail(r *proto.IDRequest) (*proto.RedeemCodeBatchResponse, error) {
	// 获取批次信息
	redeemCodeBatch, err := model.GetRedeemCodeBatchByID(l.svcCtx.DB, r.Id)
	if err != nil {
		l.Logger.Errorw("获取兑换码批次失败", logx.Field("err", err), logx.Field("id", r.Id))
		return nil, errorx.ToGrpcError(errorx.ErrInternal)
	}

	// 返回数据
	return &proto.RedeemCodeBatchResponse{
		Id:          redeemCodeBatch.ID,
		Title:       redeemCodeBatch.Title,
		Description: redeemCodeBatch.Description,
		UsageLimit:  redeemCodeBatch.UsageLimit,
		TotalCount:  redeemCodeBatch.TotalCount,
		UsedCount:   redeemCodeBatch.UsedCount,
		StartedAt:   timestamppb.New(redeemCodeBatch.StartedAt),
		EndedAt:     timestamppb.New(redeemCodeBatch.EndedAt),
		Status:      redeemCodeBatch.Status,
		CreatorId:   redeemCodeBatch.CreatorID,
		CreatorName: redeemCodeBatch.CreatorName,
		CreatedAt:   timestamppb.New(redeemCodeBatch.CreatedAt),
		UpdatedAt:   timestamppb.New(redeemCodeBatch.UpdatedAt),
	}, nil
}
