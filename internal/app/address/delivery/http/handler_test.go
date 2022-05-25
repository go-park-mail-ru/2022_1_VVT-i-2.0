package suggestHandler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/address/mock"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/delivery/http/middleware"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/logger"
	mockLogger "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/logger/mock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestSuggsHandler_Suggest(t *testing.T) {
	mockUCase := new(mock.AddrUcase)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/suggest?q=Москва, Тверская улица, 2", strings.NewReader(``))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := NewSuggsHandler(mockUCase)

	jsonResp := `{"suggests":[{"address":"Москва, Тверская улица, 12","end":true}]}
`
	err = handler.Suggest(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, jsonResp, rec.Body.String())
}

func TestSuggsHandler_SuggestInvalidQuery(t *testing.T) {
	mockUCase := new(mock.AddrUcase)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, `/suggest?q=Навалидные даннные!!!`, strings.NewReader(``))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := NewSuggsHandler(mockUCase)

	err = handler.Suggest(c)
	assert.Error(t, err)
	assert.Equal(t, http.StatusBadRequest, c.Response().Status)
}

func TestSuggsHandler_SuggestUcaseErr(t *testing.T) {
	mockUCase := new(mock.AddrUcaseErr)
	mockLogger := new(mockLogger.Logger)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/suggest?q=Москва, Тверская улица, 2", strings.NewReader(``))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set(middleware.LoggerCtxKey, &logger.ServLogger{Logger: mockLogger})
	handler := NewSuggsHandler(mockUCase)

	err = handler.Suggest(c)
	assert.Error(t, err)
	assert.Equal(t, http.StatusInternalServerError, c.Response().Status)
}
