package auth

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/auth/models"
)

type Usecase interface {
	// Login(req *models.LoginReq) (*models.UserDataUsecase, error)
	// Register(req *models.RegisterReq) (*models.UserDataUsecase, error)
	// SendCode(req *models.SendCodeReq) (bool, error)
	SendCode(*models.SendCodeUcaseReq) (models.SendCodeUcaseResp, error)
}
