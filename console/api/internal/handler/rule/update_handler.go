package rule

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"4hfire/api/internal/logic/rule"
	"4hfire/api/internal/svc"
	"4hfire/api/internal/types"
	"4hfire/common/errors"
	"4hfire/common/response"
)

func UpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lang := r.Header.Get("Lang")
		var req types.RuleUpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			logx.Error(err)
			response.Response(w, nil, errors.ParamsError(), lang)
			return
		}
		l := rule.NewUpdateLogic(r.Context(), svcCtx)
		err := l.Update(&req)
		response.Response(w, nil, err, lang)
	}
}
