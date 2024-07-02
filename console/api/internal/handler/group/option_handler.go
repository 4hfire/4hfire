package group

import (
	"net/http"

	"4hfire/api/internal/logic/group"
	"4hfire/api/internal/svc"
	"4hfire/common/response"
)

func OptionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lang := r.Header.Get("Lang")
		l := group.NewOptionLogic(r.Context(), svcCtx)
		resp, err := l.Option()
		response.Response(w, resp, err, lang)
	}
}
