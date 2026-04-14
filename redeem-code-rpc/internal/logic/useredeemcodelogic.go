package logic

import (
	"context"
	"time"

	"github.com/1348453525/user-redeem-code-gozero/model"
	"github.com/1348453525/user-redeem-code-gozero/pkg/errorx"
	"github.com/1348453525/user-redeem-code-gozero/redeem-code-rpc/internal/svc"
	proto "github.com/1348453525/user-redeem-code-gozero/redeem-code-rpc/redeemcode"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type UseRedeemCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUseRedeemCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UseRedeemCodeLogic {
	return &UseRedeemCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 使用兑换码
func (l *UseRedeemCodeLogic) UseRedeemCode(r *proto.UseRedeemCodeRequest) (*emptypb.Empty, error) {
	l.Logger.Infow("兑换码使用开始", logx.Field("r", r))
	// 开启事务
	tx := l.svcCtx.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			l.Logger.Errorw("事务panic", logx.Field("panic", r))
		}
	}()

	// 获取兑换码信息
	redeemCode, err := model.GetRedeemCodeByID(l.svcCtx.DB, r.RedeemCodeId)
	if err != nil {
		tx.Rollback()
		l.Logger.Errorw("获取兑换码失败", logx.Field("err", err), logx.Field("r", r))
		return &emptypb.Empty{}, errorx.ToGrpcError(errorx.ErrInternal)
	}

	// 检查兑换码状态
	if redeemCode.IsDel == 1 {
		tx.Rollback()
		return &emptypb.Empty{}, errorx.ToGrpcError(errorx.ErrRedeemCodeInvalid)
	}

	// 检查兑换码是否已过期
	if time.Now().After(redeemCode.ExpirationAt) {
		tx.Rollback()
		return &emptypb.Empty{}, errorx.ToGrpcError(errorx.ErrRedeemCodeExpired)
	}

	// 检查兑换码是否已达到使用上限
	if redeemCode.UsedCount >= redeemCode.UsageLimit {
		tx.Rollback()
		return &emptypb.Empty{}, errorx.ToGrpcError(errorx.ErrRedeemCodeUsedUp)
	}

	// 创建使用记录
	redeemCodeRecord := model.RedeemCodeRecord{
		UserID:       r.UserId,
		RedeemCodeID: r.RedeemCodeId,
	}
	if err := tx.Create(&redeemCodeRecord).Error; err != nil {
		tx.Rollback()
		l.Logger.Errorw("创建使用记录失败", logx.Field("err", err), logx.Field("r", r))
		return &emptypb.Empty{}, errorx.ToGrpcError(errorx.ErrInternal)
	}

	// 更新兑换码已使用数量（乐观锁：使用 updated_at 作为版本控制）
	result := tx.Model(&model.RedeemCode{}).
		Where("id = ? AND updated_at = ? AND used_count < usage_limit", r.RedeemCodeId, redeemCode.UpdatedAt).
		Update("used_count", gorm.Expr("used_count + 1"))
	if result.Error != nil {
		tx.Rollback()
		l.Logger.Errorw("更新兑换码使用数量失败", logx.Field("err", result.Error), logx.Field("r", r))
		return &emptypb.Empty{}, errorx.ToGrpcError(errorx.ErrInternal)
	}
	if result.RowsAffected == 0 {
		tx.Rollback()
		l.Logger.Errorw("兑换码使用失败：乐观锁失败，可能已被其他用户使用", logx.Field("r", r))
		return &emptypb.Empty{}, errorx.ToGrpcError(errorx.ErrRedeemCodeUsedUp)
	}

	// 如果达到使用上限，更新批次已使用数量
	if redeemCode.UsedCount+1 >= redeemCode.UsageLimit {
		if err := tx.Model(&model.RedeemCodeBatch{}).Where("id = ?", redeemCode.RedeemCodeBatchID).Update("used_count", gorm.Expr("used_count + 1")).Error; err != nil {
			tx.Rollback()
			l.Logger.Errorw("更新批次使用数量失败", logx.Field("err", err), logx.Field("r", r), logx.Field("batch_id", redeemCode.RedeemCodeBatchID))
			return &emptypb.Empty{}, errorx.ToGrpcError(errorx.ErrInternal)
		}
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		l.Logger.Errorw("提交事务失败", logx.Field("err", err), logx.Field("r", r))
		if err := tx.Rollback().Error; err != nil {
			l.Logger.Errorw("事务回滚失败", logx.Field("err", err), logx.Field("r", r))
		}
		return &emptypb.Empty{}, errorx.ToGrpcError(errorx.ErrInternal)
	}

	l.Logger.Infow("兑换码使用成功", logx.Field("r", r))
	return &emptypb.Empty{}, nil
}
