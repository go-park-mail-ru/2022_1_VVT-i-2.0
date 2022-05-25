package restaurantsHandler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/promocode/mock"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/staticManager/localStaticManager"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestRestaurantsHandler_GetAllRestaurants(t *testing.T) {
	staticManager := localStaticManager.NewLocalFileManager("http://localhost/static/", "static")

	mockUCase := new(mock.PromoUcase)

	e := echo.New()

	req, err := http.NewRequest(echo.GET, "/promo", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := NewPromocodesHandler(mockUCase, staticManager)

	json := `{"promos":[{"img":"http://localhost/static/promocodes/main/image1","logo":"http://localhost/static/promocodes/logos/logo1","text":"text1","promocode":"promo1","restName":"name1","restSlug":"slug1","discount":0.01,"priceReduction":2,"minPrice":101},{"img":"http://localhost/static/promocodes/main/image2","logo":"http://localhost/static/promocodes/logos/logo2","text":"text2","promocode":"promo2","restName":"name2","restSlug":"slug2","discount":0.01,"priceReduction":2,"minPrice":102}]}`

	err = handler.GetAllPromocodes(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, json, rec.Body.String())
}
