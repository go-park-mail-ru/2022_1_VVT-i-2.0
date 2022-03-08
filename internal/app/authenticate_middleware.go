package serv

import (
	"context"
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

type keyCtx string

const keyUserId keyCtx = "user"

type ctxUserId uint64

func (s *server) authOptMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("22")
		tokenCookie, err := r.Cookie("token")
		fmt.Println("24")
		if err != nil {
			fmt.Println("26")
			next.ServeHTTP(w, r)
			fmt.Println("2")
			return
		}

		fmt.Println("29")
		token, err := parseToken(tokenCookie.Value)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		fmt.Println("36")
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			next.ServeHTTP(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), keyUserId, ctxUserId(uint64((claims["userId"]).(float64))))

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (s *server) authRequiredMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenCookie, err := r.Cookie("token")
		if err != nil {
			http.Error(w, `{"error":"auth required"}`, http.StatusUnauthorized)
			return
		}

		token, err := parseToken(tokenCookie.Value)

		if err != nil {
			http.Error(w, `{"error":"auth required"}`, http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			http.Error(w, `{"error":"auth required"}`, http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), keyUserId, ctxUserId(uint64((claims["userId"]).(float64))))

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
