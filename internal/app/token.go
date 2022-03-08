package serv

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

var SECRET = []byte("v4fQVUeQ*4r`@TA15m)*")

func createToken(userId uint64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId})

	return token.SignedString(SECRET)
}

func parseToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return SECRET, nil
	})
}
