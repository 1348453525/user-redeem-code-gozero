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

type DeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除用户
func (l *DeleteLogic) Delete(r *proto.IDRequest) (*emptypb.Empty, error) {
	// 删除用户
	result := l.svcCtx.DB.Model(&model.User{}).Where("id=?", r.Id).Update("is_del", 1)
	if result.Error != nil {
		logx.Errorw("删除用户失败", logx.Field("err", result.Error))
		return &emptypb.Empty{}, errorx.ToGrpcError(errorx.ErrInternal)
	}

	// 检查是否删除成功
	if result.RowsAffected == 0 {
		return &emptypb.Empty{}, errorx.ToGrpcError(errorx.ErrOperationFailed)
	}
	return &emptypb.Empty{}, nil
}
