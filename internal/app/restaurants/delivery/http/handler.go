package restaurantsHandler

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/delivery/http/httpErrDescr"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/delivery/http/middleware"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/restaurants"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type RestaurantsHandler struct {
	Usecase     restaurants.Usecase
}

func NewRestaurantsHandler(usecase restaurants.Usecase) *RestaurantsHandler {
	return &RestaurantsHandler{
		Usecase:     usecase,
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

	restaurantsD := &models.RestaurantsResponse{}

	for _, rest := range restaurantsDataDelivery.Restaurants {
		item := &models.RestaurantJson{
			Id: rest.Id,
			Name: rest.Name,
			City: rest.City,
			Address: rest.Address,
			Image_path: "http://tavide.xyz:8080/static/" + rest.Image_path,
			Slug: rest.Slug,
			Min_price: rest.Min_price,
			Avg_price: rest.Avg_price,
			Rating: rest.Rating,
		}
		restaurantsD.Restaurants = append(restaurantsD.Restaurants, *item)
	}

	return ctx.JSON(http.StatusOK, models.RestaurantsResponse{Restaurants: restaurantsD.Restaurants})
}