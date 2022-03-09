package serv

import (
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var SECRET = []byte("v4fQVUeQ*4r`@TA15m)*")

func createTokenCookie(userId uint64) (http.Cookie, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId})

	tokenStr, err := token.SignedString(SECRET)

	if err != nil {
		return http.Cookie{}, err
	}

	tokenCookie := &http.Cookie{
		Name:     "token",
		Value:    tokenStr,
		Secure:   true,
		HttpOnly: true,
		Path:     "/",
		Expires:  time.Now().AddDate(0, 0, +3),
	}
	return *tokenCookie, nil
}

func parseToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return SECRET, nil
	})
}
