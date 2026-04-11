package userlogic

import (
	"context"

	"github.com/1348453525/user-redeem-code-gozero/model"
	"github.com/1348453525/user-redeem-code-gozero/pkg/errorx"
	"github.com/1348453525/user-redeem-code-gozero/user-rpc/internal/svc"
	proto "github.com/1348453525/user-redeem-code-gozero/user-rpc/proto/user"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新用户信息
func (l *UpdateLogic) Update(r *proto.UpdateUserRequest) (*emptypb.Empty, error) {
	// 查询用户是否存在
	var count int64
	result := l.svcCtx.DB.Model(&model.User{}).Where("id = ?", r.Id).Count(&count)
	if result.Error != nil {
		logx.Errorw("查询用户失败", logx.Field("err", result.Error))
		return &emptypb.Empty{}, errorx.ToGrpcError(errorx.ErrInternal)
	}
	if count == 0 {
		return &emptypb.Empty{}, errorx.ToGrpcError(errorx.ErrParam)
	}

	// 处理生日
	birthday := parseBirthday(r.Birthday)

	// 构建更新数据
	updateData := map[string]interface{}{
		"username": r.Username,
		"nickname": r.Nickname,
		"mobile":   r.Mobile,
		"gender":   r.Gender,
		"birthday": birthday,
	}
	result = l.svcCtx.DB.Model(&model.User{}).Where("id=?", r.Id).Updates(updateData)
	if result.Error != nil {
		logx.Errorw("更新用户失败", logx.Field("err", result.Error))
		return &emptypb.Empty{}, errorx.ToGrpcError(errorx.ErrInternal)
	}

	// 检查是否更新成功
	if result.RowsAffected == 0 {
		return &emptypb.Empty{}, errorx.ToGrpcError(errorx.ErrOperationFailed)
	}
	return &emptypb.Empty{}, nil
}
