package user

import "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"

type Repository interface {
	GetUserByPhone(phone string) (*models.UserDataRepo, error)
	GetUserById(id models.UserId) (*models.UserDataRepo, error)
	UpdateUser(updUser *models.UpdateUserStorage) (*models.UserDataRepo, error)
	HasUserByPhone(phone string) (bool, error)
}
