package restaurantsHandler

import (
	"net/http"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/delivery/http/httpErrDescr"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/delivery/http/middleware"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	promocode "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/promocode"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/staticManager"
	"github.com/labstack/echo/v4"
)

type PromocodesHandler struct {
	Ucase         promocode.Ucase
	StaticManager staticManager.FileManager
}

func NewPromocodesHandler(ucase promocode.Ucase, staticManager staticManager.FileManager) *PromocodesHandler {
	return &PromocodesHandler{
		Ucase:         ucase,
		StaticManager: staticManager,
	}
}

func (h *PromocodesHandler) GetAllPromocodes(ctx echo.Context) error {
	logger := middleware.GetLoggerFromCtx(ctx)
	requestId := middleware.GetRequestIdFromCtx(ctx)

	promosUcaseResp, err := h.Ucase.GetAllPromocodes()

	if err != nil {
		logger.Error(requestId, err.Error())
		return httpErrDescr.NewHTTPError(ctx, http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
	}

	promosResp := &models.GetPromocodesResp{Promos: make([]models.PromocodeResp, len(promosUcaseResp.Promos))}
	for i, promo := range promosUcaseResp.Promos {
		promosResp.Promos[i] = models.PromocodeResp(promo)
		promosResp.Promos[i].Image = h.StaticManager.GetRestaurantUrl(promo.Image)
	}

	return ctx.JSON(http.StatusOK, promosResp)
}
