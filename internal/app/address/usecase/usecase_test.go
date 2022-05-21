package ucase

// import (
// 	"testing"

// 	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/address/mock"
// 	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
// 	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
// 	"github.com/stretchr/testify/assert"
// )

// func TestSuggestUcase_SuggestCity(t *testing.T) {
// 	mockDishesRepo := new(mock.AddrRepo)
// 	ucase := NewAddrUcase(mockDishesRepo)

// 	suggs, err := ucase.Suggest(&models.SuggestUcaseReq{Address: "Мо"})
// 	assert.NoError(t, err)

// 	expect := models.SuggestResp{Suggests: defaultRes}

// 	assert.False(t, suggs.AddressFull)
// 	assert.Equal(t, expect.Suggests, suggs.Suggests)

// 	suggs, err = ucase.Suggest(nil)
// 	assert.NoError(t, err)

// 	assert.False(t, suggs.AddressFull)
// 	assert.Equal(t, expect.Suggests, suggs.Suggests)
// }

// func TestSuggestUcase_SuggestStreet(t *testing.T) {
// 	mockAddrRepo := new(mock.AddrRepo)
// 	ucase := NewAddrUcase(mockAddrRepo)

// 	suggs, err := ucase.Suggest(&models.SuggestUcaseReq{Address: "Москвa, "})
// 	assert.NoError(t, err)

// 	expect := models.SuggestResp{Suggests: []string{"Москва, Измайловская улица, ", "Москва, Измайловский проспект, "}}

// 	assert.False(t, suggs.AddressFull)
// 	assert.Equal(t, suggs.Suggests, expect.Suggests)
// }

// func TestSuggestUcase_SuggestHouse(t *testing.T) {
// 	mockDishesRepo := new(mock.AddrRepo)
// 	ucase := NewAddrUcase(mockDishesRepo)

// 	suggs, err := ucase.Suggest(&models.SuggestUcaseReq{Address: "Москва, Измайловский проспект,"})
// 	assert.NoError(t, err)

// 	expect := models.SuggestResp{Suggests: []string{"Москва, Измайловский проспект, 1"}}

// 	assert.True(t, suggs.AddressFull)
// 	assert.Equal(t, expect.Suggests, suggs.Suggests)
// }

// func TestSuggestUcase_SuggestWrongHouse(t *testing.T) {
// 	mockDishesRepo := new(mock.AddrRepoGetHouseErr)
// 	ucase := NewAddrUcase(mockDishesRepo)

// 	suggs, err := ucase.Suggest(&models.SuggestUcaseReq{Address: "Москва, Измайловский проспект, 3"})
// 	assert.NoError(t, err)

// 	expect := models.SuggestResp{Suggests: []string{"Москва, Измайловский проспект, 1"}}

// 	assert.True(t, suggs.AddressFull)
// 	assert.Equal(t, expect.Suggests, suggs.Suggests)
// }

// func TestSuggestUcase_SuggestStreetErr(t *testing.T) {
// 	mockAddrRepo := new(mock.AddrRepoErr)
// 	ucase := NewAddrUcase(mockAddrRepo)

// 	suggs, err := ucase.Suggest(&models.SuggestUcaseReq{Address: "Москвa, "})
// 	assert.NoError(t, err)

// 	assert.Equal(t, defaultRes, suggs.Suggests)
// 	assert.False(t, suggs.AddressFull)
// }

// func TestSuggestUcase_SuggestHouseErr(t *testing.T) {
// 	mockDishesRepo := new(mock.AddrRepoHouseErr)
// 	ucase := NewAddrUcase(mockDishesRepo)

// 	suggs, err := ucase.Suggest(&models.SuggestUcaseReq{Address: "Москва, Измайловский проспект,"})
// 	assert.Error(t, err)
// 	cause := servErrors.ErrorAs(err)
// 	assert.Equal(t, *cause, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, ""))
// 	assert.Nil(t, suggs)
// }
