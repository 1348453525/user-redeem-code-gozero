// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package redeemcodebatch

import (
	"net/http"

	"github.com/1348453525/user-redeem-code-gozero/api/internal/logic/redeemcodebatch"
	"github.com/1348453525/user-redeem-code-gozero/api/internal/svc"
	"github.com/1348453525/user-redeem-code-gozero/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateRedeemCodeBatchHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateRedeemCodeBatchRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := redeemcodebatch.NewUpdateRedeemCodeBatchLogic(r.Context(), svcCtx)
		err := l.UpdateRedeemCodeBatch(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
