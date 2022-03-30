package jwtManager

import (
	"fmt"
	"time"

	conf "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/config"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/authManager"
	jwt "github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
)

type JwtManager struct {
	key         []byte
	method      jwt.SigningMethod
	expDuration time.Duration
}

func NewJwtManager(cfg conf.AuthManagerConfig) *JwtManager {
	methodObj := jwt.GetSigningMethod(cfg.Method)
	if methodObj == nil {
		return nil
	}
	return &JwtManager{
		key:         []byte(cfg.Key),
		method:      methodObj,
		expDuration: cfg.ExpiryTime.Duration,
	}
}

func (manager *JwtManager) CreateToken(payload authManager.TokenPayload) (string, error) {
	// payload.Exp := 1
	payload.Exp = time.Now().Add(manager.expDuration)
	token := jwt.NewWithClaims(manager.method, jwt.MapClaims(authManager.TokenPayloadToMap(payload)))
	// return token.SignedString(manager.key)
	return token.SignedString(manager.key)
}

func (manager *JwtManager) ParseToken(token string) (authManager.TokenPayload, error) {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return manager.key, nil
	})

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok || !jwtToken.Valid {
		return authManager.TokenPayload{}, errors.Wrap(err, "jwt-token not valid")
	}
	return authManager.MapToTokenPayload(claims), nil
}
