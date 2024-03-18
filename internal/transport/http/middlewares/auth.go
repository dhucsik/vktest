package middlewares

import (
	"net/http"

	"github.com/dhucsik/vktest/internal/errors"
	"github.com/dhucsik/vktest/internal/util/jwt"
)

const (
	authorizationHeader = "Authorization"
)

type AuthMiddleware struct{}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) Handler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get(authorizationHeader)
		session, isRefresh, err := jwt.ParseJWT(token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if isRefresh {
			http.Error(w, errors.ErrUnexpectedRefresh.Error(), http.StatusUnauthorized)
			return
		}

		if session == nil {
			http.Error(w, errors.ErrInvalidSession.Error(), http.StatusUnauthorized)
			return
		}

		sessionCtx := session.SetInCtx(r.Context())
		next(w, r.WithContext(sessionCtx))
	}
}
