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

	//item := ctx.Param("id")
	//id, err := strconv.ParseInt(item, 16, 32)
	slug := ctx.Param("slug")

	commetsDataDelivery, err := h.Usecase.GetRestaurantComments(slug)

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

	var AddCommentRestaurantUseCaseReq models.AddCommentRestaurant
	if err := ctx.Bind(&AddCommentRestaurantUseCaseReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, httpErrDescr.BAD_REQUEST_BODY)
	}

	commetsDataDelivery, err := h.Usecase.AddRestaurantComment(models.UserId(user.Id), &models.AddCommentRestaurantUseCase{
		Restaurant:     AddCommentRestaurantUseCaseReq.Restaurant,
		Comment_text:   AddCommentRestaurantUseCaseReq.Comment_text,
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
		Restaurant_id: commetsDataDelivery.Restaurant_id,
		Author:        commetsDataDelivery.Author,
		Text:          commetsDataDelivery.Text,
		Stars:         commetsDataDelivery.Stars,
		Date:          commetsDataDelivery.Date,
	}

	result, _ := json.Marshal(comment)
	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

//func (h CommentsHandler) AddRestaurantComment(ctx echo.Context) error {
//	if middleware.GetUserFromCtx(ctx) != nil {
//		return echo.NewHTTPError(http.StatusConflict, httpErrDescr.ALREADY_AUTHORIZED)
//	}
//
//	logger := middleware.GetLoggerFromCtx(ctx)
//	requestId := middleware.GetRequestIdFromCtx(ctx)
//
//	var AddCommentRestaurantUseCaseReq models.AddCommentRestaurant
//	if err := ctx.Bind(&AddCommentRestaurantUseCaseReq); err != nil {
//		return echo.NewHTTPError(http.StatusBadRequest, httpErrDescr.BAD_REQUEST_BODY)
//	}
//
//	commetsDataDelivery, err := h.Usecase.AddRestaurantComment(&models.AddCommentRestaurantUseCase{
//		Restaurant:     AddCommentRestaurantUseCaseReq.Restaurant,
//		User_id:        AddCommentRestaurantUseCaseReq.User_id,
//		Comment_text:   AddCommentRestaurantUseCaseReq.Comment_text,
//		Comment_rating: AddCommentRestaurantUseCaseReq.Comment_rating,
//	})
//	if err != nil {
//		cause := servErrors.ErrorAs(err)
//		if cause == nil {
//			logger.Error(requestId, err.Error())
//			return echo.NewHTTPError(http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
//		}
//		switch cause.Code {
//		case servErrors.WRONG_AUTH_CODE:
//			return echo.NewHTTPError(http.StatusForbidden, httpErrDescr.WRONG_AUTH_CODE)
//		case servErrors.CACH_MISS_CODE, servErrors.NO_SUCH_ENTITY_IN_DB:
//			return echo.NewHTTPError(http.StatusNotFound, httpErrDescr.NO_SUCH_CODE_INFO)
//		default:
//			logger.Error(requestId, err.Error())
//			return echo.NewHTTPError(http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
//		}
//	}
//
//	if commetsDataDelivery == nil {
//		logger.Error(requestId, "from user-usecase-register returned userData==nil and err==nil, unknown error")
//		return echo.NewHTTPError(http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
//	}
//
//	comment := &models.CommentDataDelivery{
//		Id:             commetsDataDelivery.Id,
//		Restaurant:     commetsDataDelivery.Restaurant,
//		User_id:        commetsDataDelivery.User_id,
//		Comment_text:   commetsDataDelivery.Comment_text,
//		Comment_rating: commetsDataDelivery.Comment_rating,
//	}
//
//	result, _ := json.Marshal(comment)
//	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
//	return ctx.JSONBlob(http.StatusOK, result)
//}
