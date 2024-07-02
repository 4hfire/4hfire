package staff

import (
	"4hfire/common/response"
	"net/http"

	"4hfire/api/internal/logic/staff"
	"4hfire/api/internal/svc"
	"4hfire/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 管理员删除管理账户
func DeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lang := r.Header.Get("Lang")
		var req types.AdminDeleteReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := staff.NewDeleteLogic(r.Context(), svcCtx)
		err := l.Delete(&req)
		response.Response(w, nil, err, lang)
	}
}
