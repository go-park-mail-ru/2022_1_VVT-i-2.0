package restaurantsHandler

import (
	"encoding/json"
	"math"
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/delivery/http/httpErrDescr"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/delivery/http/middleware"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/restaurants"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/staticManager"
	"github.com/labstack/echo/v4"
)

type RestaurantsHandler struct {
	Usecase       restaurants.Usecase
	StaticManager staticManager.FileManager
}

func NewRestaurantsHandler(usecase restaurants.Usecase, staticManager staticManager.FileManager) *RestaurantsHandler {
	return &RestaurantsHandler{
		Usecase:       usecase,
		StaticManager: staticManager,
	}
}

// GetAllRestaurants Restaurants godoc
// @Summary      List restaurants
// @Description  Get restaurants
// @Tags         Restaurants
// @Accept       json
// @Produce      json
// @Success      200  {object} []models.RestaurantJsonForKirill
// @Router       /restaurants [get]
func (h RestaurantsHandler) GetAllRestaurants(ctx echo.Context) error {
	logger := middleware.GetLoggerFromCtx(ctx)
	requestId := middleware.GetRequestIdFromCtx(ctx)

	restaurantsDataDelivery, err := h.Usecase.GetAllRestaurants()

	if err != nil {
		cause := servErrors.ErrorAs(err)
		if cause != nil && cause.Code == servErrors.NO_SUCH_ENTITY_IN_DB {
			return echo.NewHTTPError(http.StatusForbidden, httpErrDescr.NO_SUCH_RESTAURANTS)
		}
		logger.Error(requestId, err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
	}

	if restaurantsDataDelivery == nil {
		logger.Error(requestId, "from restaurants-handler-getall returned restaurantsDataDelivery==nil and err==nil, unknown error")
		return echo.NewHTTPError(http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
	}

	restaurantsD := &models.AllRestaurantsResp{}
	for _, rest := range restaurantsDataDelivery.Restaurants {
		rating := 0.0
		if rest.ReviewCount != 0 {
			rating = math.Round(float64(rest.AggRating)*10/float64(rest.ReviewCount)) / 10
		}
		item := models.RestaurantResp{
			Id:             rest.Id,
			Name:           rest.Name,
			ImagePath:      h.StaticManager.GetRestaurantUrl(rest.ImagePath),
			Slug:           rest.Slug,
			MinPrice:       rest.MinPrice,
			Rating:         rating,
			TimeToDelivery: strconv.Itoa(rest.DownMinutsToDelivery) + "-" + strconv.Itoa(rest.UpMinutsToDelivery),
		}
		restaurantsD.Restaurants = append(restaurantsD.Restaurants, item)
	}

	result, _ := json.Marshal(restaurantsD.Restaurants)
	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

func (h RestaurantsHandler) GetRestaurantsByCategory(ctx echo.Context, category string) error {
	logger := middleware.GetLoggerFromCtx(ctx)
	requestId := middleware.GetRequestIdFromCtx(ctx)

	restaurantsDataDelivery, err := h.Usecase.GetRestaurantsByCategory(models.GetRestaurantByCategoryUcaseReq{Name: category})

	if err != nil {
		cause := servErrors.ErrorAs(err)
		if cause != nil && cause.Code == servErrors.NO_SUCH_ENTITY_IN_DB {
			return echo.NewHTTPError(http.StatusForbidden, httpErrDescr.NO_SUCH_RESTAURANTS)
		}
		logger.Error(requestId, err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
	}

	if restaurantsDataDelivery == nil {
		logger.Error(requestId, "from restaurants-handler-getall returned restaurantsDataDelivery==nil and err==nil, unknown error")
		return echo.NewHTTPError(http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
	}

	restaurantsD := &models.AllRestaurantsResp{}
	for _, rest := range restaurantsDataDelivery.Restaurants {
		rating := 0.0
		if rest.ReviewCount != 0 {
			rating = math.Round(float64(rest.AggRating)*10/float64(rest.ReviewCount)) / 10
		}
		item := models.RestaurantResp{
			Id:             rest.Id,
			Name:           rest.Name,
			ImagePath:      h.StaticManager.GetRestaurantUrl(rest.ImagePath),
			Slug:           rest.Slug,
			MinPrice:       rest.MinPrice,
			Rating:         rating,
			TimeToDelivery: strconv.Itoa(rest.DownMinutsToDelivery) + "-" + strconv.Itoa(rest.UpMinutsToDelivery),
		}
		restaurantsD.Restaurants = append(restaurantsD.Restaurants, item)
	}

	result, _ := json.Marshal(restaurantsD.Restaurants)
	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

// если есть параметр категории -- вываливать категорию
// если есть параметр поска делать поиск по категории или по ресторанам
// иначе вываливать все

func (h RestaurantsHandler) GetAllRestaurantsMain(ctx echo.Context) error {
	// searhQuery := ctx.QueryParam("q")
	category := ctx.QueryParam("category")

	if category != "" {
		return h.GetRestaurantsByCategory(ctx, category)
	}

	// if searhQuery != "" {
	// 	return h.GetRestaurantsBySeachQuery(ctx, searhQuery)
	// }

	return h.GetAllRestaurants(ctx)
}
