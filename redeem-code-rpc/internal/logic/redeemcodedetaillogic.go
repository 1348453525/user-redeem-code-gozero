package logic

import (
	"context"

	"github.com/1348453525/user-redeem-code-gozero/model"
	"github.com/1348453525/user-redeem-code-gozero/pkg/errorx"
	"github.com/1348453525/user-redeem-code-gozero/redeem-code-rpc/internal/svc"
	proto "github.com/1348453525/user-redeem-code-gozero/redeem-code-rpc/redeemcode"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RedeemCodeDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRedeemCodeDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RedeemCodeDetailLogic {
	return &RedeemCodeDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取兑换码详情
func (l *RedeemCodeDetailLogic) RedeemCodeDetail(r *proto.IDRequest) (*proto.RedeemCodeResponse, error) {
	// 查询数据
	redeemCode, err := model.GetRedeemCodeByID(l.svcCtx.DB, r.Id)
	if err != nil {
		return nil, errorx.ToGrpcError(errorx.ErrInternal)
	}

	// deleted_at=null，timestamppb.New(*redeemCode.DeletedAt)
	// panic: runtime error: invalid memory address or nil pointer dereference
	// 安全处理 DeletedAt
	var deletedAtPb *timestamppb.Timestamp
	if redeemCode.DeletedAt != nil {
		deletedAtPb = timestamppb.New(*redeemCode.DeletedAt)
	}

	// 返回数据
	return &proto.RedeemCodeResponse{
		Id:                redeemCode.ID,
		RedeemCodeBatchId: redeemCode.RedeemCodeBatchID,
		Title:             redeemCode.Title,
		Value:             redeemCode.Value,
		UsageLimit:        redeemCode.UsageLimit,
		UsedCount:         redeemCode.UsedCount,
		ExpirationAt:      timestamppb.New(redeemCode.ExpirationAt),
		IsDel:             redeemCode.IsDel,
		DeletedAt:         deletedAtPb,
		CreatedAt:         timestamppb.New(redeemCode.CreatedAt),
		UpdatedAt:         timestamppb.New(redeemCode.UpdatedAt),
	}, nil
}
