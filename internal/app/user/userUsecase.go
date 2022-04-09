package user

// type

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
)

type Usecase interface {
	Login(req *models.LoginRequest) (*models.UserDataUsecase, error)
	Register(req *models.RegisterRequest) (*models.UserDataUsecase, error)
	SendCode(req *models.SendCodeReq) (bool, error)
}
