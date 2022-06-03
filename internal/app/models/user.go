package models

import (
	"database/sql"
	"io"
)

type UserDataRepo struct {
	Id     int64
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

type UserDataUcase struct {
	Id     int64
	Name   string
	Phone  string
	Email  string
	Avatar string
}

type UpdateAvatarRepo struct {
	ImgPath string
	UserId  int64
}

//easyjson:json
type UserDataResp struct {
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	Email  string `json:"email"`
	Avatar string `json:"avatar"`
}

//easyjson:json
type LoginReq struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}

//easyjson:json
type LoginResp struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Avatar  string `json:"avatar"`
	Address string `json:"address"`
}

type LoginUcaseReq struct {
	Phone string
	Code  string
}

type LoginUcaseResp struct {
	Id      int64
	Name    string
	Phone   string
	Email   string
	Avatar  string
	Address string
}

//easyjson:json
type RegisterReq struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type RegisterUcaseReq struct {
	Phone string
	Code  string
	Name  string
	Email string
}

//easyjson:json
type SendCodeReq struct {
	Phone string `json:"phone"`
}

//easyjson:json
type SendCodeResp struct {
	IsRegistered bool `json:"registered"`
}

type SendCodeUcaseReq struct {
	Phone string
}

type SendCodeUcaseResp struct {
	IsRegistered bool
}

//easyjson:json
type UpdateUserReq struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UpdateUserUcase struct {
	Id        int64
	Name      string
	Email     string
	AvatarImg io.Reader
}

type UpdateUserStorage struct {
	Id     int64
	Name   string
	Email  string
	Avatar string
}
