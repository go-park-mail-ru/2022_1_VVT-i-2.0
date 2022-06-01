package recommendationsHandler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/delivery/http/httpErrDescr"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/delivery/http/middleware"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/recommendations"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/staticManager"
	"github.com/labstack/echo/v4"
)

const (
	defaultRecommendationCount = 2
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

func (h *RecommendationsHandler) GetRecommendations(ctx echo.Context) error {
	logger := middleware.GetLoggerFromCtx(ctx)
	requestId := middleware.GetRequestIdFromCtx(ctx)

	var req models.RecommendationsReq
	if err := ctx.Bind(&req); err != nil {
		return httpErrDescr.NewHTTPError(ctx, http.StatusBadRequest, httpErrDescr.BAD_REQUEST_BODY)
	}

	var reqUcase = models.RecommendationsUcaseReq{
		RestId:   req.RestId,
		DishesId: make([]int64, len(req.OrderList)),
		Limit:    defaultRecommendationCount,
	}

	for i, item := range req.OrderList {
		reqUcase.DishesId[i] = item.Id
	}

	recommendationsUcaseResp, err := h.Ucase.GetRecommendations(&reqUcase)

	if err != nil {
		cause := servErrors.ErrorAs(err)
		if cause != nil && cause.Code == servErrors.NO_SUCH_ENTITY_IN_DB {
			return httpErrDescr.NewHTTPError(ctx, http.StatusForbidden, httpErrDescr.NO_SUCH_RESTAURANTS)
		}
		logger.Error(requestId, err.Error())
		return httpErrDescr.NewHTTPError(ctx, http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
	}

	if recommendationsUcaseResp == nil {
		result, _ := json.Marshal(models.RecommendationsResp{
			Dishes: make([]models.RecommendationResp, 0),
		})
		ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(result)))
		return ctx.JSONBlob(http.StatusOK, result)
	}

	recommendationResp := &models.RecommendationsResp{Dishes: make([]models.RecommendationResp, 0, defaultRecommendationCount)}

	for i, recommendation := range recommendationsUcaseResp.Dishes {
		recommendationResp.Dishes = append(recommendationResp.Dishes, models.RecommendationResp(recommendation))
		recommendationResp.Dishes[i].ImagePath = h.StaticManager.GetDishesUrl(recommendation.ImagePath)
	}
	respBodyJson, _ := json.Marshal(recommendationResp)
	ctx.Response().Header().Add(echo.HeaderContentLength, strconv.Itoa(len(respBodyJson)))
	return ctx.JSONBlob(http.StatusOK, respBodyJson)
}
