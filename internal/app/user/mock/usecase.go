package mock

import (
	"errors"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/stretchr/testify/mock"
)

type UserUcase struct {
	mock.Mock
}

func (u *UserUcase) Login(req *models.LoginUcaseReq) (*models.LoginUcaseResp, error) {
	return &models.LoginUcaseResp{
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

func (u *UserUcase) GetUser(id int64) (*models.UserDataUcase, error) {
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

type UserUcaseDBErr struct {
	mock.Mock
}

func (u *UserUcaseDBErr) Login(req *models.LoginUcaseReq) (*models.LoginUcaseResp, error) {
	return nil, servErrors.NewError(servErrors.DB_ERROR, "")
}

func (u *UserUcaseDBErr) Register(req *models.RegisterUcaseReq) (*models.UserDataUcase, error) {
	return nil, servErrors.NewError(servErrors.DB_ERROR, "")
}

func (u *UserUcaseDBErr) SendCode(req *models.SendCodeUcaseReq) (models.SendCodeUcaseResp, error) {
	return models.SendCodeUcaseResp{IsRegistered: false}, servErrors.NewError(servErrors.DB_ERROR, "")
}

func (u *UserUcaseDBErr) GetUser(id int64) (*models.UserDataUcase, error) {
	return nil, servErrors.NewError(servErrors.DB_ERROR, "")
}

func (u *UserUcaseDBErr) UpdateUser(req *models.UpdateUserUcase) (*models.UserDataUcase, error) {
	return nil, servErrors.NewError(servErrors.DB_ERROR, "")
}

type UserUcaseUnknownErr struct {
	mock.Mock
}

func (u *UserUcaseUnknownErr) Login(req *models.LoginUcaseReq) (*models.LoginUcaseResp, error) {
	return nil, errors.New("unknown error")
}

func (u *UserUcaseUnknownErr) Register(req *models.RegisterUcaseReq) (*models.UserDataUcase, error) {
	return nil, errors.New("unknown error")
}

func (u *UserUcaseUnknownErr) SendCode(req *models.SendCodeUcaseReq) (models.SendCodeUcaseResp, error) {
	return models.SendCodeUcaseResp{IsRegistered: false}, errors.New("unknown error")
}

func (u *UserUcaseUnknownErr) GetUser(id int64) (*models.UserDataUcase, error) {
	return nil, errors.New("unknown error")
}

func (u *UserUcaseUnknownErr) UpdateUser(req *models.UpdateUserUcase) (*models.UserDataUcase, error) {
	return nil, errors.New("unknown error")
}
