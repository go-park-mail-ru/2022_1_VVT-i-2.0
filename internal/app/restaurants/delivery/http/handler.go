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
		fmt.Println("1")
		cause := servErrors.ErrorAs(err)
		if cause != nil && cause.Code == servErrors.NO_SUCH_ENTITY_IN_DB {
			return echo.NewHTTPError(http.StatusForbidden, httpErrDescr.NO_SUCH_USER)
		}
		logger.Error(requestId, err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
	}

	if restaurantsDataDelivery == nil {
		fmt.Println("2")
		logger.Error(requestId, "from user-usecase-get-user returned userData==nil and err==nil, unknown error")
		return echo.NewHTTPError(http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
	}

	restaurantsD := &models.RestaurantsResponseForKirill{}

	for _, rest := range restaurantsDataDelivery.Restaurants {
		item := &models.RestaurantJsonForKirill{
			Id:         rest.Id,
			Name:       rest.Name,
			City:       rest.City,
			Address:    rest.Address,
			Image_path: h.StaticManager.GetRestaurantUrl(rest.Image_path),
			Slug:           rest.Slug,
			Min_price:      rest.Min_price,
			Avg_price:      rest.Avg_price,
			Rating:         float64(rest.Rating)/float64(rest.Count_rating),
			TimeToDelivery: "25 - 30",
		}
		restaurantsD.Restaurants = append(restaurantsD.Restaurants, *item)
	}

	result, _ := json.Marshal(restaurantsD.Restaurants)
	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

// GetDishesByRestaurants Get dishes by restaurant godoc
// @Summary      List dishes by restaurant
// @Description  get dishes by restaurant
// @Tags         Restaurants
// @Accept       json
// @Produce      json
// @Success      200  {object}   models.RestaurantsDishesJsonForKirill
// @Router       /restaurant/:slug [get]
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

	restaurantD := &models.RestaurantsDishesJsonForKirill{
		Id:      restaurantDataDelivery.Id,
		Name:    restaurantDataDelivery.Name,
		City:    restaurantDataDelivery.City,
		Address: restaurantDataDelivery.Address,
		// Image_path:     "http://127.0.0.1:8080/static/static/" + restaurantDataDelivery.Image_path,
		Image_path:     h.StaticManager.GetRestaurantUrl(restaurantDataDelivery.Image_path),
		Slug:           restaurantDataDelivery.Slug,
		Min_price:      restaurantDataDelivery.Min_price,
		Avg_price:      restaurantDataDelivery.Avg_price,
		Rating:         float64(int(restaurantDataDelivery.Rating*10)) / 10,
		TimeToDelivery: "25-30",
	}

	for _, dish := range dishesDataDelivery.Dishes {
		item := &models.DishJsonForKirill{
			Id:          dish.Id,
			Restaurant:  dish.Restaurant,
			Name:        dish.Name,
			Description: dish.Description,
			Image_path:  h.StaticManager.GetDishesUrl(dish.Image_path),
			// Image_path:  "http://127.0.0.1:8080/static/dish_static/" + dish.Image_path,
			Calories: dish.Calories,
			Price:    dish.Price,
			Weight:   dish.Weight,
		}
		restaurantD.Dishes = append(restaurantD.Dishes, *item)
	}

	result, _ := json.Marshal(restaurantD)
	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

// GetCommentsRestaurantByRestaurants Get comments by restaurant godoc
// @Summary      List comments by restaurant
// @Description  get comments by restaurant
// @Tags         Comments
// @Accept       json
// @Produce      json
// @Success      200  {object}   models.CommentsDataDelivery
// @Router       /comments/:id [get]
func (h RestaurantsHandler) GetCommentsRestaurantByRestaurants(ctx echo.Context) error {
	logger := middleware.GetLoggerFromCtx(ctx)
	requestId := middleware.GetRequestIdFromCtx(ctx)

	//item := ctx.Param("id")
	//id, err := strconv.ParseInt(item, 16, 32)
	id, err := strconv.Atoi(ctx.Param("id"))
	commetsDataDelivery, err := h.Usecase.GetCommentsRestaurantByRestaurants(int(id))

	if err != nil {
		cause := servErrors.ErrorAs(err)
		if cause != nil && cause.Code == servErrors.NO_SUCH_ENTITY_IN_DB {
			return echo.NewHTTPError(http.StatusForbidden, httpErrDescr.NO_SUCH_USER)
		}
		logger.Error(requestId, err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
	}

	if commetsDataDelivery == nil {
		logger.Error(requestId, "from user-usecase-get-user returned userData==nil and err==nil, unknown error")
		return echo.NewHTTPError(http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
	}

	commentsD := &models.CommentsDataDelivery{}

	for _, comment := range commetsDataDelivery.Comment {
		item := &models.CommentDataDelivery{
			Id:         comment.Id,
			Restaurant: comment.Restaurant,
			User_id: comment.User_id,
			Comment_text: comment.Comment_text,
			Comment_rating: comment.Comment_rating,
		}
		commentsD.Comment = append(commentsD.Comment, *item)
	}

	result, _ := json.Marshal(commentsD.Comment)
	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

// AddCommentsRestaurantByRestaurants Add comments by restaurant godoc
// @Summary      Add comments by restaurant
// @Description  Add comments by restaurant
// @Tags         Comments
// @Accept       json
// @Produce      json
// @Success      200  {object}   models.CommentDataDelivery
// @Router       /comment [post]
func (h RestaurantsHandler) AddCommentsRestaurantByRestaurants(ctx echo.Context) error {
	if middleware.GetUserFromCtx(ctx) != nil {
		return echo.NewHTTPError(http.StatusConflict, httpErrDescr.ALREADY_AUTHORIZED)
	}

	logger := middleware.GetLoggerFromCtx(ctx)
	requestId := middleware.GetRequestIdFromCtx(ctx)

	var AddCommentRestaurantUseCaseReq models.AddCommentRestaurant
	if err := ctx.Bind(&AddCommentRestaurantUseCaseReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, httpErrDescr.BAD_REQUEST_BODY)
	}

	commetsDataDelivery, err := h.Usecase.AddCommentsRestaurantByRestaurants(&models.AddCommentRestaurantUseCase{
		Restaurant: AddCommentRestaurantUseCaseReq.Restaurant,
		User_id: AddCommentRestaurantUseCaseReq.User_id,
		Comment_text: AddCommentRestaurantUseCaseReq.Comment_text,
		Comment_rating: AddCommentRestaurantUseCaseReq.Comment_rating,
	})
	if err != nil {
		cause := servErrors.ErrorAs(err)
		if cause == nil {
			logger.Error(requestId, err.Error())
			return echo.NewHTTPError(http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
		}
		switch cause.Code {
		case servErrors.WRONG_AUTH_CODE:
			return echo.NewHTTPError(http.StatusForbidden, httpErrDescr.WRONG_AUTH_CODE)
		case servErrors.CACH_MISS_CODE, servErrors.NO_SUCH_ENTITY_IN_DB:
			return echo.NewHTTPError(http.StatusNotFound, httpErrDescr.NO_SUCH_CODE_INFO)
		default:
			logger.Error(requestId, err.Error())
			return echo.NewHTTPError(http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
		}
	}

	if commetsDataDelivery == nil {
		logger.Error(requestId, "from user-usecase-register returned userData==nil and err==nil, unknown error")
		return echo.NewHTTPError(http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
	}

	comment := &models.CommentDataDelivery{
		Id:         commetsDataDelivery.Id,
		Restaurant: commetsDataDelivery.Restaurant,
		User_id: commetsDataDelivery.User_id,
		Comment_text: commetsDataDelivery.Comment_text,
		Comment_rating: commetsDataDelivery.Comment_rating,
	}

	result, _ := json.Marshal(comment)
	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}