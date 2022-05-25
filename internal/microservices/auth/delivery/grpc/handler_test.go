package grpc

import (
	"context"
	"reflect"
	"testing"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/auth/mock"
	proto "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/auth/proto"
	"github.com/stretchr/testify/assert"
)

func TestAuthHandler_Login(t *testing.T) {
	mockUcase := new(mock.AuthUcase)
	handler := NewAuthHandler(mockUcase)

	req := proto.LoginReq{
		Phone: "79999999999",
		Code:  "1234",
	}

	userResp, err := handler.Login(context.Background(), &req)
	assert.NoError(t, err)

	expectResp := &proto.LoginResp{
		Id:      1,
		Name:    "Name",
		Phone:   "79999999999",
		Email:   "mail.ru",
		Avatar:  "avatar",
		Address: "адрес",
	}

	if !reflect.DeepEqual(userResp, expectResp) {
		t.Errorf("results not match, want %v, have %v", userResp, expectResp)
		return
	}
}

func TestAuthHandler_Register(t *testing.T) {
	mockUcase := new(mock.AuthUcase)
	handler := NewAuthHandler(mockUcase)

	req := proto.RegisterReq{
		Phone: "79166152595",
		Code:  "1234",
		Email: "email@mail.com",
		Name:  "Name",
	}

	userResp, err := handler.Register(context.Background(), &req)
	assert.NoError(t, err)

	expectResp := &proto.UserData{
		Id:     1,
		Name:   "Name",
		Phone:  "79999999999",
		Email:  "mail.ru",
		Avatar: "avatar",
	}

	if !reflect.DeepEqual(userResp, expectResp) {
		t.Errorf("results not match, want %v, have %v", userResp, expectResp)
		return
	}
}

func TestAuthHandler_SendCode(t *testing.T) {
	mockUcase := new(mock.AuthUcase)
	handler := NewAuthHandler(mockUcase)

	req := proto.SendCodeReq{
		Phone: "79166152595",
	}

	userResp, err := handler.SendCode(context.Background(), &req)
	assert.NoError(t, err)

	expectResp := &proto.IsRegistered{
		IsRegistered: false,
	}

	if !reflect.DeepEqual(userResp, expectResp) {
		t.Errorf("results not match, want %v, have %v", userResp, expectResp)
		return
	}
}

func TestAuthHandler_LoginErr(t *testing.T) {
	mockUcase := new(mock.AuthUcaseErr)
	handler := NewAuthHandler(mockUcase)

	req := proto.LoginReq{
		Phone: "79999999999",
		Code:  "1234",
	}

	_, err := handler.Login(context.Background(), &req)
	assert.Error(t, err)
}

func TestAuthHandler_RegisterErr(t *testing.T) {
	mockUcase := new(mock.AuthUcaseErr)
	handler := NewAuthHandler(mockUcase)

	req := proto.RegisterReq{
		Phone: "79166152595",
		Code:  "1234",
		Email: "email@mail.com",
		Name:  "Name",
	}

	_, err := handler.Register(context.Background(), &req)
	assert.Error(t, err)
}

func TestAuthHandler_SendCodeErr(t *testing.T) {
	mockUcase := new(mock.AuthUcaseErr)
	handler := NewAuthHandler(mockUcase)

	req := proto.SendCodeReq{
		Phone: "79166152595",
	}

	_, err := handler.SendCode(context.Background(), &req)

	assert.Error(t, err)
}
