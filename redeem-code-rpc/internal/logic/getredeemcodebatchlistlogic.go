package logic

import (
	"context"

	"github.com/1348453525/user-redeem-code-gozero/model"
	"github.com/1348453525/user-redeem-code-gozero/redeem-code-rpc/internal/svc"
	proto "github.com/1348453525/user-redeem-code-gozero/redeem-code-rpc/redeemcode"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRedeemCodeBatchListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRedeemCodeBatchListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRedeemCodeBatchListLogic {
	return &GetRedeemCodeBatchListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取兑换码批次列表
func (l *GetRedeemCodeBatchListLogic) GetRedeemCodeBatchList(r *proto.GetListRequest) (*proto.GetRedeemCodeBatchListResponse, error) {
	list, count := model.GetRedeemCodeBatchList(l.svcCtx.DB, r.Page, r.PageSize)
	resp := &proto.GetRedeemCodeBatchListResponse{
		Page:     r.Page,
		PageSize: r.PageSize,
		Total:    count,
	}
	for _, v := range list {
		redeemCodeBatchTmp := &proto.RedeemCodeBatchResponse{
			Id:          v.ID,
			Title:       v.Title,
			Description: v.Description,
			UsageLimit:  v.UsageLimit,
			TotalCount:  v.TotalCount,
			UsedCount:   v.UsedCount,
			StartedAt:   timestamppb.New(v.StartedAt),
			EndedAt:     timestamppb.New(v.EndedAt),
			Status:      v.Status,
			CreatorId:   v.CreatorID,
			CreatorName: v.CreatorName,
			CreatedAt:   timestamppb.New(v.CreatedAt),
			UpdatedAt:   timestamppb.New(v.UpdatedAt),
		}
		resp.Data = append(resp.Data, redeemCodeBatchTmp)
	}
	return resp, nil
}
