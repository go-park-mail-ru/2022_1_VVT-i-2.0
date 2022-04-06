package authManager

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
)

type TokenPayload struct {
	Id models.UserId
}

const (
	idTitle = "id"
)

type AuthManager interface {
	CreateToken(payload TokenPayload) (string, error)
	ParseToken(token string) (TokenPayload, error)
}

func TokenPayloadToMap(payload TokenPayload) map[string]interface{} {
	return map[string]interface{}{
		idTitle: payload.Id,
	}
}

func MapToTokenPayload(payloadMap map[string]interface{}) TokenPayload {
	return TokenPayload{
		Id: models.UserId(payloadMap[idTitle].(float64)),
	}
}
