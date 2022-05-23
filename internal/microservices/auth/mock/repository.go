package mock

// import (
// 	"database/sql"

// 	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/auth/models"
// 	"github.com/stretchr/testify/mock"
// )

// type AuthRepo struct {
// 	mock.Mock
// }

// func (u AuthRepo) AddUser(storage *models.AddUserRepoReq) (*models.UserDataRepo, error) {
// 	return &models.UserDataRepo{
// 		Id:     1,
// 		Name:   "Name",
// 		Phone:  "79999999999",
// 		Email:  "email@mail.com",
// 		Avatar: sql.NullString{String: "avatar.png"}}, nil
// }

// func (u AuthRepo) GetUserByPhone(phone models.UserByPhoneRepoReq) (*models.UserDataRepo, error) {
// 	return &models.UserDataRepo{
// 		Id:     1,
// 		Name:   "Name",
// 		Phone:  "79999999999",
// 		Email:  "email@mail.com",
// 		Avatar: sql.NullString{String: "avatar.png"}}, nil
// }

// func (u AuthRepo) HasUserByPhone(phone models.UserByPhoneRepoReq) (models.HasSuchUserRepoResp, error) {
// 	return models.HasSuchUserRepoResp{IsRegistered: true}, nil
// }
