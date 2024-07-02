package middleware

import (
	"4hfire/common/lib/jwt"
	"context"
	"github.com/mssola/useragent"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	JwtInstance jwt.JWT
}

func NewAuthMiddleware(
	JwtInstance jwt.JWT) *AuthMiddleware {
	return &AuthMiddleware{JwtInstance: JwtInstance}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqCtx := r.Context()
		//	401
		token := r.Header.Get("Authorization")
		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		token = strings.TrimPrefix(token, "Bearer ")
		payload, err := m.JwtInstance.Parse(token)
		if err != nil {
			logx.Errorf("parse token error: %v", err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		account, ok := payload["account"].(string)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		id, ok := payload["id"].(uint64)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		uid, ok := payload["uid"].(string)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		ctx := context.WithValue(reqCtx, "account", account)
		ctx = context.WithValue(ctx, "name", id)
		ctx = context.WithValue(ctx, "uid", uid)
		val := r.Header.Get("User-Agent")
		ua := useragent.New(val)
		ctx = context.WithValue(ctx, "os", ua.OS())
		browser, _ := ua.Browser()
		ctx = context.WithValue(ctx, "broswer", browser)
		ip := r.Header.Get("X-Forwarded-For")
		if ip == "" {
			ip = r.Header.Get("X-Real-Ip")
		}
		if ip == "" {
			ip = r.RemoteAddr
		}
		if strings.Contains(ip, ",") {
			ip = strings.Split(ip, ",")[0]
		}
		if strings.Contains(ip, ":") {
			ip = ip[0:strings.Index(ip, ":")]
		}
		ctx = context.WithValue(ctx, "ip", ip)
		newReq := r.WithContext(ctx)
		next(w, newReq)
	}
}
