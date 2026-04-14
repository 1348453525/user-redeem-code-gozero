// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package user

import (
	"net/http"

	"github.com/1348453525/user-redeem-code-gozero/api/internal/logic/user"
	"github.com/1348453525/user-redeem-code-gozero/api/internal/svc"
	"github.com/1348453525/user-redeem-code-gozero/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewGetListLogic(r.Context(), svcCtx)
		resp, err := l.GetList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
