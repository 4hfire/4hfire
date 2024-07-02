package staff

import (
	"4hfire/common/response"
	"net/http"

	"4hfire/api/internal/logic/staff"
	"4hfire/api/internal/svc"
	"4hfire/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 管理员关闭otp
func CloseOTPHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lang := r.Header.Get("Lang")
		var req types.CloseOTPReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := staff.NewCloseOTPLogic(r.Context(), svcCtx)
		err := l.CloseOTP(&req)
		response.Response(w, nil, err, lang)
	}
}
