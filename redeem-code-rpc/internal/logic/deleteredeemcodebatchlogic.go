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
func (l *DeleteRedeemCodeBatchLogic) DeleteRedeemCodeBatch(r *proto.IDRequest) (*emptypb.Empty, error) {
	// 参数验证
	if r.Id <= 0 {
		return &emptypb.Empty{}, errorx.ToGrpcError(errorx.ErrParam)
	}

	// 验证批次是否存在
	if _, err := model.GetRedeemCodeBatchByID(l.svcCtx.DB, r.Id); err != nil {
		l.Logger.Errorw("获取兑换码批次失败", logx.Field("err", err), logx.Field("id", r.Id))
		return &emptypb.Empty{}, errorx.ToGrpcError(errorx.ErrInternal)
	}

	// 使用事务确保原子性
	tx := l.svcCtx.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			l.Logger.Errorw("事务panic", logx.Field("panic", r))
		}
	}()

	// 更新批次状态为删除
	if err := tx.Model(&model.RedeemCodeBatch{}).Where("id=?", r.Id).Update("status", 2).Error; err != nil {
		tx.Rollback()
		l.Logger.Errorw("删除兑换码批次失败", logx.Field("err", err), logx.Field("id", r.Id))
		return &emptypb.Empty{}, errorx.ToGrpcError(errorx.ErrInternal)
	}

	// 更新关联的兑换码状态
	if err := tx.Model(&model.RedeemCode{}).Where("redeem_code_batch_id=?", r.Id).Update("is_del", 1).Error; err != nil {
		tx.Rollback()
		l.Logger.Errorw("删除兑换码失败", logx.Field("err", err), logx.Field("batchID", r.Id))
		return &emptypb.Empty{}, errorx.ToGrpcError(errorx.ErrInternal)
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		l.Logger.Errorw("提交事务失败", logx.Field("err", err))
		return &emptypb.Empty{}, errorx.ToGrpcError(errorx.ErrOperationFailed)
	}

	return &emptypb.Empty{}, nil
}
