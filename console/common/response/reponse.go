/*
 * @Author: Young
 * @Date: 2022-05-11 10:58:40
 * LastEditors: lihy lihy@zhiannet.com
 * LastEditTime: 2022-11-17 14:25:25
 * @FilePath: /zero-trust/console/IAM/common/response/reponse.go
 */

package response

import (
	"4hfire/common/errors"
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type Body struct {
	Code int32       `json:"code"`
	Time int64       `json:"time"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Response(w http.ResponseWriter, resp interface{}, err error, lang string) {
	body := Body{Data: resp, Time: time.Now().Unix()}

	// body.Code = code

	if err != nil {
		errObj, _ := err.(*errors.Error)
		if errObj != nil {
			if errObj.Code() == errors.ErrCodeUnauth {
				httpx.WriteJson(w, http.StatusUnauthorized, nil)
				return
			}
			body.Code = errObj.Code()
			if lang == "" {
				lang = "zh"
			}
			_, ok := errors.ErrMap[lang][errObj.Code()]
			if !ok {
				body.Msg = err.Error()
			} else {
				body.Msg = errors.ErrMap[lang][errObj.Code()]
			}
		} else {
			// 未知错误码
			body.Code = errors.ErrCodeUnkonwn
			body.Msg = err.Error()
		}
	}

	httpx.OkJson(w, body)
}

// func Response(w http.ResponseWriter, resp interface{}, err error) {
// 	body := Body{Data: resp, Time: time.Now().Unix()}

// 	if err != nil {
// 		errObj, _ := err.(*errors.Error)
// 		if errObj != nil {
// 			if errObj.Code() == errors.ErrUnauth {
// 				httpx.WriteJson(w, http.StatusUnauthorized, nil)
// 				return
// 			}
// 			body.Code = errObj.Code()
// 			body.Msg = errObj.Msg()
// 		} else {
// 			// 未知错误码
// 			body.Code = errors.ErrUnkonwn
// 			body.Msg = "系统错误"
// 		}
// 	}

// 	httpx.OkJson(w, body)
// }

func ResponseRaw(w http.ResponseWriter, resp interface{}) {
	var body []byte
	switch v := resp.(type) {
	case []byte:
		body = v
	case string:
		body = []byte(v)
	default:
		body = []byte("1")
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
