package logic

import (
	"context"

	"github.com/1348453525/user-redeem-code-gozero/model"
	"github.com/1348453525/user-redeem-code-gozero/pkg/errorx"
	"github.com/1348453525/user-redeem-code-gozero/pkg/helper"
	"github.com/1348453525/user-redeem-code-gozero/redeem-code-rpc/internal/svc"
	proto "github.com/1348453525/user-redeem-code-gozero/redeem-code-rpc/redeemcode"
	"github.com/gofrs/uuid/v5"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRedeemCodeBatchLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateRedeemCodeBatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRedeemCodeBatchLogic {
	return &CreateRedeemCodeBatchLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建兑换码批次
func (l *CreateRedeemCodeBatchLogic) CreateRedeemCodeBatch(r *proto.CreateRedeemCodeBatchRequest) (*proto.RedeemCodeBatchResponse, error) {
	// 参数验证
	if r.Title == "" || r.TotalCount <= 0 || r.UsageLimit <= 0 {
		return nil, errorx.ToGrpcError(errorx.ErrParam)
	}

	// 解析日期
	startedAt, err := helper.ParseDatetime(r.StartedAt)
	if err != nil {
		l.Logger.Errorw("解析开始时间失败", logx.Field("err", err), logx.Field("startedAt", r.StartedAt))
		return nil, errorx.ToGrpcError(errorx.ErrParam)
	}
	endedAt, err := helper.ParseDatetime(r.EndedAt)
	if err != nil {
		l.Logger.Errorw("解析结束时间失败", logx.Field("err", err), logx.Field("endedAt", r.EndedAt))
		return nil, errorx.ToGrpcError(errorx.ErrParam)
	}

	// 验证时间范围
	if endedAt.Before(*startedAt) {
		l.Logger.Errorw("结束时间早于开始时间", logx.Field("startedAt", r.StartedAt), logx.Field("endedAt", r.EndedAt))
		return nil, errorx.ToGrpcError(errorx.ErrParam)
	}

	// 使用事务确保原子性
	tx := l.svcCtx.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			l.Logger.Errorw("事务panic", logx.Field("panic", r))
		}
	}()

	// 创建兑换码批次
	newRedeemCodeBatch := model.RedeemCodeBatch{
		Title:       r.Title,
		Description: r.Description,
		UsageLimit:  r.UsageLimit,
		TotalCount:  r.TotalCount,
		StartedAt:   *startedAt,
		EndedAt:     *endedAt,
		Status:      1,
		CreatorID:   r.CreatorId,
		CreatorName: r.CreatorName,
	}
	if result := tx.Create(&newRedeemCodeBatch); result.Error != nil {
		tx.Rollback()
		l.Logger.Errorw("创建兑换码批次失败", logx.Field("err", result.Error))
		return nil, errorx.ToGrpcError(errorx.ErrOperationFailed)
	}

	// 批量创建兑换码（分批次处理，避免内存问题）
	const batchSize = 1000
	totalCount := int(r.TotalCount)
	for i := 0; i < totalCount; i += batchSize {
		end := i + batchSize
		if end > totalCount {
			end = totalCount
		}

		var newRedeemCodes []model.RedeemCode
		for j := i; j < end; j++ {
			redeemCode := model.RedeemCode{
				RedeemCodeBatchID: newRedeemCodeBatch.ID,
				Title:             r.Title,
				Value:             uuid.Must(uuid.NewV7()).String(),
				UsageLimit:        r.UsageLimit,
				ExpirationAt:      *endedAt,
				IsDel:             2,
			}
			newRedeemCodes = append(newRedeemCodes, redeemCode)
		}

		if result := tx.Create(&newRedeemCodes); result.Error != nil {
			tx.Rollback()
			l.Logger.Errorw("批量创建兑换码失败", logx.Field("err", result.Error))
			return nil, errorx.ToGrpcError(errorx.ErrOperationFailed)
		}
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		l.Logger.Errorw("提交事务失败", logx.Field("err", err))
		return nil, errorx.ToGrpcError(errorx.ErrOperationFailed)
	}

	// 返回数据
	return &proto.RedeemCodeBatchResponse{
		Id:          newRedeemCodeBatch.ID,
		Title:       newRedeemCodeBatch.Title,
		Description: newRedeemCodeBatch.Description,
		UsageLimit:  newRedeemCodeBatch.UsageLimit,
		TotalCount:  newRedeemCodeBatch.TotalCount,
		UsedCount:   newRedeemCodeBatch.UsedCount,
		StartedAt:   timestamppb.New(newRedeemCodeBatch.StartedAt),
		EndedAt:     timestamppb.New(newRedeemCodeBatch.EndedAt),
		Status:      newRedeemCodeBatch.Status,
		CreatorId:   newRedeemCodeBatch.CreatorID,
		CreatorName: newRedeemCodeBatch.CreatorName,
		CreatedAt:   timestamppb.New(newRedeemCodeBatch.CreatedAt),
		UpdatedAt:   timestamppb.New(newRedeemCodeBatch.UpdatedAt),
	}, nil
}
