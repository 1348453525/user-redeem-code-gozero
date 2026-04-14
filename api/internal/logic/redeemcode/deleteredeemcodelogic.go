// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package redeemcode

import (
	"context"

	"github.com/1348453525/user-redeem-code-gozero/api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteRedeemCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteRedeemCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRedeemCodeLogic {
	return &DeleteRedeemCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteRedeemCodeLogic) DeleteRedeemCode() error {
	// todo: add your logic here and delete this line

	return nil
}
