package jwtManager

import (
	"time"

	conf "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/config"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/authManager"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
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

func (manager *JwtManager) GetEpiryTime() time.Duration {
	return manager.expDuration
}

func (manager *JwtManager) CreateToken(payload *authManager.TokenPayload) (string, error) {
	payload.Exp = time.Now().Add(manager.expDuration)
	token := jwt.NewWithClaims(manager.method, jwt.MapClaims(authManager.TokenPayloadToMap(*payload)))
	if token == nil {
		return "", servErrors.NewError(servErrors.CREATE_TOKEN, "error creating jwt-token")
	}
	tokenSigned, err := token.SignedString(manager.key)
	if err != nil {
		return "", servErrors.NewError(servErrors.CREATE_TOKEN, "error signing jwt-token: "+err.Error())
	}
	return tokenSigned, nil
}

func (manager *JwtManager) ParseToken(token string) (*authManager.TokenPayload, error) {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return manager.key, nil
	})

	if jwtToken == nil || err != nil {
		return nil, servErrors.NewError(servErrors.PARSE_TOKEN, "error parsing jwt-token: "+err.Error())
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok || !jwtToken.Valid {
		return nil, servErrors.NewError(servErrors.WRONG_TOKEN_CLAIMS, "jwt-token not valid")
	}
	return authManager.MapToTokenPayload(claims), nil
}
