package orderHandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/delivery/http/httpErrDescr"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/delivery/http/middleware"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/order"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/staticManager"
	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	Usecase       order.Usecase
	StaticManager staticManager.FileManager
}

func NewOrderHandler(usecase order.Usecase, staticManager staticManager.FileManager) *OrderHandler {
	return &OrderHandler{
		Usecase:       usecase,
		StaticManager: staticManager,
	}
}

func (h OrderHandler) CreateOrder(ctx echo.Context) error {
	user := middleware.GetUserFromCtx(ctx)
	if user == nil {
		return httpErrDescr.NewHTTPError(ctx, http.StatusUnauthorized, httpErrDescr.AUTH_REQUIRED)
	}
	logger := middleware.GetLoggerFromCtx(ctx)
	requestId := middleware.GetRequestIdFromCtx(ctx)

	var orderReq models.OrderReq
	if err := ctx.Bind(&orderReq); err != nil {
		fmt.Println(err.Error())
		fmt.Println(ctx.Request().Body)
		return httpErrDescr.NewHTTPError(ctx, http.StatusBadRequest, httpErrDescr.BAD_REQUEST_BODY)
	}
	if _, err := govalidator.ValidateStruct(orderReq); err != nil {
		fmt.Println(err)
		return httpErrDescr.NewHTTPError(ctx, http.StatusBadRequest, httpErrDescr.INVALID_DATA)
	}

	newOrderId, err := h.Usecase.CreateOrder(&models.OrderUcaseReq{UserId: user.Id, Address: orderReq.Address, Cart: orderReq.Cart, Comment: orderReq.Comment})
	if err != nil {
		cause := servErrors.ErrorAs(err)
		if cause == nil {
			logger.Error(requestId, err.Error())
			//return httpErrDescr.NewHTTPError(ctx, http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
			return httpErrDescr.NewHTTPError(ctx, http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
		}
		switch cause.Code {
		case servErrors.DB_INSERT:
			return httpErrDescr.NewHTTPError(ctx, http.StatusConflict, httpErrDescr.CREATING_ORDER)
		default:
			logger.Error(requestId, err.Error())
			return httpErrDescr.NewHTTPError(ctx, http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
		}
	}
	return ctx.JSON(http.StatusOK, models.OrderResp{OrderId: newOrderId.OrderId})
}

func (h OrderHandler) GetUserOrders(ctx echo.Context) error {
	user := middleware.GetUserFromCtx(ctx)
	if user == nil {
		return httpErrDescr.NewHTTPError(ctx, http.StatusUnauthorized, httpErrDescr.AUTH_REQUIRED)
	}
	logger := middleware.GetLoggerFromCtx(ctx)
	requestId := middleware.GetRequestIdFromCtx(ctx)

	orders, err := h.Usecase.GetUserOrders(&models.GetUserOrdersUcaseReq{UserId: int64(user.Id)})
	if err != nil {
		cause := servErrors.ErrorAs(err)
		if cause != nil && cause.Code == servErrors.NO_SUCH_ENTITY_IN_DB {
			//return httpErrDescr.NewHTTPError(ctx, http.StatusForbidden, httpErrDescr.NO_SUCH_USER)
			return httpErrDescr.NewHTTPError(ctx, http.StatusForbidden, httpErrDescr.NO_SUCH_USER)
		}
		logger.Error(requestId, err.Error())
		return httpErrDescr.NewHTTPError(ctx, http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
	}

	resp := models.GetUserOrdersResp{Orders: make([]models.ShortOrder, len(orders.Orders))}
	for i, order := range orders.Orders {
		resp.Orders[i] = models.ShortOrder(order)
	}
	result, _ := json.Marshal(resp)
	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

func (h OrderHandler) GetUserOrderStatuses(ctx echo.Context) error {
	user := middleware.GetUserFromCtx(ctx)
	if user == nil {
		return httpErrDescr.NewHTTPError(ctx, http.StatusUnauthorized, httpErrDescr.AUTH_REQUIRED)
	}
	logger := middleware.GetLoggerFromCtx(ctx)
	requestId := middleware.GetRequestIdFromCtx(ctx)

	orderStatuses, err := h.Usecase.GetUserOrderStatuses(&models.GetUserOrderStatusesUcaseReq{UserId: int64(user.Id)})
	if err != nil {
		cause := servErrors.ErrorAs(err)
		if cause != nil && cause.Code == servErrors.NO_SUCH_ENTITY_IN_DB {
			//return httpErrDescr.NewHTTPError(ctx, http.StatusForbidden, httpErrDescr.NO_SUCH_USER)
			return httpErrDescr.NewHTTPError(ctx, http.StatusForbidden, httpErrDescr.NO_SUCH_USER)
		}
		logger.Error(requestId, err.Error())
		return httpErrDescr.NewHTTPError(ctx, http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
	}

	resp := models.GetUserOrderStatusesResp{OrderStatuses: make([]models.OrderStatus, len(orderStatuses.OrderStatuses))}
	for i, order := range orderStatuses.OrderStatuses {
		resp.OrderStatuses[i] = models.OrderStatus(order)
	}
	fmt.Println(resp)
	result, _ := json.Marshal(resp)
	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}

func (h OrderHandler) GetUserOrder(ctx echo.Context) error {
	user := middleware.GetUserFromCtx(ctx)
	if user == nil {
		return httpErrDescr.NewHTTPError(ctx, http.StatusUnauthorized, httpErrDescr.AUTH_REQUIRED)
	}
	logger := middleware.GetLoggerFromCtx(ctx)
	requestId := middleware.GetRequestIdFromCtx(ctx)

	orderId, err := strconv.Atoi(ctx.Param("orderId"))
	if err != nil {
		return httpErrDescr.NewHTTPError(ctx, http.StatusBadRequest, httpErrDescr.BAD_ORDER_ID)
	}

	orderUcaseData, err := h.Usecase.GetUserOrder(&models.GetUserOrderUcaseReq{UserId: int64(user.Id), OrderId: int64(orderId)})
	if err != nil {
		cause := servErrors.ErrorAs(err)
		if cause != nil && cause.Code == servErrors.THIS_ORDER_DOESNOT_BELONG_USER {
			//return httpErrDescr.NewHTTPError(ctx, http.StatusForbidden, httpErrDescr.THIS_ORDER_DOESNOT_BELONG_USER)
			return httpErrDescr.NewHTTPError(ctx, http.StatusForbidden, httpErrDescr.THIS_ORDER_DOESNOT_BELONG_USER)
		}

		if cause != nil && cause.Code == servErrors.NO_SUCH_ENTITY_IN_DB {
			return httpErrDescr.NewHTTPError(ctx, http.StatusForbidden, httpErrDescr.NO_SUCH_USER)
		}
		logger.Error(requestId, err.Error())
		return httpErrDescr.NewHTTPError(ctx, http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
	}
	fmt.Println(orderUcaseData.Status)
	resp := models.GetUserOrderResp{OrderId: orderUcaseData.OrderId, Address: orderUcaseData.Address, Date: orderUcaseData.Date, RestaurantName: orderUcaseData.RestaurantName, Status: orderUcaseData.Status, TotalPrice: orderUcaseData.TotalPrice, Cart: make([]models.OrderPositionResp, len(orderUcaseData.Cart))}
	for i, order := range orderUcaseData.Cart {
		order.ImagePath = h.StaticManager.GetDishesUrl(order.ImagePath)
		resp.Cart[i] = models.OrderPositionResp(order)
	}
	fmt.Println(resp)
	result, _ := json.Marshal(resp)
	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}
