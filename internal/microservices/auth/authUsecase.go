package auth

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/auth/models"
)

type Usecase interface {
	Login(req *models.LoginUcaseReq) (*models.UserDataUcase, error)
	Register(req *models.RegisterUcaseReq) (*models.UserDataUcase, error)
	SendCode(*models.SendCodeUcaseReq) (models.SendCodeUcaseResp, error)
}
