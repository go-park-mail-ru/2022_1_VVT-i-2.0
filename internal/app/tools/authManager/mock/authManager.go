package mock

import (
	"errors"
	"time"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/authManager"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/stretchr/testify/mock"
)

type AuthManager struct {
	mock.Mock
}

func (a *AuthManager) CreateToken(payload *authManager.TokenPayload) (string, error) {
	if payload == nil {
		return "", servErrors.NewError(servErrors.CREATE_TOKEN, "")
	}
	return "", nil
}

func (a *AuthManager) ParseToken(token string) (*authManager.TokenPayload, error) {
	panic("implement me")
}

func (a *AuthManager) GetEpiryTime() time.Duration {
	return time.Duration(1)
}

type AuthManagerErr struct {
	mock.Mock
}

func (a *AuthManagerErr) CreateToken(payload *authManager.TokenPayload) (string, error) {
	return "", errors.New("unknown error")
}

func (a *AuthManagerErr) ParseToken(token string) (*authManager.TokenPayload, error) {
	return nil, errors.New("unknown error")
}

func (a *AuthManagerErr) GetEpiryTime() time.Duration {
	return time.Duration(1)
}
