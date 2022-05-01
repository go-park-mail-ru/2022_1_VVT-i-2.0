package restaurantsHandler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	conf "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/config"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/domain/interfaces"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/staticManager/localStaticManager"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRestaurantsHandler_GetAllRestaurants(t *testing.T) {
	config := conf.NewConfig()
	err := conf.ReadConfigFile("../config/serv.toml", config)

	staticManager := localStaticManager.NewLocalFileManager(config.ServConfig.StaticUrl, config.ServConfig.StaticPath)

	mockUCase := new(interfaces.CommentsUsecase)

	e := echo.New()

	req, err := http.NewRequest(echo.GET, "/restaurants", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := RestaurantsHandler{
		Usecase:       mockUCase,
		StaticManager: staticManager,
	}

	json := `[{"id":1,"restName":"Name","imgPath":"restaurants/imgPath","slug":"slug","price":1,"rating":1,"timeToDeliver":"1-1"}]`

	err = handler.GetAllRestaurants(c)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, json, rec.Body.String())
}
