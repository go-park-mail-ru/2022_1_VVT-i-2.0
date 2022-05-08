package user

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
)

type Ucase interface {
	Register(req *models.RegisterUcaseReq) (*models.UserDataUcase, error)
	Login(req *models.LoginUcaseReq) (*models.UserDataUcase, error)
	SendCode(req *models.SendCodeUcaseReq) (models.SendCodeUcaseResp, error)
	GetUser(id models.UserId) (*models.UserDataUcase, error)
	UpdateUser(req *models.UpdateUserUcase) (*models.UserDataUcase, error)
}
