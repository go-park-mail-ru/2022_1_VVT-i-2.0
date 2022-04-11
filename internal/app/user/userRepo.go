package user

import "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"

type Repository interface {
	AddUser(*models.UserAddDataStorage) (models.UserId, error)
	GetUserByPhone(phone string) (*models.UserDataStorage, error)
	GetUserById(id models.UserId) (*models.UserDataStorage, error)
	UpdateUser(updUser *models.UpdateUser) (*models.UserDataStorage, error)
	HasUserByPhone(phone string) (bool, error)
}
