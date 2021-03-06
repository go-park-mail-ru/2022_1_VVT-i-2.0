package suggestHandler

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
	Ucase addr.Ucase
}

func NewSuggsHandler(ucase addr.Ucase) *SuggsHandler {
	return &SuggsHandler{
		Ucase: ucase,
	}
}

func (h SuggsHandler) Suggest(ctx echo.Context) error {
	logger := middleware.GetLoggerFromCtx(ctx)
	requestId := middleware.GetRequestIdFromCtx(ctx)
	var suggsReq models.SuggestReq
	suggsReq.Address = ctx.QueryParam("q")
	if _, err := govalidator.ValidateStruct(suggsReq); err != nil {
		return httpErrDescr.NewHTTPError(ctx, http.StatusBadRequest, httpErrDescr.INVALID_DATA)
	}
	user := middleware.GetUserFromCtx(ctx)
	userId := 0
	if user != nil {
		userId = int(user.Id)
	}

	suggsUcaseResp, err := h.Ucase.Suggest(&models.SuggestUcaseReq{Address: suggsReq.Address, UserId: int64(userId)})

	if err != nil {
		logger.Error(requestId, err.Error())
		return httpErrDescr.NewHTTPError(ctx, http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
	}

	suggsResp := models.SuggestsResp{Suggests: make([]models.OneSuggestResp, len(suggsUcaseResp.Suggests))}
	for i, sugg := range suggsUcaseResp.Suggests {
		suggsResp.Suggests[i] = models.OneSuggestResp(sugg)
	}

	return ctx.JSON(http.StatusOK, suggsResp)

}
