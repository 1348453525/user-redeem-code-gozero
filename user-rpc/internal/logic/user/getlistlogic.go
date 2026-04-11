package userlogic

import (
	"context"

	"github.com/1348453525/user-redeem-code-gozero/model"
	"github.com/1348453525/user-redeem-code-gozero/user-rpc/internal/svc"
	proto "github.com/1348453525/user-redeem-code-gozero/user-rpc/proto/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetListLogic {
	return &GetListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户列表
func (l *GetListLogic) GetList(r *proto.GetUserListRequest) (*proto.GetUserListResponse, error) {
	list, count := model.GetUserList(l.svcCtx.DB, r.Page, r.PageSize)
	resp := &proto.GetUserListResponse{
		Page:     r.Page,
		PageSize: r.PageSize,
		Total:    count,
	}
	for _, v := range list {
		userInfoDvo := buildUserInfoDvo(v)
		resp.Data = append(resp.Data, &userInfoDvo)
	}
	return resp, nil
}
