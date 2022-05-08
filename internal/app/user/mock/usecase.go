package mock

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/stretchr/testify/mock"
)

type UserUcase struct {
	mock.Mock
}

func (u *UserUcase) Login(req *models.LoginUcaseReq) (*models.UserDataUcase, error) {
	return &models.UserDataUcase{
		Id:     1,
		Name:   "Name",
		Phone:  "79999999999",
		Email:  "email@mail.com",
		Avatar: "avatar.png"}, nil
}

func (u *UserUcase) Register(req *models.RegisterUcaseReq) (*models.UserDataUcase, error) {
	return &models.UserDataUcase{
		Id:     1,
		Name:   "Name",
		Phone:  "79999999999",
		Email:  "email@mail.com",
		Avatar: "avatar.png"}, nil
}

func (u *UserUcase) SendCode(req *models.SendCodeUcaseReq) (models.SendCodeUcaseResp, error) {
	return models.SendCodeUcaseResp{IsRegistered: true}, nil
}

func (u *UserUcase) GetUser(id models.UserId) (*models.UserDataUcase, error) {
	return &models.UserDataUcase{
		Id:     1,
		Name:   "Name",
		Phone:  "79999999999",
		Email:  "email@mail.com",
		Avatar: "avatar.png"}, nil
}

func (u *UserUcase) UpdateUser(req *models.UpdateUserUcase) (*models.UserDataUcase, error) {
	return &models.UserDataUcase{
		Id:     1,
		Name:   "Name",
		Phone:  "79999999999",
		Email:  "email@mail.com",
		Avatar: "avatar.png"}, nil
}

type UserUcaseErr struct {
	mock.Mock
}

func (u *UserUcaseErr) Login(req *models.LoginUcaseReq) (*models.UserDataUcase, error) {
	return nil, servErrors.NewError(servErrors.DB_ERROR, "")
}

func (u *UserUcaseErr) Register(req *models.RegisterUcaseReq) (*models.UserDataUcase, error) {
	return nil, servErrors.NewError(servErrors.DB_ERROR, "")
}

func (u *UserUcaseErr) SendCode(req *models.SendCodeUcaseReq) (models.SendCodeUcaseResp, error) {
	return models.SendCodeUcaseResp{IsRegistered: false}, servErrors.NewError(servErrors.DB_ERROR, "")
}

func (u *UserUcaseErr) GetUser(id models.UserId) (*models.UserDataUcase, error) {
	return nil, servErrors.NewError(servErrors.DB_ERROR, "")
}

func (u *UserUcaseErr) UpdateUser(req *models.UpdateUserUcase) (*models.UserDataUcase, error) {
	return nil, servErrors.NewError(servErrors.DB_ERROR, "")
}
