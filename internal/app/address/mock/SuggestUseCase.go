package mock

// import (
// 	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
// 	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
// 	"github.com/stretchr/testify/mock"
// )

// type AddrUcase struct {
// 	mock.Mock
// }

// func (s AddrUcase) Suggest(req *models.SuggestUcaseReq) (*models.SuggestUcaseResp, error) {
// 	return &models.SuggestUcaseResp{AddressFull: false, Suggests: []string{"Москва, Тверская улица, 12"}}, nil
// }

// type AddrUcaseErr struct {
// 	mock.Mock
// }

// func (s AddrUcaseErr) Suggest(req *models.SuggestUcaseReq) (*models.SuggestUcaseResp, error) {
// 	return &models.SuggestUcaseResp{}, servErrors.NewError(servErrors.DB_ERROR, "")
// }
