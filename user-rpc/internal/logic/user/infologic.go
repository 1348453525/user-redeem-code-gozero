package userlogic

import (
	"context"
	"errors"

	"github.com/1348453525/user-redeem-code-gozero/model"
	"github.com/1348453525/user-redeem-code-gozero/pkg/errorx"
	"github.com/1348453525/user-redeem-code-gozero/user-rpc/internal/svc"
	proto "github.com/1348453525/user-redeem-code-gozero/user-rpc/proto/user"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type InfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoLogic {
	return &InfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户信息
func (l *InfoLogic) Info(r *proto.IDRequest) (*proto.UserInfoResponse, error) {
	// 查询用户
	userInfo, err := model.GetUserByID(l.svcCtx.DB, r.Id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errorx.ToGrpcError(errorx.ErrUserNotExisted)
	}
	if err != nil {
		return nil, errorx.ToGrpcError(err)
	}

	// 返回数据
	resp := buildUserInfoDvo(userInfo)
	return &resp, nil
}
