package auth

import (
	"4hfire/common/response"
	"net/http"

	"4hfire/api/internal/logic/auth"
	"4hfire/api/internal/svc"
)

// 管理-管理用户详情
func InfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lang := r.Header.Get("Lang")
		l := auth.NewInfoLogic(r.Context(), svcCtx)
		resp, err := l.Info()
		response.Response(w, resp, err, lang)
	}
}
