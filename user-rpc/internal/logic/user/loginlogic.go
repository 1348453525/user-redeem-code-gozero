package userlogic

import (
	"context"
	"errors"

	"github.com/1348453525/user-redeem-code-gozero/model"
	"github.com/1348453525/user-redeem-code-gozero/pkg/errorx"
	"github.com/1348453525/user-redeem-code-gozero/pkg/jwt"
	"github.com/1348453525/user-redeem-code-gozero/user-rpc/internal/svc"
	proto "github.com/1348453525/user-redeem-code-gozero/user-rpc/proto/user"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 登录
func (l *LoginLogic) Login(r *proto.LoginRequest) (*proto.LoginResponse, error) {
	// 验证用户名长度
	if len(r.Username) < 3 || len(r.Username) > 20 {
		return nil, errorx.ToGrpcError(errorx.ErrParam)
	}

	// 验证密码长度
	if len(r.Password) < 6 {
		return nil, errorx.ToGrpcError(errorx.ErrParam)
	}

	// 查询用户
	var user model.User
	result := l.svcCtx.DB.Where("username = ?", r.Username).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errorx.ToGrpcError(errorx.ErrParam)
		}
		logx.Errorw("查询用户失败", logx.Field("err", result.Error))
		return nil, errorx.ToGrpcError(errorx.ErrInternal)
	}

	// 检验状态
	if user.IsDel == 1 {
		return nil, errorx.ToGrpcError(errorx.ErrUserDisabled)
	}

	// 校验密码
	if !verifyPassword(r.Password, user.Password) {
		return nil, errorx.ToGrpcError(errorx.ErrPasswordError)
	}

	// 生成 token
	token, err := jwt.GenerateToken(user.ID)
	if err != nil {
		logx.Errorw("生成token失败", logx.Field("err", result.Error))
		return nil, errorx.ToGrpcError(errorx.ErrInternal)
	}

	// 返回数据
	userInfo := buildUserInfoDvo(&user)
	resp := &proto.LoginResponse{
		Info:  &userInfo,
		Token: token,
	}
	return resp, nil
}
