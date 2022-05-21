package ucase

// import (
// 	"io"
// 	"reflect"
// 	"testing"

// 	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
// 	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/staticManager/localStaticManager"
// 	mockAuthCli "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/auth/mock"
// 	"github.com/stretchr/testify/assert"

// 	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/user/mock"
// )

// func TestUserUseCase_SendCode(t *testing.T) {
// 	mockUserRepo := new(mock.UserRepository)
// 	mockAuthCli := new(mockAuthCli.GrpcAuthHandler)
// 	staticManager := localStaticManager.NewLocalFileManager("", "")
// 	useCase := NewUcase(mockUserRepo, staticManager, mockAuthCli)

// 	sendCodeReq := &models.SendCodeUcaseReq{
// 		Phone: "79999999999",
// 	}

// 	resp, err := useCase.SendCode(sendCodeReq)
// 	assert.NoError(t, err)

// 	mockUser := models.SendCodeUcaseResp{
// 		IsRegistered: false,
// 	}

// 	if !reflect.DeepEqual(resp, mockUser) {
// 		t.Errorf("results not match, want %v, have %v", resp, mockUser)
// 		return
// 	}
// }

// func TestUserUseCase_Login(t *testing.T) {
// 	mockUserRepo := new(mock.UserRepository)
// 	mockAuthCli := new(mockAuthCli.GrpcAuthHandler)
// 	staticManager := localStaticManager.NewLocalFileManager("", "")
// 	useCase := NewUcase(mockUserRepo, staticManager, mockAuthCli)

// 	addUser := &models.LoginUcaseReq{
// 		Phone: "79999999999",
// 		Code:  "1234",
// 	}

// 	restData, err := useCase.Login(addUser)
// 	assert.NoError(t, err)

// 	mockUser := &models.UserDataUcase{
// 		Id:     1,
// 		Name:   "Name",
// 		Phone:  "79999999999",
// 		Email:  "email@mail.com",
// 		Avatar: "avatar.png",
// 	}

// 	if !reflect.DeepEqual(restData, mockUser) {
// 		t.Errorf("results not match, want %v, have %v", restData, mockUser)
// 		return
// 	}
// }

// func TestUserUsecase_Register(t *testing.T) {
// 	mockUserRepo := new(mock.UserRepository)
// 	mockAuthCli := new(mockAuthCli.GrpcAuthHandler)
// 	staticManager := localStaticManager.NewLocalFileManager("", "")
// 	useCase := NewUcase(mockUserRepo, staticManager, mockAuthCli)

// 	user := &models.RegisterUcaseReq{
// 		Name:  "Name",
// 		Phone: "79999999999",
// 		Email: "email@mail.com",
// 	}

// 	restData, err := useCase.Register(user)
// 	assert.NoError(t, err)

// 	mockUser := &models.UserDataUcase{
// 		Id:     1,
// 		Name:   "Name",
// 		Phone:  "79999999999",
// 		Email:  "email@mail.com",
// 		Avatar: "avatar.png",
// 	}

// 	if !reflect.DeepEqual(restData, mockUser) {
// 		t.Errorf("results not match, want %v, have %v", restData, mockUser)
// 		return
// 	}
// }

// func TestUserUsecase_GetUser(t *testing.T) {
// 	mockUserRepo := new(mock.UserRepository)
// 	mockAuthCli := new(mockAuthCli.GrpcAuthHandler)
// 	staticManager := localStaticManager.NewLocalFileManager("", "")
// 	useCase := NewUcase(mockUserRepo, staticManager, mockAuthCli)

// 	user := models.UserId(1)

// 	restData, err := useCase.GetUser(user)
// 	assert.NoError(t, err)

// 	mockUser := &models.UserDataUcase{
// 		Id:     1,
// 		Name:   "Name",
// 		Phone:  "79999999999",
// 		Email:  "email@mail.com",
// 		Avatar: "avatar.png",
// 	}

// 	if !reflect.DeepEqual(restData, mockUser) {
// 		t.Errorf("results not match, want %v, have %v", restData, mockUser)
// 		return
// 	}
// }

// func TestUserUsecase_UpdateUser(t *testing.T) {
// 	mockUserRepo := new(mock.UserRepository)
// 	mockAuthCli := new(mockAuthCli.GrpcAuthHandler)
// 	staticManager := localStaticManager.NewLocalFileManager("", "")
// 	useCase := NewUcase(mockUserRepo, staticManager, mockAuthCli)

// 	user := &models.UpdateUserUcase{
// 		Id:        1,
// 		Name:      "Name",
// 		Email:     "email@mail.com",
// 		AvatarImg: io.Reader(nil),
// 	}

// 	restData, err := useCase.UpdateUser(user)
// 	assert.NoError(t, err)

// 	mockUser := &models.UserDataUcase{
// 		Id:     1,
// 		Name:   "Name",
// 		Phone:  "79999999999",
// 		Email:  "email@mail.com",
// 		Avatar: "avatar.png",
// 	}

// 	if !reflect.DeepEqual(restData, mockUser) {
// 		t.Errorf("results not match, want %v, have %v", restData, mockUser)
// 		return
// 	}
// }
