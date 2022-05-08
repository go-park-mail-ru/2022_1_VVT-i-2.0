package ucase

import (
	"context"
	"io"

	"github.com/disintegration/imaging"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/staticManager"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/user"
	authProto "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/auth/proto"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"google.golang.org/grpc/status"
)

// TODO: удалить
var LOGIN_CODE string

// avatarSide           = 300
// avatarSide = 30
// )

type UserUcase struct {
	UserRepo      user.Repository
	StaticManager staticManager.FileManager
	Authorizer    authProto.AuthServiceClient
}

func NewUcase(userRepo user.Repository, staticManager staticManager.FileManager, authorizer authProto.AuthServiceClient) *UserUcase {
	return &UserUcase{
		UserRepo:      userRepo,
		StaticManager: staticManager,
		Authorizer:    authorizer,
	}
}

func (u *UserUcase) SendCode(req *models.SendCodeUcaseReq) (models.SendCodeUcaseResp, error) {
	isRegistered, err := u.Authorizer.SendCode(context.Background(), &authProto.SendCodeReq{Phone: req.Phone})
	if err != nil {
		return models.SendCodeUcaseResp{IsRegistered: false}, servErrors.NewError(int(status.Code(err)), err.Error())
	}
	return models.SendCodeUcaseResp{IsRegistered: isRegistered.IsRegistered}, err
}

func (u *UserUcase) Register(req *models.RegisterUcaseReq) (*models.UserDataUcase, error) {
	userData, err := u.Authorizer.Register(context.Background(), &authProto.RegisterReq{Phone: req.Phone, Code: req.Code, Name: req.Name, Email: req.Email})
	if err != nil {
		return nil, servErrors.NewError(int(status.Code(err)), err.Error())
	}
	return &models.UserDataUcase{Id: models.UserId(userData.Id), Phone: userData.Phone, Name: userData.Name, Email: userData.Email, Avatar: userData.Avatar}, err
}

func (u *UserUcase) Login(req *models.LoginUcaseReq) (*models.UserDataUcase, error) {
	userData, err := u.Authorizer.Login(context.Background(), &authProto.LoginReq{Phone: req.Phone, Code: req.Code})
	if err != nil {
		return nil, servErrors.NewError(int(status.Code(err)), err.Error())
	}
	return &models.UserDataUcase{Id: models.UserId(userData.Id), Phone: userData.Phone, Name: userData.Name, Email: userData.Email, Avatar: userData.Avatar}, err
}

func (u *UserUcase) GetUser(id models.UserId) (*models.UserDataUcase, error) {
	userData, err := u.UserRepo.GetUserById(id)
	if err != nil {
		return nil, errors.Wrapf(err, "error getting user by id %d", id)
	}
	return &models.UserDataUcase{
		Id:     userData.Id,
		Phone:  userData.Phone,
		Name:   userData.Name,
		Email:  userData.Email,
		Avatar: userData.Avatar.String,
	}, nil
}

func (u *UserUcase) UpdateUser(updates *models.UpdateUserUcase) (*models.UserDataUcase, error) {
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
		}
		return nil, errors.Wrap(err, "error updating user")
	}
	if updUser == nil {
		if newAvatarName != "" {
			u.StaticManager.RemoveAvatar(newAvatarName)
		}
		updUser, err = u.UserRepo.GetUserById(updates.Id)
		if err != nil {
			return nil, errors.Wrapf(err, "error getting user by id %d", updates.Id)
		}
	}
	return &models.UserDataUcase{
		Id:     updUser.Id,
		Phone:  updUser.Phone,
		Name:   updUser.Name,
		Email:  updUser.Email,
		Avatar: updUser.Avatar.String,
	}, nil
}

func (u *UserUcase) saveNewAvatar(avatar io.Reader) (string, error) {
	avatarImg, err := imaging.Decode(avatar)
	if err != nil {
		return "", servErrors.NewError(servErrors.DECODE_IMG, err.Error())
	}

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

	err = u.StaticManager.SafeAvatar(avatarImg, avatarName)
	if err != nil {
		return "", servErrors.NewError(servErrors.CANT_SAVE_AVATAR, err.Error())
	}

	return avatarName, nil
}

func newAvatarName() string {
	fname, _ := uuid.NewUUID()
	return fname.String()
}
