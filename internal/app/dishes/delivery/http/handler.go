package restaurantsHandler

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/delivery/http/httpErrDescr"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/delivery/http/middleware"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/dishes"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/staticManager"
	"github.com/labstack/echo/v4"
	"math"
	"net/http"
	"strconv"
)

type DishesHandler struct {
	Ucase         dishes.Ucase
	StaticManager staticManager.FileManager
}

func NewDishesHandler(ucase dishes.Ucase, staticManager staticManager.FileManager) *DishesHandler {
	return &DishesHandler{
		Ucase:         ucase,
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
	if slug == "" {
		return httpErrDescr.NewHTTPError(ctx, http.StatusBadRequest, httpErrDescr.INVALID_DATA)
	}

	restaurantDishes, err := h.Ucase.GetRestaurantDishes(models.GetRestaurantDishesUcaseReq{Slug: slug})

	if err != nil {
		cause := servErrors.ErrorAs(err)
		if cause != nil && cause.Code == servErrors.NO_SUCH_ENTITY_IN_DB {
			return httpErrDescr.NewHTTPError(ctx, http.StatusForbidden, httpErrDescr.NO_SUCH_RESTAURANT)
		}
		logger.Error(requestId, err.Error())
		return httpErrDescr.NewHTTPError(ctx, http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
	}

	if restaurantDishes == nil {
		logger.Error(requestId, "from user-ucase-get-user returned userData==nil and err==nil, unknown error")
		return httpErrDescr.NewHTTPError(ctx, http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
	}

	rating := 0.0
	if restaurantDishes.ReviewCount != 0 {
		rating = math.Round(float64(restaurantDishes.AggRating)*10/float64(restaurantDishes.ReviewCount)) / 10
	}
	resp := &models.GetRestaurantDishesResp{
		Id:             restaurantDishes.Id,
		Name:           restaurantDishes.Name,
		ImagePath:      h.StaticManager.GetRestaurantUrl(restaurantDishes.ImagePath),
		Slug:           restaurantDishes.Slug,
		MinPrice:       restaurantDishes.MinPrice,
		Rating:         rating,
		ReviewCount:    restaurantDishes.ReviewCount,
		TimeToDelivery: strconv.Itoa(restaurantDishes.DownMinutsToDelivery) + "-" + strconv.Itoa(restaurantDishes.UpMinutsToDelivery),
		Dishes:         make([]models.DishResp, len(restaurantDishes.Dishes)),
	}

	for i, dish := range restaurantDishes.Dishes {
		resp.Dishes[i] = models.DishResp{
			Id:           dish.Id,
			RestaurantId: dish.RestaurantId,
			Name:         dish.Name,
			Description:  dish.Description,
			ImagePath:    h.StaticManager.GetDishesUrl(dish.ImagePath),
			Calories:     dish.Calories,
			Price:        dish.Price,
			Weight:       dish.Weight,
		}
	}

	result, _ := json.Marshal(resp)
	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}
