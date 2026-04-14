// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package redeemcodebatch

import (
	"net/http"

	"github.com/1348453525/user-redeem-code-gozero/api/internal/logic/redeemcodebatch"
	"github.com/1348453525/user-redeem-code-gozero/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RedeemCodeBatchDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := redeemcodebatch.NewRedeemCodeBatchDetailLogic(r.Context(), svcCtx)
		resp, err := l.RedeemCodeBatchDetail()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
