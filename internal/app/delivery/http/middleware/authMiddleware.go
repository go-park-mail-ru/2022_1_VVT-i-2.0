package middleware

import (
	"context"
	"net/http"

	models "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/serv/models"
)

type KeyCtx string

const KeyUserId KeyCtx = "user"
const TokenKeyCookie = "token"

func (mw *CommonMiddlewareChain) AuthOptMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		tokenCookie, err := ctx.Request().Cookie(TokenKeyCookie)
		if err != nil {
			return next(ctx)
		}

		// token, err :=  token.ParseToken(tokenCookie.Value)
		token, err :=  mw.AuthManager.ParseToken(tokenCookie)
		if err != nil {
			return next(ctx)
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			next.ServeHTTP(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), KeyUserId, models.UserId((claims["userId"]).(float64)))

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func AuthRequiredMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenCookie, err := r.Cookie("token")
		if err != nil {
			http.Error(w, `{"error":"auth required"}`, http.StatusUnauthorized)
			return
		}

		token, err := token.ParseToken(tokenCookie.Value)

		if err != nil {
			http.Error(w, `{"error":"auth required"}`, http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			http.Error(w, `{"error":"auth required"}`, http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), KeyUserId, models.UserId((claims["userId"]).(float64)))

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
