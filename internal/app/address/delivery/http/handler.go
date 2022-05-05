package suggestHandler

// отклонять слишком длиннные запросы

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	addr "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/address"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/delivery/http/httpErrDescr"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/delivery/http/middleware"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	_ "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/validator"
	"github.com/labstack/echo/v4"
)

type SuggsHandler struct {
	Usecase addr.Usecase
}

func NewSuggsHandler(usecase addr.Usecase) *SuggsHandler {
	return &SuggsHandler{
		Usecase: usecase,
	}
}

func (h SuggsHandler) Suggest(ctx echo.Context) error {
	// TODO отключить auth-middleware
	// TODO добавить валидацию адреса
	logger := middleware.GetLoggerFromCtx(ctx)
	requestId := middleware.GetRequestIdFromCtx(ctx)
	var suggsReq models.SuggestReq
	suggsReq.Address = ctx.QueryParam("q")
	// fmt.Println(suggsReq)
	// fmt.Println(ctx.Request().URL.Query())
	if _, err := govalidator.ValidateStruct(suggsReq); err != nil {
		return httpErrDescr.NewHTTPError(ctx, http.StatusBadRequest, httpErrDescr.INVALID_DATA)
	}

	suggsResp, err := h.Usecase.Suggest(&suggsReq)

	if err != nil {
		logger.Error(requestId, err.Error())
	}

	return ctx.JSON(http.StatusOK, suggsResp)

}
