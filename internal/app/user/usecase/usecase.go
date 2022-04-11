package usecase

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	cacher "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/cacher"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/notification"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/user"
	"github.com/pkg/errors"
)

// TODO: удалить
var LOGIN_CODE string

const (
	codeUpBound          = 10000 // > 0
	codeExpiration int32 = 300   // 5 min
	codeSendMsg          = "Ваш код для входа в Foobrinto: "
	msgEncoding          = "unicode"
)

type UserUsecase struct {
	Notificator notification.Notificator
	Cacher      cacher.Cacher
	UserRepo    user.Repository
}

func NewUsecase(notificator notification.Notificator, cacher cacher.Cacher, userRepo user.Repository) *UserUsecase {
	return &UserUsecase{
		Notificator: notificator,
		Cacher:      cacher,
		UserRepo:    userRepo,
	}
}

func generateLoginCode() string {
	randNum, _ := rand.Int(rand.Reader, big.NewInt(codeUpBound))
	return strconv.Itoa(int(randNum.Int64()) + codeUpBound)[1:]
}

func (u *UserUsecase) SendCode(req *models.SendCodeReq) (bool, error) {
	loginCode := generateLoginCode()
	LOGIN_CODE = loginCode //TODO: удалить
	fmt.Printf("~~~~~~~code: %s ~~~~~~~~\n", loginCode)
	err := u.Cacher.Set(cacher.NewItem(req.Phone, []byte(loginCode), codeExpiration))
	if err != nil {
		return false, errors.Wrap(err, "error saving [auth code destination]-code item to cach")
	}

	// err = u.Notificator.SendCode(req.Phone, loginCode)
	// if err != nil {
	// 	return false, errors.Wrap(err, "error sending message e with code to auth code destination")
	// }

	hasSuchUser, err := u.UserRepo.HasUserByPhone(req.Phone)
	if err != nil {
		return false, errors.Wrap(err, "error finding out if there is such user in database")
	}
	return hasSuchUser, nil
}

func (u *UserUsecase) isCodeCorrect(codeDst string, code string) (bool, error) {
	cachItem, err := u.Cacher.Get(codeDst)

	if err != nil || string(cachItem.Value) != code {
		return false, errors.Wrap(err, "code validation error")
	}
	return true, nil
}

func (u *UserUsecase) Login(req *models.LoginRequest) (*models.UserDataUsecase, error) {
	isCorrect, err := u.isCodeCorrect(req.Phone, req.Code)
	if err != nil {
		return nil, errors.Wrap(err, "code check failed")
	}
	if !isCorrect {
		return nil, servErrors.NewError(servErrors.WRONG_AUTH_CODE, servErrors.WRONG_AUTH_CODE_DESCR)
	}
	userData, err := u.UserRepo.GetUserByPhone(req.Phone)
	if err != nil {
		return nil, errors.Wrap(err, "error getting user by phone")
	}
	return &models.UserDataUsecase{
		Id:    userData.Id,
		Phone: userData.Phone,
		Name:  userData.Name,
		Email: userData.Email,
	}, nil
}

func (u *UserUsecase) Register(req *models.RegisterReq) (*models.UserDataUsecase, error) {
	isCorrect, err := u.isCodeCorrect(req.Phone, req.Code)
	if err != nil {
		return nil, errors.Wrap(err, "code check failed")
	}
	if !isCorrect {
		return nil, servErrors.NewError(servErrors.WRONG_AUTH_CODE, servErrors.WRONG_AUTH_CODE_DESCR)
	}

	id, err := u.UserRepo.AddUser(&models.UserAddDataStorage{Phone: req.Phone, Email: req.Email, Name: req.Name})
	if err != nil {
		return nil, errors.Wrap(err, "error adding user to storage")
	}
	return &models.UserDataUsecase{
		Id:    id,
		Phone: req.Phone,
		Name:  req.Name,
		Email: req.Email,
	}, nil
}

func (u *UserUsecase) GetUser(id models.UserId) (*models.UserDataUsecase, error) {
	userData, err := u.UserRepo.GetUserById(id)
	if err != nil {
		return nil, errors.Wrapf(err, "error getting user by id %d", id)
	}
	return &models.UserDataUsecase{
		Id:    userData.Id,
		Phone: userData.Phone,
		Name:  userData.Name,
		Email: userData.Email,
	}, nil
}

func (u *UserUsecase) UpdateUser(updates *models.UpdateUser) (*models.UserDataUsecase, error) {
	updUser, err := u.UserRepo.UpdateUser(updates)
	if err != nil {
		return nil, errors.Wrap(err, "error updating user")
	}
	return &models.UserDataUsecase{
		Id:    updUser.Id,
		Phone: updUser.Phone,
		Name:  updUser.Name,
		Email: updUser.Email,
	}, nil
}
