package restaurantsHandler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/domain/interfaces"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/staticManager/localStaticManager"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDishesHandler_GetDishesByRestaurant(t *testing.T) {
	slug := "my_slug"

	// config := conf.NewConfig()
	// err := conf.ReadConfigFile("../config/serv.toml", config)
	// if err != nil {
	// 	require.NoError(t, err)
	// 	return
	// }

	// staticManager := localStaticManager.NewLocalFileManager(config.ServConfig.StaticUrl, config.ServConfig.StaticPath)
	staticManager := localStaticManager.NewLocalFileManager("", "")

	mockUCase := new(interfaces.DishesUcase)
	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/restaurant/my_slug", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("slug", slug)
	handler := DishesHandler{
		Usecase:       mockUCase,
		StaticManager: staticManager,
	}

	json := `{"id":1,"restName":"Name","imgPath":"restaurants/imgPath","slug":"slug","minPrice":1,"rating":1,"timeToDeliver":"1-1","dishes":[{"id":1,"restaurant":1,"productName":"Name","description":"Description","imgPath":"dishes/imgPath","info":1,"price":1,"weight":1}]}`

	err = handler.GetDishesByRestaurants(c)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, json, rec.Body.String())
}
