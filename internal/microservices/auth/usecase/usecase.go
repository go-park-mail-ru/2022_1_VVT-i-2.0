package ucase

import (
	"crypto/rand"
	"math/big"
	"strconv"

	cacher "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/cacher"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/notification"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/auth"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/auth/models"
	"github.com/pkg/errors"
)

const (
	codeUpBound          = 10000 // > 0
	codeExpiration int32 = 300   // 5 min
)

type AuthUcase struct {
	Notificator notification.Notificator
	Cacher      cacher.Cacher
	AuthRepo    auth.Repository
}

func NewAuthUcase(notificator notification.Notificator, cacher cacher.Cacher, authRepo auth.Repository) *AuthUcase {
	return &AuthUcase{
		Notificator: notificator,
		Cacher:      cacher,
		AuthRepo:    authRepo,
	}
}

func generateLoginCode() string {
	randNum, _ := rand.Int(rand.Reader, big.NewInt(codeUpBound))
	return strconv.Itoa(int(randNum.Int64()) + codeUpBound)[1:]
}

func (u *AuthUcase) SendCode(req *models.SendCodeUcaseReq) (models.SendCodeUcaseResp, error) {
	loginCode := generateLoginCode()

	err := u.Notificator.SendCode(req.Phone, loginCode)

	if err != nil {
		cause := servErrors.ErrorAs(err)
		if cause == nil || cause.Code != servErrors.FLASHCALL_PHONE_ALREADY_IN_QUEUE {
			return models.SendCodeUcaseResp{}, errors.Wrap(err, "error sending message e with code to auth code destination")
		}
	} else {
		err = u.Cacher.Set(cacher.NewItem(req.Phone, []byte(loginCode), codeExpiration))
		if err != nil {
			return models.SendCodeUcaseResp{}, errors.Wrap(err, "error saving [auth code destination]-code item to cach")
		}
	}

	hasSuchUser, err := u.AuthRepo.HasUserByPhone(models.UserByPhoneRepoReq{Phone: req.Phone})
	if err != nil {
		return models.SendCodeUcaseResp{}, errors.Wrap(err, "error finding out if there is such user in database")
	}
	return models.SendCodeUcaseResp(hasSuchUser), nil
}

func (u *AuthUcase) isCodeCorrect(codeDst string, code string) (bool, error) {
	// cachItem, err := u.Cacher.Get(codeDst)

	// if err != nil || string(cachItem.Value) != code {
	// 	return false, errors.Wrap(err, "code validation error")
	// }
	return true, nil
}

func (u *AuthUcase) Login(req *models.LoginUcaseReq) (*models.LogitUcaseResp, error) {
	isCorrect, err := u.isCodeCorrect(req.Phone, req.Code)
	if err != nil {
		return nil, errors.Wrap(err, "code check failed")
	}
	if !isCorrect {
		return nil, servErrors.NewError(servErrors.WRONG_AUTH_CODE, servErrors.WRONG_AUTH_CODE_DESCR)
	}
	userData, err := u.AuthRepo.GetUserByPhone(models.UserByPhoneRepoReq{Phone: req.Phone})
	if err != nil {
		return nil, errors.Wrap(err, "error getting user by phone")
	}
	topAddressRepoResp, err := u.AuthRepo.GetTopUserAddr(&models.GetTopUserAddrRepoReq{UserId: int64(userData.Id)})
	if err != nil {
		return nil, errors.Wrap(err, "error getting user address")
	}
	return &models.LogitUcaseResp{
		Id:     userData.Id,
		Phone:  userData.Phone,
		Name:   userData.Name,
		Email:  userData.Email,
		Avatar: userData.Avatar.String,
		Addres: topAddressRepoResp.Address,
	}, nil
}

func (u *AuthUcase) Register(req *models.RegisterUcaseReq) (*models.UserDataUcase, error) {
	isCorrect, err := u.isCodeCorrect(req.Phone, req.Code)
	if err != nil {
		return nil, errors.Wrap(err, "code check failed")
	}
	if !isCorrect {
		return nil, servErrors.NewError(servErrors.WRONG_AUTH_CODE, servErrors.WRONG_AUTH_CODE_DESCR)
	}

	userDataStorage, err := u.AuthRepo.AddUser(&models.AddUserRepoReq{Phone: req.Phone, Email: req.Email, Name: req.Name})
	if err != nil {
		return nil, errors.Wrap(err, "error adding user to storage")
	}

	return &models.UserDataUcase{
		Id:    userDataStorage.Id,
		Phone: userDataStorage.Phone,
		Name:  userDataStorage.Name,
		Email: userDataStorage.Email,
	}, nil
}
