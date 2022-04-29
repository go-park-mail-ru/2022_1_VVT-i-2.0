package usecase

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/domain/interfaces"
	data "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/domain/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/stretchr/testify/require"
	"io"
	"reflect"
	"testing"
)

func TestUserUseCase_Login(t *testing.T) {
	mockUserRepo := new(interfaces.UserRepository)
	mockCacher := new(interfaces.Memcacher)
	useCase := UserUsecase{
		Cacher: 	mockCacher,
		UserRepo: 	mockUserRepo,
	}

	addUser := &models.LoginReq{
		Phone: "79166152595",
		Code: "1234",
	}

	restData, err := useCase.Login(addUser)
	require.NoError(t, err)

	mockUser := &models.UserDataUsecase{
		Id: 	models.UserId(data.User.Id),
		Name:   data.User.Name,
		Phone: 	data.User.Phone,
		Email: 	data.User.Email,
		Avatar: data.User.Avatar.String,
	}

	if !reflect.DeepEqual(restData, mockUser) {
		t.Errorf("results not match, want %v, have %v", restData, mockUser)
		return
	}
}

func TestUserUsecase_Register(t *testing.T) {
	mockUserRepo := new(interfaces.UserRepository)
	mockCacher := new(interfaces.Memcacher)
	useCase := UserUsecase{
		Cacher: 	mockCacher,
		UserRepo: 	mockUserRepo,
	}

	user := &models.RegisterReq{
		Phone: 	data.User.Phone,
		Code: "1234",
		Name:   data.User.Name,
		Email: 	data.User.Email,
	}

	restData, err := useCase.Register(user)
	require.NoError(t, err)

	mockUser := &models.UserDataUsecase{
		Id: models.UserId(data.User.Id),
		Name:   data.User.Name,
		Phone: data.User.Phone,
		Email: data.User.Email,
		Avatar: "",
	}

	if !reflect.DeepEqual(restData, mockUser) {
		t.Errorf("results not match, want %v, have %v", restData, mockUser)
		return
	}
}

func TestUserUsecase_GetUser(t *testing.T) {
	mockUserRepo := new(interfaces.UserRepository)
	mockCacher := new(interfaces.Memcacher)
	useCase := UserUsecase{
		Cacher: 	mockCacher,
		UserRepo: 	mockUserRepo,
	}

	user := models.UserId(1)

	restData, err := useCase.GetUser(user)
	require.NoError(t, err)

	mockUser := &models.UserDataUsecase{
		Id: models.UserId(data.User.Id),
		Name:   data.User.Name,
		Phone: data.User.Phone,
		Email: data.User.Email,
		Avatar:  data.User.Avatar.String,
	}

	if !reflect.DeepEqual(restData, mockUser) {
		t.Errorf("results not match, want %v, have %v", restData, mockUser)
		return
	}
}

func TestUserUsecase_UpdateUser(t *testing.T) {
	mockUserRepo := new(interfaces.UserRepository)
	mockCacher := new(interfaces.Memcacher)
	useCase := UserUsecase{
		Cacher: 	mockCacher,
		UserRepo: 	mockUserRepo,
	}

	user := &models.UpdateUserUsecase{
		Id: 	models.UserId(data.User.Id),
		Name:   data.User.Name,
		Email:  data.User.Email,
		AvatarImg: io.Reader(nil),
	}

	restData, err := useCase.UpdateUser(user)
	require.NoError(t, err)

	mockUser := &models.UserDataUsecase{
		Id: models.UserId(data.User.Id),
		Name:   data.User.Name,
		Phone: data.User.Phone,
		Email: data.User.Email,
		Avatar:  data.User.Avatar.String,
	}

	if !reflect.DeepEqual(restData, mockUser) {
		t.Errorf("results not match, want %v, have %v", restData, mockUser)
		return
	}
}

//func TestUserUsecase_saveNewAvatar(t *testing.T) {
//	mockUserRepo := new(interfaces.UserRepository)
//	mockCacher := new(interfaces.Memcacher)
//	useCase := UserUsecase{
//		Cacher: 	mockCacher,
//		UserRepo: 	mockUserRepo,
//	}
//
//	restData, err := useCase.saveNewAvatar()
//	require.NoError(t, err)
//
//	mockUser := &models.UserDataUsecase{
//		Id: models.UserId(data.User.Id),
//		Name:   data.User.Name,
//		Phone: data.User.Phone,
//		Email: data.User.Email,
//		Avatar:  data.User.Avatar.String,
//	}
//
//	if !reflect.DeepEqual(restData, mockUser) {
//		t.Errorf("results not match, want %v, have %v", restData, mockUser)
//		return
//	}
//}