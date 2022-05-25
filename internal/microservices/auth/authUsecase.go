package auth

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/auth/models"
)

type Ucase interface {
	Login(req *models.LoginUcaseReq) (*models.LogitUcaseResp, error)
	Register(req *models.RegisterUcaseReq) (*models.UserDataUcase, error)
	SendCode(*models.SendCodeUcaseReq) (models.SendCodeUcaseResp, error)
}
