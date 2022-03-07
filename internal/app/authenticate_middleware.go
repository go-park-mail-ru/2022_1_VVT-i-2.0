package serv

import (
	"context"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"net/http"
)

type keyCtx string

const keyUserId keyCtx = "userId"

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

		ctx := context.WithValue(r.Context(), keyUserId, claims[string(keyUserId)])
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// TODO: set https only for cookies
