package auth

import (
	"4hfire/common/response"
	"net/http"

	"4hfire/api/internal/logic/auth"
	"4hfire/api/internal/svc"
	"4hfire/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 用户登录时查看otp开启状态
func OtpStateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lang := r.Header.Get("Lang")
		var req types.OTPStateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := auth.NewOtpStateLogic(r.Context(), svcCtx)
		resp, err := l.OtpState(&req)
		response.Response(w, resp, err, lang)
	}
}
