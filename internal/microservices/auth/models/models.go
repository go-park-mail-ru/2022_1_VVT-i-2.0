package models

import (
	"database/sql"
)

type UserId uint64

type UserDataRepo struct {
	Id     UserId
	Name   string
	Phone  string
	Email  string
	Avatar sql.NullString
}

type AddUserRepoReq struct {
	Name   string
	Phone  string
	Email  string
	Avatar string
}

type UserDataUcase struct {
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
	Name   string
	Phone  string
	Email  string
	Avatar string
}

type LoginUcaseReq struct {
	Phone string
	Code  string
}

type LogitUcaseResp struct {
	Id     UserId
	Name   string
	Phone  string
	Email  string
	Avatar string
	Addres string
}

type RegisterUcaseReq struct {
	Phone string
	Code  string
	Name  string
	Email string
}

type SendCodeUcaseReq struct {
	Phone string
}

type SendCodeUcaseResp struct {
	IsRegistered bool
}

type UserByPhoneRepoReq struct {
	Phone string
}

type HasSuchUserRepoResp struct {
	IsRegistered bool
}

type GetTopUserAddrRepoReq struct {
	UserId int64
}

type GetTopUserAddrRepoResp struct {
	Address string
}
