package orderHandler

import (
	"fmt"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/delivery/http/httpErrDescr"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/delivery/http/middleware"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/order"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	Usecase order.Usecase
}

func NewOrderHandler(usecase order.Usecase) *OrderHandler {
	return &OrderHandler{
		Usecase: usecase,
	}
}

func (h OrderHandler) Order(ctx echo.Context) error {
	user := middleware.GetUserFromCtx(ctx)
	if user == nil {
		return echo.NewHTTPError(http.StatusUnauthorized, httpErrDescr.AUTH_REQUIRED)
	}
	logger := middleware.GetLoggerFromCtx(ctx)
	requestId := middleware.GetRequestIdFromCtx(ctx)

	var orderReq models.OrderReq
	if err := ctx.Bind(&orderReq); err != nil {
		fmt.Println(err)
		fmt.Println(ctx.Request().Body)
		return echo.NewHTTPError(http.StatusBadRequest, httpErrDescr.BAD_REQUEST_BODY)
	}
	if _, err := govalidator.ValidateStruct(orderReq); err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, httpErrDescr.INVALID_DATA)
	}

	newOrderId, err := h.Usecase.Order(&models.OrderUcaseInput{UserId: user.Id, Address: orderReq.Address, Cart: orderReq.Cart, Comment: orderReq.Comment})
	if err != nil {
		cause := servErrors.ErrorAs(err)
		if cause == nil {
			logger.Error(requestId, err.Error())
			return echo.NewHTTPError(http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
		}
		switch cause.Code {
		case servErrors.DB_INSERT:
			return echo.NewHTTPError(http.StatusConflict, httpErrDescr.CREATING_ORDER)
		default:
			logger.Error(requestId, err.Error())
			return echo.NewHTTPError(http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
		}
	}
	if newOrderId == nil || newOrderId.OrderId == 0 {
		logger.Error(requestId, "from user-usecase-register returned userData==nil and err==nil, unknown error")
		return echo.NewHTTPError(http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
	}

	return ctx.JSON(http.StatusOK, models.OrderRepoResp{OrderId: newOrderId.OrderId})
}
