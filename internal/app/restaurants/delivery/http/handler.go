package restaurantsHandler

import (
	"encoding/json"
	"fmt"

	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/delivery/http/httpErrDescr"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/delivery/http/middleware"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/restaurants"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/labstack/echo/v4"
)

type RestaurantsHandler struct {
	Usecase restaurants.Usecase
}

func NewRestaurantsHandler(usecase restaurants.Usecase) *RestaurantsHandler {
	return &RestaurantsHandler{
		Usecase: usecase,
	}
}

func (h RestaurantsHandler) GetAllRestaurants(ctx echo.Context) error {
	logger := middleware.GetLoggerFromCtx(ctx)
	requestId := middleware.GetRequestIdFromCtx(ctx)

	restaurantsDataDelivery, err := h.Usecase.GetAllRestaurants()

	if err != nil {
		cause := servErrors.ErrorAs(err)
		if cause != nil && cause.Code == servErrors.NO_SUCH_ENTITY_IN_DB {
			return echo.NewHTTPError(http.StatusForbidden, httpErrDescr.NO_SUCH_USER)
		}
		logger.Error(requestId, err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
	}

	if restaurantsDataDelivery == nil {
		logger.Error(requestId, "from user-usecase-get-user returned userData==nil and err==nil, unknown error")
		return echo.NewHTTPError(http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
	}

	restaurantsD := &models.RestaurantsResponseForKirill{}

	for _, rest := range restaurantsDataDelivery.Restaurants {
		item := &models.RestaurantJsonForKirill{
			Id: rest.Id,
			Name: rest.Name,
			City: rest.City,
			Address: rest.Address,
			Image_path: "http://tavide.xyz:8080/static/static/" + rest.Image_path,
			Slug: rest.Slug,
			Min_price: rest.Min_price,
			Avg_price: rest.Avg_price,
			Rating: float64(int(rest.Rating * 10)) / 10,
			TimeToDelivery: "25 - 30",
		}
		restaurantsD.Restaurants = append(restaurantsD.Restaurants, *item)
	}

	result, _ := json.Marshal(restaurantsD.Restaurants)
	fmt.Printf("json string: %s\n", string(result))
	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

func (h RestaurantsHandler) GetDishesByRestaurants(ctx echo.Context) error {
	logger := middleware.GetLoggerFromCtx(ctx)
	requestId := middleware.GetRequestIdFromCtx(ctx)

	slug := ctx.Param("slug")
	restaurantDataDelivery, err := h.Usecase.GetRestaurantBySluf(slug)

	if err != nil {
		cause := servErrors.ErrorAs(err)
		if cause != nil && cause.Code == servErrors.NO_SUCH_ENTITY_IN_DB {
			return echo.NewHTTPError(http.StatusForbidden, httpErrDescr.NO_SUCH_USER)
		}
		logger.Error(requestId, err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
	}

	if restaurantDataDelivery == nil {
		logger.Error(requestId, "from user-usecase-get-user returned userData==nil and err==nil, unknown error")
		return echo.NewHTTPError(http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
	}

	dishesDataDelivery, err := h.Usecase.GetDishByRestaurant(restaurantDataDelivery.Id)

	if err != nil {
		cause := servErrors.ErrorAs(err)
		if cause != nil && cause.Code == servErrors.NO_SUCH_ENTITY_IN_DB {
			return echo.NewHTTPError(http.StatusForbidden, httpErrDescr.NO_SUCH_USER)
		}
		logger.Error(requestId, err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
	}

	if dishesDataDelivery == nil {
		logger.Error(requestId, "from user-usecase-get-user returned userData==nil and err==nil, unknown error")
		return echo.NewHTTPError(http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
	}

	restaurantD := &models.RestaurantsDishJsonForKirill{
		Id: restaurantDataDelivery.Id,
		Name: restaurantDataDelivery.Name,
		City: restaurantDataDelivery.City,
		Address: restaurantDataDelivery.Address,
		Image_path: "http://tavide.xyz:8080/static/static/" + restaurantDataDelivery.Image_path,
		Slug: restaurantDataDelivery.Slug,
		Min_price: restaurantDataDelivery.Min_price,
		Avg_price: restaurantDataDelivery.Avg_price,
		Rating: float64(int(restaurantDataDelivery.Rating * 10)) / 10,
		TimeToDelivery: "25-30",
	}

	for _, dish := range dishesDataDelivery.Dishes {
		item := &models.DishJsonForKirill{
			Id: dish.Id,
			Restaurant: dish.Restaurant,
			Name: dish.Name,
			Description: dish.Description,
			Image_path:  "http://tavide.xyz:8080/static/dish_static/" + dish.Image_path,
			Calories:    dish.Calories,
			Price:       dish.Price,
		}
		restaurantD.Dishes = append(restaurantD.Dishes, *item)
	}

	result, _ := json.Marshal(restaurantD)
	fmt.Printf("json string: %s\n", string(result))
	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}
