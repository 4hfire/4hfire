package auth

import (
	"4hfire/common/response"
	"net/http"

	"4hfire/api/internal/logic/auth"
	"4hfire/api/internal/svc"
)

// 管理用户退出
func LogoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lang := r.Header.Get("Lang")
		l := auth.NewLogoutLogic(r.Context(), svcCtx)
		err := l.Logout()
		response.Response(w, nil, err, lang)
	}
}
