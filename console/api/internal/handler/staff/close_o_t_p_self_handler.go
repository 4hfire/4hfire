package staff

import (
	"4hfire/common/response"
	"net/http"

	"4hfire/api/internal/logic/staff"
	"4hfire/api/internal/svc"
)

// 关闭自己otp
func CloseOTPSelfHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lang := r.Header.Get("Lang")
		l := staff.NewCloseOTPSelfLogic(r.Context(), svcCtx)
		err := l.CloseOTPSelf()
		response.Response(w, nil, err, lang)
	}
}
