package staff

import (
	"4hfire/common/response"
	"net/http"

	"4hfire/api/internal/logic/staff"
	"4hfire/api/internal/svc"
)

// 管理员获取otp开启信息
func OpenOTPInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lang := r.Header.Get("Lang")
		l := staff.NewOpenOTPInfoLogic(r.Context(), svcCtx)
		resp, err := l.OpenOTPInfo()
		response.Response(w, resp, err, lang)
	}
}
