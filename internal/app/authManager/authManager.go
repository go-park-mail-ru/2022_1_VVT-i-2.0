package authManager

import (
	"fmt"
	"time"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
)

type TokenPayload struct {
	Id  models.UserId `valid:"userId, required"`
	Exp time.Time     `valid:"expired, required"`
}

const (
	idTitle      = "id"
	expiresTitle = "expired"
)

func NewTokenPayload(id models.UserId) *TokenPayload {
	return &TokenPayload{
		Id: id,
	}
}

type AuthManager interface {
	CreateToken(payload TokenPayload) (string, error)
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
	fmt.Println(expStr)
	exp, _ := time.Parse(time.RFC3339, expStr)
	fmt.Println(exp)
	return &TokenPayload{
		Id:  models.UserId(payloadMap[idTitle].(float64)),
		Exp: exp,
	}
}
