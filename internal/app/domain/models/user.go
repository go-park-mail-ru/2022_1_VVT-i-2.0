package models

import (
	"database/sql"
)

type UserId uint64

type mockUser struct {
	Id     UserId
	Name   string
	Phone  string
	Email  string
	Avatar sql.NullString
}

var User = &mockUser{
	Id: 1,
	Name: "Name",
	Phone: "79166152595",
	Email: "email@email.com",
	Avatar: sql.NullString{String: "avatar"},
}
