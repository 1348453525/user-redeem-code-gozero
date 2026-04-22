// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package user

import (
	"context"

	"github.com/1348453525/user-redeem-code-gozero/api/internal/svc"
	"github.com/1348453525/user-redeem-code-gozero/api/internal/types"
	userclient "github.com/1348453525/user-redeem-code-gozero/user-rpc/client/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.UserInfoResponse, err error) {
	respRpc, err := l.svcCtx.UserRpc.Register(l.ctx, &userclient.RegisterRequest{
		Username:        req.Username,
		Password:        req.Password,
		ConfirmPassword: req.ConfirmPassword,
		Nickname:        req.Nickname,
		Mobile:          req.Mobile,
		Gender:          req.Gender,
		Birthday:        req.Birthday,
	})
	if err != nil {
		logx.Errorw("UserRpc.Register error", logx.Field("err", err))
		return nil, err
	}

	return &types.UserInfoResponse{
		Id:       respRpc.Id,
		Username: respRpc.Username,
		Nickname: respRpc.Nickname,
		Mobile:   respRpc.Mobile,
		Gender:   respRpc.Gender,
		Birthday: respRpc.Birthday,
	}, nil
}
