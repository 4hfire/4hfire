package staff

import (
	"4hfire/common/response"
	"net/http"

	"4hfire/api/internal/logic/staff"
	"4hfire/api/internal/svc"
	"4hfire/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 管理员账户切换账号禁用状态
func SwitchForbidHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lang := r.Header.Get("Lang")
		var req types.AdminForbidReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := staff.NewSwitchForbidLogic(r.Context(), svcCtx)
		err := l.SwitchForbid(&req)
		response.Response(w, nil, err, lang)
	}
}
