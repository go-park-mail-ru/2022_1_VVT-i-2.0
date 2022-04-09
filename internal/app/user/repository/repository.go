package repository

import "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"

type UserRepo struct {
}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (r *UserRepo) GetUserByPhone(phone string) (*models.UserDataStorage, error) {
	return &models.UserDataStorage{Id: 1, Phone: "79000000000", Name: "UserByPhone", Email: "natali-skv@mail.ru"}, nil
}

func (r *UserRepo) AddUser(newUser *models.UserAddDataStorage) (*models.UserDataStorage, error) {
	return &models.UserDataStorage{Id: 2, Phone: newUser.Phone, Name: newUser.Name, Email: newUser.Email}, nil
}

func (r *UserRepo) GetUserById(id models.UserId) (*models.UserDataStorage, error) {
	return &models.UserDataStorage{Id: 2, Phone: "79000000000", Name: "Kati", Email: "natali-skv@mail.ru"}, nil
}

func (r *UserRepo) UpdateData(newData models.UserDataUpdateReq) error {
	return nil
}

func (r *UserRepo) HasUserByPhone(phone string) (bool, error) {
	return true, nil
}
