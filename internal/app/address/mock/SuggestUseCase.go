package mock

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/stretchr/testify/mock"
)

type AddrUcase struct {
	mock.Mock
}

func (s *AddrUcase) Suggest(req *models.SuggestUcaseReq) (*models.SuggestsUcaseResp, error) {
	return &models.SuggestsUcaseResp{Suggests: []models.OneSuggestUcaseResp{{Address: "Москва, Тверская улица, 12", Full: true}}}, nil
}

type AddrUcaseErr struct {
	mock.Mock
}

func (s *AddrUcaseErr) Suggest(req *models.SuggestUcaseReq) (*models.SuggestsUcaseResp, error) {
	return &models.SuggestsUcaseResp{}, servErrors.NewError(servErrors.DB_ERROR, "")
}
