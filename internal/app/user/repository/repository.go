package repository

import "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"

type UserRepo struct {
}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (r *UserRepo) GetUserByPhone(phone string) (*models.UserDataStorage, error) {
	return &models.UserDataStorage{Phone: "79000000000", Name: "Kati", Email: "natali-skv@mail.ru"}, nil
}

func (r *UserRepo) GetUserById(id models.UserId) (*models.UserDataStorage, error) {
	return &models.UserDataStorage{Phone: "79000000000", Name: "Kati", Email: "natali-skv@mail.ru"}, nil
}

func (r *UserRepo) UpdateData(newData models.UserDataUpdateReq) error {
	return nil
}

func (r *UserRepo) HasUserByPhone(phone string) (bool, error) {
	return true, nil
}
