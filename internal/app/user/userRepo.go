package user

import "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"

type Repository interface {
	// AddUser(*models.UserDataStorage) (*models.UserDataStorage, error)
	// TODO  не знаю как назвать модели правильно
	GetUserByPhone(phone string) (*models.UserDataStorage, error)
	GetUserById(id models.UserId) (*models.UserDataStorage, error)
	UpdateData(newData models.UserDataUpdateReq) error
	HasUserByPhone(phone string) (bool, error)
}
