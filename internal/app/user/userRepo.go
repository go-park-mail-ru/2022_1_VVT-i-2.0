package user

import "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"

type Repository interface {
	GetUserById(id models.UserId) (*models.UserDataRepo, error)
	UpdateUser(updUser *models.UpdateUserStorage) (*models.UserDataRepo, error)
}
