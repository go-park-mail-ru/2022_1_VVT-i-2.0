package interfaces

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/authManager"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/stretchr/testify/mock"
	"time"
)

type AuthManager struct {
	mock.Mock
}

func (a AuthManager) CreateToken(payload *authManager.TokenPayload) (string, error) {
	if payload == nil {
		return "", servErrors.NewError(servErrors.TEST_ERROR, "")
	}
	return "", nil
}

func (a AuthManager) ParseToken(token string) (*authManager.TokenPayload, error) {
	panic("implement me")
}

func (a AuthManager) GetEpiryTime() time.Duration {
	return time.Duration(1)
}

