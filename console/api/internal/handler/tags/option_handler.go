package tags

import (
	"net/http"

	"4hfire/api/internal/logic/tags"
	"4hfire/api/internal/svc"
	"4hfire/common/response"
)

func OptionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lang := r.Header.Get("Lang")
		l := tags.NewOptionLogic(r.Context(), svcCtx)
		resp, err := l.Option()
		response.Response(w, resp, err, lang)
	}
}
