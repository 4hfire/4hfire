package staff

import (
	"4hfire/common/response"
	"net/http"

	"4hfire/api/internal/logic/staff"
	"4hfire/api/internal/svc"
	"4hfire/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 管理员更新数据
func UpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lang := r.Header.Get("Lang")
		var req types.AdminUpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := staff.NewUpdateLogic(r.Context(), svcCtx)
		err := l.Update(&req)
		response.Response(w, nil, err, lang)
	}
}
