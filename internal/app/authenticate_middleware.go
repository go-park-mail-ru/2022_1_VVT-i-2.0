package serv

import (
	"context"
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

type keyCtx string

const keyUser keyCtx = "user"

type ctxStruct struct {
	user userDataStruct
}

var SECRET = []byte("v4fQVUeQ*4r`@TA15m)*")

func (s *server) authOptMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenCookie, err := r.Cookie("token")
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}
		token, err := jwt.Parse(tokenCookie.Value, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return SECRET, nil
		})
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			next.ServeHTTP(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), keyUser, ctxStruct{userDataStruct{id: uint64((claims["userId"]).(float64)), name: (claims["username"]).(string), address: (claims["userAddress"]).(string)}})

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
		token, err := jwt.Parse(tokenCookie.Value, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return SECRET, nil
		})
		if err != nil {
			http.Error(w, `{"error":"auth required"}`, http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			http.Error(w, `{"error":"auth required"}`, http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), keyUser, ctxStruct{userDataStruct{id: uint64((claims["userId"]).(float64)), name: (claims["username"]).(string), address: (claims["userAddress"]).(string)}})

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
