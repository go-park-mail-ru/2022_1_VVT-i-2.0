package recommendationsHandler

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/delivery/http/middleware"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/recommendations/mock"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/staticManager/localStaticManager"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRecommendationsHandler_GetRecommendations(t *testing.T) {
	request := models.RecommendationsOrderLists{
		RestId: 1,
		OrderList: []models.RecommendationsOrderPosition{
			{
				Id: 1,
				Count: 1,
			},
		},
	}

	j, _ := json.Marshal(request)

	mockUCase := new(mock.RecommendationsUcase)
	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/recommendation", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set(middleware.UserCtxKey, middleware.UserCtx{Id: 1})

	staticManager := localStaticManager.NewLocalFileManager("http:/static", "/static")

	handler := RecommendationsHandler{
		Ucase: mockUCase,
		StaticManager: staticManager,
	}

	json := `{"dishes":[{"id":2,"category":2,"restaurant":1,"productName":"Name2","description":"Description","imgPath":"http:/staticdishes/DishImagePath","info":200,"price":10,"weight":100},{"id":3,"category":3,"restaurant":1,"productName":"Name3","description":"Description","imgPath":"http:/staticdishes/DishImagePath","info":200,"price":10,"weight":100}]}`

	err = handler.GetRecommendations(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, json, rec.Body.String())

}