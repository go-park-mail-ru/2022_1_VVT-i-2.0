package mock

import (
	"errors"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/auth/models"
)

type AuthUcase struct {
}

func (u *AuthUcase) Login(req *models.LoginUcaseReq) (*models.LogitUcaseResp, error) {
	return &models.LogitUcaseResp{Id: 1, Name: "Name", Phone: "79999999999", Email: "mail.ru", Avatar: "avatar", Addres: "адрес"}, nil
}

func (u *AuthUcase) Register(req *models.RegisterUcaseReq) (*models.UserDataUcase, error) {
	return &models.UserDataUcase{Id: 1, Name: "Name", Phone: "79999999999", Email: "mail.ru", Avatar: "avatar"}, nil
}

func (u *AuthUcase) SendCode(*models.SendCodeUcaseReq) (models.SendCodeUcaseResp, error) {
	return models.SendCodeUcaseResp{IsRegistered: false}, nil
}

type AuthUcaseErr struct {
}

func (u *AuthUcaseErr) Login(req *models.LoginUcaseReq) (*models.LogitUcaseResp, error) {
	return nil, errors.New("unknown error")
}

func (u *AuthUcaseErr) Register(req *models.RegisterUcaseReq) (*models.UserDataUcase, error) {
	return nil, errors.New("unknown error")
}

func (u *AuthUcaseErr) SendCode(*models.SendCodeUcaseReq) (models.SendCodeUcaseResp, error) {
	return models.SendCodeUcaseResp{IsRegistered: false}, errors.New("unknown error")
}
