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
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/staticManager"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/validator"
	"github.com/labstack/echo/v4"
)

type RestaurantsHandler struct {
	Ucase         restaurants.Ucase
	StaticManager staticManager.FileManager
}

func NewRestaurantsHandler(ucase restaurants.Ucase, staticManager staticManager.FileManager) *RestaurantsHandler {
	return &RestaurantsHandler{
		Ucase:         ucase,
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

	ucaseRest, err := h.Ucase.GetAllRestaurants()

	if err != nil {
		logger.Error(requestId, err.Error())
		return httpErrDescr.NewHTTPError(ctx, http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
	}

	restaurantsResp := &models.AllRestaurantsResp{Restaurants: make([]models.RestaurantResp, len(ucaseRest.Restaurants))}
	for i, rest := range ucaseRest.Restaurants {
		rating := 0.0
		if rest.ReviewCount != 0 {
			rating = math.Round(float64(rest.AggRating)*10/float64(rest.ReviewCount)) / 10
		}
		restaurantsResp.Restaurants[i] = models.RestaurantResp{
			Id:             rest.Id,
			Name:           rest.Name,
			ImagePath:      h.StaticManager.GetRestaurantUrl(rest.ImagePath),
			Slug:           rest.Slug,
			MinPrice:       rest.MinPrice,
			Rating:         rating,
			TimeToDelivery: strconv.Itoa(rest.DownMinutsToDelivery) + "-" + strconv.Itoa(rest.UpMinutsToDelivery),
		}
	}

	result, _ := json.Marshal(restaurantsResp.Restaurants)
	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

func (h RestaurantsHandler) GetRestaurantsByCategory(ctx echo.Context, category string) error {
	logger := middleware.GetLoggerFromCtx(ctx)
	requestId := middleware.GetRequestIdFromCtx(ctx)

	ucaseResp, err := h.Ucase.GetRestaurantsByCategory(models.GetRestaurantByCategoryUcaseReq{Name: category})

	if err != nil {
		logger.Error(requestId, err.Error())
		return httpErrDescr.NewHTTPError(ctx, http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
	}

	restaurantsResp := &models.AllRestaurantsResp{Restaurants: make([]models.RestaurantResp, len(ucaseResp.Restaurants))}
	for i, rest := range ucaseResp.Restaurants {
		rating := 0.0
		if rest.ReviewCount != 0 {
			rating = math.Round(float64(rest.AggRating)*10/float64(rest.ReviewCount)) / 10
		}
		restaurantsResp.Restaurants[i] = models.RestaurantResp{
			Id:             rest.Id,
			Name:           rest.Name,
			ImagePath:      h.StaticManager.GetRestaurantUrl(rest.ImagePath),
			Slug:           rest.Slug,
			MinPrice:       rest.MinPrice,
			Rating:         rating,
			TimeToDelivery: strconv.Itoa(rest.DownMinutsToDelivery) + "-" + strconv.Itoa(rest.UpMinutsToDelivery),
		}
	}

	result, _ := json.Marshal(restaurantsResp.Restaurants)
	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

func (h RestaurantsHandler) GetRestaurantsBySeachQuery(ctx echo.Context, query string) error {
	logger := middleware.GetLoggerFromCtx(ctx)
	requestId := middleware.GetRequestIdFromCtx(ctx)

	ucaseResp, err := h.Ucase.GetRestaurantBySearchQuery(models.GetRestaurantBySearchQueryUcaseReq{Query: query})

	if err != nil {
		logger.Error(requestId, err.Error())
		return httpErrDescr.NewHTTPError(ctx, http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
	}

	restaurantsResp := &models.AllRestaurantsResp{Restaurants: make([]models.RestaurantResp, len(ucaseResp.Restaurants))}
	for i, rest := range ucaseResp.Restaurants {
		rating := 0.0
		if rest.ReviewCount != 0 {
			rating = math.Round(float64(rest.AggRating)*10/float64(rest.ReviewCount)) / 10
		}
		restaurantsResp.Restaurants[i] = models.RestaurantResp{
			Id:             rest.Id,
			Name:           rest.Name,
			ImagePath:      h.StaticManager.GetRestaurantUrl(rest.ImagePath),
			Slug:           rest.Slug,
			MinPrice:       rest.MinPrice,
			Rating:         rating,
			TimeToDelivery: strconv.Itoa(rest.DownMinutsToDelivery) + "-" + strconv.Itoa(rest.UpMinutsToDelivery),
		}
	}

	result, _ := json.Marshal(restaurantsResp.Restaurants)
	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

func (h RestaurantsHandler) GetAllRestaurantsMain(ctx echo.Context) error {
	searhQuery := ctx.QueryParam("q")
	category := ctx.QueryParam("category")

	if len(category) != 0 && !validator.IsRestaurantCategory(category) {
		return httpErrDescr.NewHTTPError(ctx, http.StatusBadRequest, httpErrDescr.INVALID_CATEGORY)
	}

	if !validator.IsRestaurantQuery(searhQuery) {
		return httpErrDescr.NewHTTPError(ctx, http.StatusBadRequest, httpErrDescr.INVALID_QUERY)
	}

	if category != "" {
		return h.GetRestaurantsByCategory(ctx, category)
	}

	if searhQuery != "" {
		return h.GetRestaurantsBySeachQuery(ctx, searhQuery)
	}

	return h.GetAllRestaurants(ctx)
}
