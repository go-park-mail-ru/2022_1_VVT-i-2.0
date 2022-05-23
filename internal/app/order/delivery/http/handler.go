package orderHandler

import (
	"encoding/json"
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
	Ucase         order.Ucase
	StaticManager staticManager.FileManager
}

func NewOrderHandler(ucase order.Ucase, staticManager staticManager.FileManager) *OrderHandler {
	return &OrderHandler{
		Ucase:         ucase,
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
		return httpErrDescr.NewHTTPError(ctx, http.StatusBadRequest, httpErrDescr.BAD_REQUEST_BODY)
	}
	if _, err := govalidator.ValidateStruct(orderReq); err != nil {
		return httpErrDescr.NewHTTPError(ctx, http.StatusBadRequest, httpErrDescr.INVALID_DATA)
	}

	newOrderId, err := h.Ucase.CreateOrder(&models.OrderUcaseReq{
		UserId:    user.Id,
		Address:   orderReq.Address,
		Cart:      orderReq.Cart,
		Promocode: orderReq.Promocode,
		Comment:   orderReq.Comment})
	if err != nil {
		cause := servErrors.ErrorAs(err)
		if cause == nil {
			logger.Error(requestId, err.Error())
			return httpErrDescr.NewHTTPError(ctx, http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
		}
		switch cause.Code {
		case servErrors.DB_INSERT:
			return httpErrDescr.NewHTTPError(ctx, http.StatusConflict, httpErrDescr.CREATING_ORDER)
		case servErrors.NO_SUCH_ADDRESS:
			return httpErrDescr.NewHTTPError(ctx, http.StatusConflict, httpErrDescr.NO_SUCH_ADDRESS)
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

	orders, err := h.Ucase.GetUserOrders(&models.GetUserOrdersUcaseReq{UserId: int64(user.Id)})
	if err != nil {
		cause := servErrors.ErrorAs(err)
		if cause != nil && cause.Code == servErrors.NO_SUCH_ENTITY_IN_DB {
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

	orderStatuses, err := h.Ucase.GetUserOrderStatuses(&models.GetUserOrderStatusesUcaseReq{UserId: int64(user.Id)})
	if err != nil {
		cause := servErrors.ErrorAs(err)
		if cause != nil && cause.Code == servErrors.NO_SUCH_ENTITY_IN_DB {
			return httpErrDescr.NewHTTPError(ctx, http.StatusForbidden, httpErrDescr.NO_SUCH_USER)
		}
		logger.Error(requestId, err.Error())
		return httpErrDescr.NewHTTPError(ctx, http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
	}

	resp := models.GetUserOrderStatusesResp{OrderStatuses: make([]models.OrderStatus, len(orderStatuses.OrderStatuses))}
	for i, order := range orderStatuses.OrderStatuses {
		resp.OrderStatuses[i] = models.OrderStatus(order)
	}
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

	orderUcaseData, err := h.Ucase.GetUserOrder(&models.GetUserOrderUcaseReq{UserId: int64(user.Id), OrderId: int64(orderId)})
	if err != nil {
		cause := servErrors.ErrorAs(err)
		if cause != nil && cause.Code == servErrors.THIS_ORDER_DOESNOT_BELONG_USER {
			return httpErrDescr.NewHTTPError(ctx, http.StatusForbidden, httpErrDescr.THIS_ORDER_DOESNOT_BELONG_USER)
		}

		if cause != nil && cause.Code == servErrors.NO_SUCH_ENTITY_IN_DB {
			return httpErrDescr.NewHTTPError(ctx, http.StatusForbidden, httpErrDescr.NO_SUCH_USER)
		}
		logger.Error(requestId, err.Error())
		return httpErrDescr.NewHTTPError(ctx, http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
	}
	resp := models.GetUserOrderResp{
		OrderId:        orderUcaseData.OrderId,
		Address:        orderUcaseData.Address,
		Date:           orderUcaseData.Date,
		RestaurantName: orderUcaseData.RestaurantName,
		RestaurantSlug: orderUcaseData.RestaurantSlug,
		Status:         orderUcaseData.Status,
		TotalPrice:     orderUcaseData.TotalPrice,
		Discount:       orderUcaseData.Discount,
		Cart:           make([]models.OrderPositionResp, len(orderUcaseData.Cart))}
	for i, order := range orderUcaseData.Cart {
		order.ImagePath = h.StaticManager.GetDishesUrl(order.ImagePath)
		resp.Cart[i] = models.OrderPositionResp(order)
	}
	result, _ := json.Marshal(resp)
	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}
