package ucase

// import (
// 	"reflect"
// 	"testing"

// 	mockCacher "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/cacher/mock"
// 	mockNatificator "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/notification/mock"
// 	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/auth/mock"
// 	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/auth/models"
// 	"github.com/stretchr/testify/require"
// )

// func TestUserUseCase_Login(t *testing.T) {
// 	mockUserRepo := new(mock.AuthRepo)
// 	mockCacher := new(mockCacher.Memcacher)
// 	mockNotificator := new(mockNatificator.Notificator)
// 	ucase := NewAuthUcase(mockNotificator, mockCacher, mockUserRepo)

// 	req := &models.LoginUcaseReq{
// 		Phone: "79166152595",
// 		Code:  "1234",
// 	}

// 	userResp, err := ucase.Login(req)
// 	require.NoError(t, err)

// 	expectResp := &models.UserDataUcase{
// 		Id:     1,
// 		Name:   "Name",
// 		Phone:  "79999999999",
// 		Email:  "email@mail.com",
// 		Avatar: "avatar.png",
// 	}

// 	if !reflect.DeepEqual(userResp, expectResp) {
// 		t.Errorf("results not match, want %v, have %v", userResp, expectResp)
// 		return
// 	}
// }

// func TestUserUseCase_Register(t *testing.T) {
// 	mockUserRepo := new(mock.AuthRepo)
// 	mockCacher := new(mockCacher.Memcacher)
// 	mockNotificator := new(mockNatificator.Notificator)
// 	ucase := NewAuthUcase(mockNotificator, mockCacher, mockUserRepo)

// 	req := &models.RegisterUcaseReq{
// 		Phone: "79166152595",
// 		Code:  "1234",
// 		Email: "email@mail.com",
// 		Name:  "Name",
// 	}

// 	userResp, err := ucase.Register(req)
// 	require.NoError(t, err)

// 	expectResp := &models.UserDataUcase{
// 		Id:     1,
// 		Name:   "Name",
// 		Phone:  "79999999999",
// 		Email:  "email@mail.com",
// 		Avatar: "",
// 	}

// 	if !reflect.DeepEqual(userResp, expectResp) {
// 		t.Errorf("results not match, want %v, have %v", userResp, expectResp)
// 		return
// 	}
// }

// func TestUserUseCase_SendCode(t *testing.T) {
// 	mockUserRepo := new(mock.AuthRepo)
// 	mockCacher := new(mockCacher.Memcacher)
// 	mockNotificator := new(mockNatificator.Notificator)
// 	ucase := NewAuthUcase(mockNotificator, mockCacher, mockUserRepo)

// 	req := &models.SendCodeUcaseReq{
// 		Phone: "79166152595",
// 	}

// 	resp, err := ucase.SendCode(req)
// 	require.NoError(t, err)

// 	expectResp := models.SendCodeUcaseResp{IsRegistered: true}

// 	if !reflect.DeepEqual(resp, expectResp) {
// 		t.Errorf("results not match, want %v, have %v", resp, expectResp)
// 		return
// 	}
// }
