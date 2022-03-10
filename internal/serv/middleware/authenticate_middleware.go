package middleware

import (
	"context"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	models "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/serv/models"
	token "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/serv/token"
)

type KeyCtx string

const KeyUserId KeyCtx = "user"

func AuthOptMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenCookie, err := r.Cookie("token")
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		token, err := token.ParseToken(tokenCookie.Value)
		if err != nil {
			next.ServeHTTP(w, r)
			return
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
