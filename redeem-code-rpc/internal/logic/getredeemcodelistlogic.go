package logic

import (
	"context"

	"github.com/1348453525/user-redeem-code-gozero/model"
	"github.com/1348453525/user-redeem-code-gozero/redeem-code-rpc/internal/svc"
	proto "github.com/1348453525/user-redeem-code-gozero/redeem-code-rpc/redeemcode"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRedeemCodeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRedeemCodeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRedeemCodeListLogic {
	return &GetRedeemCodeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取兑换码列表
func (l *GetRedeemCodeListLogic) GetRedeemCodeList(r *proto.GetListRequest) (*proto.GetRedeemCodeListResponse, error) {
	list, count := model.GetRedeemCodeList(l.svcCtx.DB, r.Page, r.PageSize)
	resp := &proto.GetRedeemCodeListResponse{
		Page:     r.Page,
		PageSize: r.PageSize,
		Total:    count,
	}
	for _, v := range list {
		// 安全处理 DeletedAt
		var deletedAtPb *timestamppb.Timestamp
		if v.DeletedAt != nil {
			deletedAtPb = timestamppb.New(*v.DeletedAt)
		}

		redeemCode := &proto.RedeemCodeResponse{
			Id:                v.ID,
			RedeemCodeBatchId: v.RedeemCodeBatchID,
			Title:             v.Title,
			Value:             v.Value,
			UsageLimit:        v.UsageLimit,
			UsedCount:         v.UsedCount,
			ExpirationAt:      timestamppb.New(v.ExpirationAt),
			IsDel:             v.IsDel,
			DeletedAt:         deletedAtPb,
			CreatedAt:         timestamppb.New(v.CreatedAt),
			UpdatedAt:         timestamppb.New(v.UpdatedAt),
		}
		resp.Data = append(resp.Data, redeemCode)
	}
	return resp, nil
}
