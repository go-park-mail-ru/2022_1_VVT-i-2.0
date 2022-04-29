package interfaces

import (
	data "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/domain/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/stretchr/testify/mock"
)

type UserUseCase struct {
	mock.Mock
}

func (u *UserUseCase) Login(req *models.LoginReq) (*models.UserDataUsecase, error) {
	if req == nil {
		return nil, servErrors.NewError(servErrors.TEST_ERROR, "")
	}
	mockUser := &models.UserDataUsecase{
		Id: models.UserId(data.User.Id),
		Name:   data.User.Name,
		Phone: data.User.Phone,
		Email: data.User.Email,
		Avatar: data.User.Avatar.String,
	}
	return mockUser, nil
}

func (u *UserUseCase) Register(req *models.RegisterReq) (*models.UserDataUsecase, error) {
	if req == nil {
		return nil, servErrors.NewError(servErrors.TEST_ERROR, "")
	}
	mockUser := &models.UserDataUsecase{
		Id: models.UserId(data.User.Id),
		Name:   data.User.Name,
		Phone: data.User.Phone,
		Email: data.User.Email,
		Avatar: data.User.Avatar.String,
	}
	return mockUser, nil
}

func (u *UserUseCase) SendCode(req *models.SendCodeReq) (bool, error) {
	if req == nil {
		return false, servErrors.NewError(servErrors.TEST_ERROR, "")
	}
	return true, nil
}

func (u *UserUseCase) GetUser(id models.UserId) (*models.UserDataUsecase, error) {
	panic("implement me")
}

func (u *UserUseCase) UpdateUser(req *models.UpdateUserUsecase) (*models.UserDataUsecase, error) {
	panic("implement me")
}


