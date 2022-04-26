package models

import (
	"database/sql"
)

type UserId uint64

type UserDataStorage struct {
	Id     UserId
	Name   string
	Phone  string
	Email  string
	Avatar sql.NullString
}

type UserAddDataStorage struct {
	Name   string
	Phone  string
	Email  string
	Avatar string
}

type UserDataUsecase struct {
	Id     UserId
	Name   string
	Phone  string
	Email  string
	Avatar string
}

type UpdateAvatarRepo struct {
	ImgPath string
	UserId  UserId
}

type UserDataResp struct {
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	Email  string `json:"email"`
	Avatar string `json:"avatar"`
}

type LoginReq struct {
	Phone string `json:"phone" valid:"phone, required"`
	Code  string `json:"code" valid:"required"`
}

type RegisterReq struct {
	Phone string `json:"phone" valid:"phone, required"`
	Code  string `json:"code" valid:"required"`
	Name  string `json:"name" valid:"name,required"`
	Email string `json:"email" valid:"email,required"`
}

type SendCodeUcaseReq struct {
	Phone string `json:"phone" valid:"phone, required"`
}

type SendCodeUcaseResp struct {
	IsRegistered bool `json:"registered"`
}

type SendCodeRepoReq struct {
	Phone string `json:"phone" valid:"phone, required"`
}

type SendCodeRepoResp struct {
	IsRegistered bool `json:"registered"`
}
