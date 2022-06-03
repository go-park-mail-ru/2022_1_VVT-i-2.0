package authManager

import (
	"time"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/validator"
)

type TokenPayload struct {
	Id  int64
	Exp time.Time
}

const (
	idTitle      = "id"
	expiresTitle = "expired"
)

func NewTokenPayload(id int64) *TokenPayload {
	return &TokenPayload{
		Id: id,
	}
}

type AuthManager interface {
	CreateToken(payload *TokenPayload) (string, error)
	ParseToken(token string) (*TokenPayload, error)
	GetEpiryTime() time.Duration
}

func TokenPayloadToMap(payload TokenPayload) map[string]interface{} {
	return map[string]interface{}{
		idTitle:      payload.Id,
		expiresTitle: payload.Exp,
	}
}

func MapToTokenPayload(payloadMap map[string]interface{}) *TokenPayload {
	expStr, ok := payloadMap[expiresTitle].(string)
	if !ok {
		return nil
	}
	exp, _ := time.Parse(time.RFC3339, expStr)
	return &TokenPayload{
		Id:  int64(payloadMap[idTitle].(float64)),
		Exp: exp,
	}
}

func IsValidToken(payload *TokenPayload) bool {
	return validator.IsNotExpired(payload.Exp) && validator.IsUserId(payload.Id)
}
