package ucase

import (
	"testing"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/address/mock"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/stretchr/testify/assert"
)

func TestSuggestUcase_SuggestCity(t *testing.T) {
	mockDishesRepo := new(mock.AddrRepo)
	ucase := NewAddrUcase(mockDishesRepo)

	suggs, err := ucase.Suggest(&models.SuggestUcaseReq{Address: "Мо"})
	assert.NoError(t, err)

	assert.Equal(t, cityResp, suggs.Suggests)
}

func TestSuggestUcase_SuggestCityUser(t *testing.T) {
	mockDishesRepo := new(mock.AddrRepo)
	ucase := NewAddrUcase(mockDishesRepo)

	suggs, err := ucase.Suggest(&models.SuggestUcaseReq{Address: "Мо", UserId: 1})
	assert.NoError(t, err)

	assert.Equal(t, []models.OneSuggestUcaseResp{{Address: "Москва, Петровка,38", Full: true}, {Address: "Москва, ", Full: false}}, suggs.Suggests)
}

func TestSuggestUcase_SuggestStreet(t *testing.T) {
	mockAddrRepo := new(mock.AddrRepo)
	ucase := NewAddrUcase(mockAddrRepo)

	suggs, err := ucase.Suggest(&models.SuggestUcaseReq{Address: "Москвa, "})
	assert.NoError(t, err)

	expect := models.SuggestsUcaseResp{Suggests: []models.OneSuggestUcaseResp{{Address: "Москва, Измайловский проспект, ", Full: false}, {Address: "Москва, Измайловская улица, ", Full: false}}}
	assert.Equal(t, suggs.Suggests, expect.Suggests)
}

func TestSuggestUcase_SuggestStreetUser(t *testing.T) {
	mockAddrRepo := new(mock.AddrRepo)
	ucase := NewAddrUcase(mockAddrRepo)

	suggs, err := ucase.Suggest(&models.SuggestUcaseReq{Address: "Москвa, ", UserId: 1})
	assert.NoError(t, err)

	expect := models.SuggestsUcaseResp{Suggests: []models.OneSuggestUcaseResp{{Address: "Москва, Петровка,38", Full: true}, {Address: "Москва, Измайловский проспект, ", Full: false}, {Address: "Москва, Измайловская улица, ", Full: false}}}
	assert.Equal(t, suggs.Suggests, expect.Suggests)
}

func TestSuggestUcase_SuggestHouse(t *testing.T) {
	mockDishesRepo := new(mock.AddrRepo)
	ucase := NewAddrUcase(mockDishesRepo)

	suggs, err := ucase.Suggest(&models.SuggestUcaseReq{Address: "Москва, Измайловский проспект,"})
	assert.NoError(t, err)

	expect := models.SuggestsUcaseResp{Suggests: []models.OneSuggestUcaseResp{{Address: "Москва, Измайловский проспект, 1", Full: true}}}

	assert.Equal(t, expect.Suggests, suggs.Suggests)
}

func TestSuggestUcase_SuggestHouseUser(t *testing.T) {
	mockDishesRepo := new(mock.AddrRepo)
	ucase := NewAddrUcase(mockDishesRepo)

	suggs, err := ucase.Suggest(&models.SuggestUcaseReq{Address: "Москва, Измайловский проспект,", UserId: 1})
	assert.NoError(t, err)

	expect := models.SuggestsUcaseResp{Suggests: []models.OneSuggestUcaseResp{{Address: "Москва, Петровка,38", Full: true}, {Address: "Москва, Измайловский проспект, 1", Full: true}}}

	assert.Equal(t, expect.Suggests, suggs.Suggests)
}

func TestSuggestUcase_SuggestWrongHouse(t *testing.T) {
	mockDishesRepo := new(mock.AddrRepoGetHouseErr)
	ucase := NewAddrUcase(mockDishesRepo)

	suggs, err := ucase.Suggest(&models.SuggestUcaseReq{Address: "Москва, Измайловский проспект, 3"})
	assert.NoError(t, err)

	expect := models.SuggestsUcaseResp{Suggests: []models.OneSuggestUcaseResp{{Address: "Москва, Измайловский проспект, 1", Full: true}}}

	assert.Equal(t, expect.Suggests, suggs.Suggests)
}

func TestSuggestUcase_SuggestStreetErr(t *testing.T) {
	mockAddrRepo := new(mock.AddrRepoErr)
	ucase := NewAddrUcase(mockAddrRepo)

	suggs, err := ucase.Suggest(&models.SuggestUcaseReq{Address: "Москвa, "})
	assert.NoError(t, err)

	expect := models.SuggestsUcaseResp{Suggests: []models.OneSuggestUcaseResp{{Address: "Москва, ", Full: false}}}
	assert.Equal(t, expect.Suggests, suggs.Suggests)
}
