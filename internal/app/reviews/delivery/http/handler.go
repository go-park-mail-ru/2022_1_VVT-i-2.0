package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/delivery/http/httpErrDescr"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/delivery/http/middleware"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	comments "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/reviews"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/staticManager"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/validator"
	"github.com/labstack/echo/v4"
)

type RestaurantReviewsHandler struct {
	Ucase         comments.Usecase
	StaticManager staticManager.FileManager
}

func NewRestaurantReviewsHandler(ucase comments.Usecase) *RestaurantReviewsHandler {
	return &RestaurantReviewsHandler{
		Ucase: ucase,
	}
}

// GetRestaurantComments Get comments by restaurant godoc
// @Summary      List comments by restaurant
// @Description  get comments by restaurant
// @Tags         Comments
// @Accept       json
// @Produce      json
// @Success      200  {object}   models.CommentsDataDelivery
// @Router       /comments/:id [get]
func (h RestaurantReviewsHandler) GetRestaurantComments(ctx echo.Context) error {
	logger := middleware.GetLoggerFromCtx(ctx)
	requestId := middleware.GetRequestIdFromCtx(ctx)

	slug := ctx.Param("slug")
	if !validator.IsSlug(slug) {
		return httpErrDescr.NewHTTPError(ctx, http.StatusBadRequest, httpErrDescr.INVALID_DATA)
	}

	reviewsUcaseResp, err := h.Ucase.GetRestaurantReviews(&models.GetRestaurantReviewsUcaseReq{
		Slug: slug,
	})

	if err != nil {
		logger.Error(requestId, err.Error())
		return httpErrDescr.NewHTTPError(ctx, http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
	}

	reviewsResp := &models.GetRestaurantReviews{Reviews: make([]models.ReviewResp, len(reviewsUcaseResp.Reviews))}

	for i, comment := range reviewsUcaseResp.Reviews {
		reviewsResp.Reviews[i] = models.ReviewResp(comment)
	}

	respBodyJson, _ := json.Marshal(reviewsResp.Reviews)
	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(respBodyJson)))
	return ctx.JSONBlob(http.StatusOK, respBodyJson)
}

// AddRestaurantReview Add comments by restaurant godoc
// @Summary      Add comments by restaurant
// @Description  Add comments by restaurant
// @Tags         Comments
// @Accept       json
// @Produce      json
// @Success      200  {object}   models.CommentDataDelivery
// @Router       /comment [post]
func (h RestaurantReviewsHandler) AddRestaurantReview(ctx echo.Context) error {
	user := middleware.GetUserFromCtx(ctx)
	if user == nil {
		return ctx.JSON(http.StatusUnauthorized, httpErrDescr.AUTH_REQUIRED)
	}

	logger := middleware.GetLoggerFromCtx(ctx)
	requestId := middleware.GetRequestIdFromCtx(ctx)

	var req models.AddRestaurantReviewReq
	if err := ctx.Bind(&req); err != nil {
		return httpErrDescr.NewHTTPError(ctx, http.StatusBadRequest, httpErrDescr.BAD_REQUEST_BODY)
	}

	if !validator.IsOrderId(req.OrderId) {
		return httpErrDescr.NewHTTPError(ctx, http.StatusBadRequest, httpErrDescr.BAD_ORDER_ID)
	}
	if !validator.IsSlug(req.Slug) {
		return httpErrDescr.NewHTTPError(ctx, http.StatusBadRequest, httpErrDescr.INVALID_DATA)
	}
	if !validator.IsReview(req.Text) {
		return httpErrDescr.NewHTTPError(ctx, http.StatusBadRequest, httpErrDescr.INVALID_REVIEW)
	}
	if !validator.IsStars(req.Rating) {
		return httpErrDescr.NewHTTPError(ctx, http.StatusBadRequest, httpErrDescr.INVALID_RATING)
	}

	reviewUcaseResp, err := h.Ucase.AddRestaurantReview(&models.AddRestaurantReviewUcaseReq{
		OrderId: req.OrderId,
		UserId:  user.Id,
		Slug:    req.Slug,
		Text:    req.Text,
		Rating:  req.Rating,
	})
	if err != nil {
		cause := servErrors.ErrorAs(err)
		if cause == nil {
			logger.Error(requestId, err.Error())
			return httpErrDescr.NewHTTPError(ctx, http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
		}
		if cause.Code == servErrors.NO_SUCH_ENTITY_IN_DB {
			return httpErrDescr.NewHTTPError(ctx, http.StatusNotFound, httpErrDescr.NO_SUCH_RESTAURANT)
		}
		if cause.Code == servErrors.ORDER_ALREADY_REVIEWED {
			return httpErrDescr.NewHTTPError(ctx, http.StatusNotFound, httpErrDescr.ORDER_ALREADY_REVIEWED)
		}
		logger.Error(requestId, err.Error())
		return httpErrDescr.NewHTTPError(ctx, http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)

	}

	comment := &models.AddRestaurantReviewResp{
		Author: reviewUcaseResp.Author,
		Text:   reviewUcaseResp.Text,
		Stars:  reviewUcaseResp.Stars,
		Date:   reviewUcaseResp.Date,
	}

	result, _ := json.Marshal(comment)
	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}
