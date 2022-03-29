package jwtManager

import (
	"fmt"

	conf "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/config"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/authManager"
	jwt "github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
)

type JwtManager struct {
	key    string
	method jwt.SigningMethod
}

type JwtConfig struct {
	Key    string
	Method string
}

func NewJwtManager(cfg conf.AuthManagerConfig) *JwtManager {
	methodObj := jwt.GetSigningMethod(cfg.Method)
	if methodObj == nil {
		return nil
	}
	return &JwtManager{
		key:    cfg.Key,
		method: methodObj,
	}
}

func (manager *JwtManager) CreateToken(payload authManager.TokenPayload) (string, error) {
	token := jwt.NewWithClaims(manager.method, jwt.MapClaims(authManager.TokenPayloadToMap(payload)))
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
