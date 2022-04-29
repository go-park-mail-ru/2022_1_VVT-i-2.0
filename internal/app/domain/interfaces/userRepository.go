package interfaces

import (
	data "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/domain/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/stretchr/testify/mock"
)

type UserRepository struct {
	mock.Mock
}

func (u UserRepository) AddUser(storage *models.UserAddDataStorage) (*models.UserDataStorage, error) {
	if storage == nil {
		return nil, servErrors.NewError(servErrors.TEST_ERROR, "")
	}
	mockUser := &models.UserDataStorage{
		Id: models.UserId(data.User.Id),
		Name:   data.User.Name,
		Phone: data.User.Phone,
		Email: data.User.Email,
		Avatar: data.User.Avatar,
	}
	return mockUser, nil
}

func (u UserRepository) GetUserByPhone(phone string) (*models.UserDataStorage, error) {
	if phone == "" {
		return nil, servErrors.NewError(servErrors.TEST_ERROR, "")
	}
	mockUser := &models.UserDataStorage{
		Id: models.UserId(data.User.Id),
		Name:   data.User.Name,
		Phone: data.User.Phone,
		Email: data.User.Email,
		Avatar: data.User.Avatar,
	}
	return mockUser, nil
}

func (u UserRepository) GetUserById(id models.UserId) (*models.UserDataStorage, error) {
	if id == 0 {
		return nil, servErrors.NewError(servErrors.TEST_ERROR, "")
	}
	mockUser := &models.UserDataStorage{
		Id: models.UserId(data.User.Id),
		Name:   data.User.Name,
		Phone: data.User.Phone,
		Email: data.User.Email,
		Avatar: data.User.Avatar,
	}
	return mockUser, nil
}

func (u UserRepository) UpdateUser(updUser *models.UpdateUserStorage) (*models.UserDataStorage, error) {
	if updUser == nil {
		return nil, servErrors.NewError(servErrors.TEST_ERROR, "")
	}
	mockUser := &models.UserDataStorage{
		Id: models.UserId(data.User.Id),
		Name:   data.User.Name,
		Phone: data.User.Phone,
		Email: data.User.Email,
		Avatar: data.User.Avatar,
	}
	return mockUser, nil
}

func (u UserRepository) HasUserByPhone(phone string) (bool, error) {
	if phone == "" {
		return false, servErrors.NewError(servErrors.TEST_ERROR, "")
	}
	return true, nil
}
