package serv

import (
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	models "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/serv/models"
)

var secret = []byte("v4fQVUeQ*4r`@TA15m)*")

func CreateTokenCookie(userId models.UserId) (http.Cookie, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId})

	tokenStr, err := token.SignedString(secret)

	if err != nil {
		return http.Cookie{}, err
	}

	tokenCookie := &http.Cookie{
		Name:     "token",
		Value:    tokenStr,
		HttpOnly: true,
		Expires:  time.Now().AddDate(0, 0, +3),
	}
	return *tokenCookie, nil
}

func ParseToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return secret, nil
	})
}
