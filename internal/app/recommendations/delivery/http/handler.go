package recommendationsHandler

import (
	"encoding/json"
	"fmt"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/delivery/http/httpErrDescr"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/recommendations"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/staticManager"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type RecommendationsHandler struct {
	Ucase         recommendations.Ucase
	StaticManager staticManager.FileManager
}

func NewRecommendationsHandler(ucase recommendations.Ucase, staticManager staticManager.FileManager) *RecommendationsHandler {
	return &RecommendationsHandler{
		Ucase:         ucase,
		StaticManager: staticManager,
	}
}

func (h RecommendationsHandler) GetRecommendations(ctx echo.Context) error {
	fmt.Println("начало работы хелндера рекоммендаций")

	//logger := middleware.GetLoggerFromCtx(ctx)
	//requestId := middleware.GetRequestIdFromCtx(ctx)

	var ordersList models.RecommendationsOrderLists
	if err := ctx.Bind(&ordersList); err != nil {
		fmt.Println("не верные входные данные")
		return httpErrDescr.NewHTTPError(ctx, http.StatusBadRequest, httpErrDescr.BAD_REQUEST_BODY)
	}
	fmt.Println(ordersList)

	var OrederListsReq = models.RecommendationsOrderListsUsecaseReq{
		RestId: ordersList.RestId,
		DishesId: make([]int64, len(ordersList.OrderList)),
	}

	for i, item := range ordersList.OrderList {
		OrederListsReq.DishesId[i] = item.Id
	}

	fmt.Println(OrederListsReq)

	recommendations, err := h.Ucase.GetRecommendations(OrederListsReq)

	if err != nil {
		cause := servErrors.ErrorAs(err)
		if cause != nil && cause.Code == servErrors.NO_SUCH_ENTITY_IN_DB {
			return httpErrDescr.NewHTTPError(ctx, http.StatusForbidden, httpErrDescr.NO_SUCH_RESTAURANTS)
		}
		//logger.Error(requestId, err.Error())
		return httpErrDescr.NewHTTPError(ctx, http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
	}

	var rec = &models.DishRecommendationListsDelivery{
		Dishes: make([]models.DishRecommendationDelivery, 3),
	}

	for i := range rec.Dishes {
		var newRec = models.DishRecommendationDelivery {
			Id:				recommendations.Dishes[i].Id,
			Category:		recommendations.Dishes[i].Category,
			RestaurantId:	recommendations.Dishes[i].RestaurantId,
			Name:			recommendations.Dishes[i].Name,
			Description:	recommendations.Dishes[i].Description,
			ImagePath:		recommendations.Dishes[i].ImagePath,
			Calories:		recommendations.Dishes[i].Calories,
			Price:			recommendations.Dishes[i].Price,
			Weight:			recommendations.Dishes[i].Weight,
		}
		rec.Dishes[i] = newRec
	}

	fmt.Println("выход из функции рекоммендаций")
	result, _ := json.Marshal(rec)
	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
	return ctx.JSONBlob(http.StatusOK, result)
}