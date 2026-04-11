package userlogic

import (
	"context"

	"github.com/1348453525/user-redeem-code-gozero/user-rpc/internal/svc"
	"github.com/1348453525/user-redeem-code-gozero/user-rpc/proto/user"
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
func (l *UpdateLogic) Update(in *user.UpdateUserRequest) (*emptypb.Empty, error) {
	// todo: add your logic here and delete this line

	return &emptypb.Empty{}, nil
}
