package restaurantsHandler

import (
	"encoding/json"
	"math"
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/delivery/http/httpErrDescr"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/delivery/http/middleware"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/dishes"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/staticManager"
	"github.com/labstack/echo/v4"
)

type DishesHandler struct {
	Usecase       dishes.Usecase
	StaticManager staticManager.FileManager
}

func NewDishesHandler(usecase dishes.Usecase, staticManager staticManager.FileManager) *DishesHandler {
	return &DishesHandler{
		Usecase:       usecase,
		StaticManager: staticManager,
	}
}

// GetDishesByRestaurants Get dishes by restaurant godoc
// @Summary      List dishes by restaurant
// @Description  get dishes by restaurant
// @Tags         Restaurants
// @Accept       json
// @Produce      json
// @Success      200  {object}   models.RestaurantsDishesJsonForKirill
// @Router       /restaurant/:slug [get]
func (h DishesHandler) GetDishesByRestaurants(ctx echo.Context) error {
	logger := middleware.GetLoggerFromCtx(ctx)
	requestId := middleware.GetRequestIdFromCtx(ctx)

	slug := ctx.Param("slug")
	restaurantData, err := h.Usecase.GetRestaurantBySlug(slug)

	if err != nil {
		cause := servErrors.ErrorAs(err)
		if cause != nil && cause.Code == servErrors.NO_SUCH_ENTITY_IN_DB {
			return echo.NewHTTPError(http.StatusForbidden, httpErrDescr.NO_SUCH_USER)
		}
		logger.Error(requestId, err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
	}

	if restaurantData == nil {
		logger.Error(requestId, "from user-usecase-get-user returned userData==nil and err==nil, unknown error")
		return echo.NewHTTPError(http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
	}

	dishesData, err := h.Usecase.GetDishesByRestaurant(restaurantData.Id)

	if err != nil {
		cause := servErrors.ErrorAs(err)
		if cause != nil && cause.Code == servErrors.NO_SUCH_ENTITY_IN_DB {
			return echo.NewHTTPError(http.StatusForbidden, httpErrDescr.NO_SUCH_USER)
		}
		logger.Error(requestId, err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
	}

	if dishesData == nil {
		logger.Error(requestId, "from user-usecase-get-user returned userData==nil and err==nil, unknown error")
		return echo.NewHTTPError(http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
	}

	rating := 0.0
	if restaurantData.ReviewCount != 0 {
		rating = math.Round(float64(restaurantData.AggRating)*10/float64(restaurantData.ReviewCount)) / 10
	}
	restaurantD := &models.RestaurantDishesResp{
		Id:             restaurantData.Id,
		Name:           restaurantData.Name,
		ImagePath:      h.StaticManager.GetRestaurantUrl(restaurantData.ImagePath),
		Slug:           restaurantData.Slug,
		MinPrice:       restaurantData.MinPrice,
		Rating:         rating,
		ReviewCount: 	restaurantData.ReviewCount,
		TimeToDelivery: strconv.Itoa(restaurantData.DownMinutsToDelivery) + "-" + strconv.Itoa(restaurantData.UpMinutsToDelivery),
	}

	for _, dish := range dishesData.Dishes {
		item := &models.DishResp{
			Id:           dish.Id,
			RestaurantId: dish.RestaurantId,
			Name:         dish.Name,
			Description:  dish.Description,
			ImagePath:    h.StaticManager.GetDishesUrl(dish.ImagePath),
			Calories:     dish.Calories,
			Price:        dish.Price,
			Weight:       dish.Weight,
		}
		restaurantD.Dishes = append(restaurantD.Dishes, *item)
	}

	result, _ := json.Marshal(restaurantD)
	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}
