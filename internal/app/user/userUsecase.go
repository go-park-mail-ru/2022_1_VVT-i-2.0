package user

// type

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
)

type Usecase interface {
	Login(req *models.LoginRequest) (*models.UserDataUsecase, error)
	Register(req *models.RegisterReq) (*models.UserDataUsecase, error)
	SendCode(req *models.SendCodeReq) (bool, error)
	GetUser(id models.UserId) (*models.UserDataUsecase, error)
	UpdateUser(req *models.UpdateUser) (*models.UserDataUsecase, error)
}
