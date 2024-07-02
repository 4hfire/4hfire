package tags

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"4hfire/api/internal/logic/tags"
	"4hfire/api/internal/svc"
	"4hfire/api/internal/types"
	"4hfire/common/errors"
	"4hfire/common/response"
)

func DeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lang := r.Header.Get("Lang")
		var req types.TagDeleteReq
		if err := httpx.Parse(r, &req); err != nil {
			logx.Error(err)
			response.Response(w, nil, errors.ParamsError(), lang)
			return
		}
		l := tags.NewDeleteLogic(r.Context(), svcCtx)
		err := l.Delete(&req)
		response.Response(w, nil, err, lang)
	}
}
