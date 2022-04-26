package usecase

import (
	"context"
	"fmt"
	"io"

	"github.com/disintegration/imaging"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	cacher "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/cacher"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/notification"
	authProto "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/proto/auth"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/staticManager"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/user"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// TODO: удалить
var LOGIN_CODE string

const (
	codeUpBound          = 10000 // > 0
	codeExpiration int32 = 300   // 5 min
	// avatarSide           = 300
	// avatarSide = 30
)

type UserUsecase struct {
	Notificator   notification.Notificator
	Cacher        cacher.Cacher
	UserRepo      user.Repository
	StaticManager staticManager.FileManager
	Authorizer    authProto.AuthServiceClient
}

func NewUsecase(notificator notification.Notificator, cacher cacher.Cacher, userRepo user.Repository, staticManager staticManager.FileManager, authorizer authProto.AuthServiceClient) *UserUsecase {
	return &UserUsecase{
		Notificator:   notificator,
		Cacher:        cacher,
		UserRepo:      userRepo,
		StaticManager: staticManager,
		Authorizer:    authorizer,
	}
}

func (u *UserUsecase) SendCode(req *models.SendCodeUcaseReq) (models.SendCodeUcaseResp, error) {
	isRegistered, err := u.Authorizer.SendCode(context.Background(), &authProto.SendCodeReq{Phone: req.Phone})
	return models.SendCodeUcaseResp{IsRegistered: isRegistered.IsRegistered}, err
}

func (u *UserUsecase) isCodeCorrect(codeDst string, code string) (bool, error) {
	cachItem, err := u.Cacher.Get(codeDst)

	if err != nil || string(cachItem.Value) != code {
		return false, errors.Wrap(err, "code validation error")
	}
	return true, nil
}

func (u *UserUsecase) Login(req *models.LoginReq) (*models.UserDataUsecase, error) {
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
		Id:     userData.Id,
		Phone:  userData.Phone,
		Name:   userData.Name,
		Email:  userData.Email,
		Avatar: userData.Avatar.String,
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

	userDataStorage, err := u.UserRepo.AddUser(&models.UserAddDataStorage{Phone: req.Phone, Email: req.Email, Name: req.Name})
	if err != nil {
		return nil, errors.Wrap(err, "error adding user to storage")
	}
	return &models.UserDataUsecase{
		Id:    userDataStorage.Id,
		Phone: userDataStorage.Phone,
		Name:  userDataStorage.Name,
		Email: userDataStorage.Email,
	}, nil
}

func (u *UserUsecase) GetUser(id models.UserId) (*models.UserDataUsecase, error) {
	userData, err := u.UserRepo.GetUserById(id)
	if err != nil {
		return nil, errors.Wrapf(err, "error getting user by id %d", id)
	}
	return &models.UserDataUsecase{
		Id:     userData.Id,
		Phone:  userData.Phone,
		Name:   userData.Name,
		Email:  userData.Email,
		Avatar: userData.Avatar.String,
	}, nil
}

func (u *UserUsecase) UpdateUser(updates *models.UpdateUserUsecase) (*models.UserDataUsecase, error) {
	var newAvatarName string
	var err error
	if updates.AvatarImg != nil {
		newAvatarName, err = u.saveNewAvatar(updates.AvatarImg)
		if err != nil {
			return nil, errors.Wrap(err, "error saving new avatar to server")
		}
	}
	updUser, err := u.UserRepo.UpdateUser(&models.UpdateUserStorage{Id: updates.Id, Email: updates.Email, Name: updates.Name, Avatar: newAvatarName})
	if err != nil {
		if newAvatarName != "" {
			u.StaticManager.RemoveAvatar(newAvatarName)
			// os.Remove(avatarPath + newAvatarName)
		}
		return nil, errors.Wrap(err, "error updating user")
	}
	if updUser == nil {
		if newAvatarName != "" {
			u.StaticManager.RemoveAvatar(newAvatarName)
			// os.Remove(avatarPath + newAvatarName)
		}
		updUser, err = u.UserRepo.GetUserById(updates.Id)
		if err != nil {
			return nil, errors.Wrapf(err, "error getting user by id %d", updates.Id)
		}
	}
	return &models.UserDataUsecase{
		Id:     updUser.Id,
		Phone:  updUser.Phone,
		Name:   updUser.Name,
		Email:  updUser.Email,
		Avatar: updUser.Avatar.String,
	}, nil
}

// /*
func (u *UserUsecase) saveNewAvatar(avatar io.Reader) (string, error) {
	avatarImg, err := imaging.Decode(avatar)
	if err != nil {
		fmt.Println(err.Error())
		return "", servErrors.NewError(servErrors.DECODE_IMG, err.Error())
	}

	// if avatarImg.Bounds().Max.X < avatarImg.Bounds().Max.Y {
	// 	avatarImg = imaging.Resize(avatarImg, avatarSide, 0, imaging.Lanczos)
	// } else {
	// 	avatarImg = imaging.Resize(avatarImg, 0, avatarSide, imaging.Lanczos)
	// }

	var avatarName string
	for i := 0; i < 10; i++ {
		avatarName = newAvatarName() + ".png"
		if u.StaticManager.IsNotSuchAvatarExist(avatarName) {
			break
		}
		avatarName = ""
	}
	if avatarName == "" {
		return "", servErrors.NewError(servErrors.CANT_CREATE_AVATAR_NAME, "")
	}

	// err = imaging.Save(avatarImg, staticPath+avatarPath+avatarName)
	err = u.StaticManager.SafeAvatar(avatarImg, avatarName)
	fmt.Println(err)
	fmt.Println(avatarName)
	if err != nil {
		return "", servErrors.NewError(servErrors.CANT_SAVE_AVATAR, err.Error())
	}

	return avatarName, nil
}

func newAvatarName() string {
	fname, _ := uuid.NewUUID()
	return fname.String()
}
