package authManager

import "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"

// interface manager

type TokenPayload struct {
	id models.UserId
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
		idTitle: payload.id,
	}
}

func MapToTokenPayload(payloadMap map[string]interface{}) TokenPayload {
	return TokenPayload{
		id: payloadMap[idTitle].(models.UserId),
	}
}
