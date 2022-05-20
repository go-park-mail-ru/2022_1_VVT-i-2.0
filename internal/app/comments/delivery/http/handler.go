package restaurantsHandler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/comments"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/delivery/http/httpErrDescr"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/delivery/http/middleware"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/staticManager"
	"github.com/labstack/echo/v4"
)

type CommentsHandler struct {
	Usecase       comments.Usecase
	StaticManager staticManager.FileManager
}

func NewCommentsHandler(usecase comments.Usecase) *CommentsHandler {
	return &CommentsHandler{
		Usecase: usecase,
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
func (h CommentsHandler) GetRestaurantComments(ctx echo.Context) error {
	logger := middleware.GetLoggerFromCtx(ctx)
	requestId := middleware.GetRequestIdFromCtx(ctx)

	slug := ctx.Param("slug")

	commetsDataDelivery, err := h.Usecase.GetRestaurantComments(models.GetRestaurantCommentsUcaseReq{
		Slug: slug,
	})

	if err != nil {
		cause := servErrors.ErrorAs(err)
		if cause != nil && cause.Code == servErrors.NO_SUCH_ENTITY_IN_DB {
			return httpErrDescr.NewHTTPError(ctx, http.StatusForbidden, httpErrDescr.NO_SUCH_USER)
		}
		logger.Error(requestId, err.Error())
		return httpErrDescr.NewHTTPError(ctx, http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
	}

	commentsD := &models.GetCommentsDataDelivery{Comment: make([]models.GetCommentDataDelivery, len(commetsDataDelivery.Comment))}

	for i, comment := range commetsDataDelivery.Comment {
		commentsD.Comment[i] = models.GetCommentDataDelivery{
			Author: comment.Author,
			Text:   comment.Text,
			Stars:  comment.Stars,
			Date:   comment.Date,
		}
	}

	result, _ := json.Marshal(commentsD.Comment)
	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

// AddRestaurantComment Add comments by restaurant godoc
// @Summary      Add comments by restaurant
// @Description  Add comments by restaurant
// @Tags         Comments
// @Accept       json
// @Produce      json
// @Success      200  {object}   models.CommentDataDelivery
// @Router       /comment [post]
func (h CommentsHandler) AddRestaurantComment(ctx echo.Context) error {
	user := middleware.GetUserFromCtx(ctx)
	if user == nil {
		return ctx.JSON(http.StatusUnauthorized, httpErrDescr.AUTH_REQUIRED)
	}

	logger := middleware.GetLoggerFromCtx(ctx)
	requestId := middleware.GetRequestIdFromCtx(ctx)

	var AddCommentRestaurantUseCaseReq models.AddCommentRestaurantReq
	if err := ctx.Bind(&AddCommentRestaurantUseCaseReq); err != nil {
		return httpErrDescr.NewHTTPError(ctx, http.StatusBadRequest, httpErrDescr.BAD_REQUEST_BODY)
	}

	commetsDataDelivery, err := h.Usecase.AddRestaurantComment(models.AddCommentRestaurantUcaseReq{
		UserId: 		user.Id,
		Slug:         	AddCommentRestaurantUseCaseReq.Slug,
		CommentText:  	AddCommentRestaurantUseCaseReq.CommentText,
		CommentRating:	AddCommentRestaurantUseCaseReq.CommentRating,
	})
	if err != nil {
		cause := servErrors.ErrorAs(err)
		if cause == nil {
			logger.Error(requestId, err.Error())
			return httpErrDescr.NewHTTPError(ctx, http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
		}
		switch cause.Code {
		case servErrors.WRONG_AUTH_CODE:
			return httpErrDescr.NewHTTPError(ctx, http.StatusForbidden, httpErrDescr.WRONG_AUTH_CODE)
		case servErrors.CACH_MISS_CODE, servErrors.NO_SUCH_ENTITY_IN_DB:
			return httpErrDescr.NewHTTPError(ctx, http.StatusNotFound, httpErrDescr.NO_SUCH_CODE_INFO)
		default:
			logger.Error(requestId, err.Error())
			return httpErrDescr.NewHTTPError(ctx, http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
		}
	}

	comment := &models.CommentDataDelivery{
		RestaurantId: commetsDataDelivery.RestaurantId,
		Author:       commetsDataDelivery.Author,
		Text:         commetsDataDelivery.Text,
		Stars:        commetsDataDelivery.Stars,
		Date:         commetsDataDelivery.Date,
	}

	result, _ := json.Marshal(comment)
	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}