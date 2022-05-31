package mock

import (
	"database/sql"
	"errors"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/stretchr/testify/mock"
)

type UserRepository struct {
	mock.Mock
}

func (u *UserRepository) AddUser(storage *models.UserAddDataStorage) (*models.UserDataRepo, error) {
	if storage == nil {
		return nil, servErrors.NewError(servErrors.DB_ERROR, "")
	}
	return &models.UserDataRepo{
		Id:     1,
		Name:   "Name",
		Phone:  "79999999999",
		Email:  "email@mail.com",
		Avatar: sql.NullString{String: "avatar.png"}}, nil
}

func (u *UserRepository) GetUserByPhone(phone string) (*models.UserDataRepo, error) {
	if phone == "" {
		return nil, servErrors.NewError(servErrors.DB_ERROR, "")
	}
	return &models.UserDataRepo{
		Id:     1,
		Name:   "Name",
		Phone:  "79999999999",
		Email:  "email@mail.com",
		Avatar: sql.NullString{String: "avatar.png"}}, nil
}

func (u *UserRepository) GetUserById(id int64) (*models.UserDataRepo, error) {
	if id == 0 {
		return nil, servErrors.NewError(servErrors.DB_ERROR, "")
	}
	return &models.UserDataRepo{
		Id:     1,
		Name:   "Name",
		Phone:  "79999999999",
		Email:  "email@mail.com",
		Avatar: sql.NullString{String: "avatar.png"}}, nil
}

func (u *UserRepository) UpdateUser(updUser *models.UpdateUserStorage) (*models.UserDataRepo, error) {
	if updUser == nil {
		return nil, servErrors.NewError(servErrors.DB_ERROR, "")
	}
	return &models.UserDataRepo{
		Id:     1,
		Name:   "Name",
		Phone:  "79999999999",
		Email:  "email@mail.com",
		Avatar: sql.NullString{String: "avatar.png"}}, nil
}

func (u *UserRepository) HasUserByPhone(phone string) (bool, error) {
	if phone == "" {
		return false, servErrors.NewError(servErrors.DB_ERROR, "")
	}
	return true, nil
}

type UserRepositoryErr struct {
	mock.Mock
}

func (u *UserRepositoryErr) AddUser(storage *models.UserAddDataStorage) (*models.UserDataRepo, error) {
	if storage == nil {
		return nil, servErrors.NewError(servErrors.DB_ERROR, "")
	}
	return nil, errors.New("unknown error")
}

func (u *UserRepositoryErr) GetUserByPhone(phone string) (*models.UserDataRepo, error) {
	if phone == "" {
		return nil, servErrors.NewError(servErrors.DB_ERROR, "")
	}

	return nil, errors.New("unknown error")
}

func (u *UserRepositoryErr) GetUserById(id int64) (*models.UserDataRepo, error) {
	if id == 0 {
		return nil, servErrors.NewError(servErrors.DB_ERROR, "")
	}

	return nil, errors.New("unknown error")
}

func (u *UserRepositoryErr) UpdateUser(updUser *models.UpdateUserStorage) (*models.UserDataRepo, error) {
	if updUser == nil {
		return nil, servErrors.NewError(servErrors.DB_ERROR, "")
	}

	return nil, errors.New("unknown error")
}

func (u *UserRepositoryErr) HasUserByPhone(phone string) (bool, error) {
	if phone == "" {
		return false, servErrors.NewError(servErrors.DB_ERROR, "")
	}
	return false, errors.New("unknown error")
}
