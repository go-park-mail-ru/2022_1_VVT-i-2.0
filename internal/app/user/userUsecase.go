package user

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
)

type Usecase interface {
	Register(req *models.RegisterUcaseReq) (*models.UserDataUcase, error)
	Login(req *models.LoginUcaseReq) (*models.UserDataUcase, error)
	SendCode(req *models.SendCodeUcaseReq) (models.SendCodeUcaseResp, error)
	GetUser(id models.UserId) (*models.UserDataUcase, error)
	UpdateUser(req *models.UpdateUserUsecase) (*models.UserDataUcase, error)
}
