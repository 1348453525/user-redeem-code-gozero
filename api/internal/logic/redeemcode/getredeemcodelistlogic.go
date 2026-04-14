// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package redeemcode

import (
	"context"

	"github.com/1348453525/user-redeem-code-gozero/api/internal/svc"
	"github.com/1348453525/user-redeem-code-gozero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRedeemCodeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRedeemCodeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRedeemCodeListLogic {
	return &GetRedeemCodeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRedeemCodeListLogic) GetRedeemCodeList() (resp *types.GetRedeemCodeListResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
