package userHandler

import (
	"fmt"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/authManager"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/delivery/http/httpErrDescr"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/delivery/http/middleware"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/user"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	Usecase     user.Usecase
	AuthManager authManager.AuthManager
}

func NewUserHandler(usecase user.Usecase, authManager authManager.AuthManager) *UserHandler {
	return &UserHandler{
		Usecase:     usecase,
		AuthManager: authManager,
	}
}

func (h UserHandler) Login(ctx echo.Context) error {
	if middleware.GetUserFromCtx(ctx) != nil {
		return echo.NewHTTPError(http.StatusConflict, httpErrDescr.ALREADY_AUTHORIZED)
	}
	var loginReq models.LoginRequest
	if err := ctx.Bind(&loginReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, httpErrDescr.BAD_REQUEST_BODY)
	}
	if _, err := govalidator.ValidateStruct(loginReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, httpErrDescr.INVALID_DATA)
	}

	userDataResp, err := h.Usecase.Login(&loginReq)
	if err != nil {
		cause := servErrors.ErrorAs(err)
		if cause != nil && cause.Code == servErrors.CACH_MISS_CODE {
			return echo.NewHTTPError(http.StatusNotFound, httpErrDescr.NO_SUCH_CODE_INFO)
		}
		return echo.NewHTTPError(http.StatusForbidden, httpErrDescr.WRONG_AUTH_CODE)
	}
	return ctx.JSON(http.StatusOK, userDataResp)
}

func (h UserHandler) Register(ctx echo.Context) error {
	if middleware.GetUserFromCtx(ctx) != nil {
		return echo.NewHTTPError(http.StatusConflict, httpErrDescr.ALREADY_AUTHORIZED)
	}
	var registerReq models.RegisterRequest
	if err := ctx.Bind(&registerReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, httpErrDescr.BAD_REQUEST_BODY)
	}
	fmt.Println(registerReq)
	if _, err := govalidator.ValidateStruct(registerReq); err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, httpErrDescr.INVALID_DATA)
	}

	userDataResp, err := h.Usecase.Register(&registerReq)
	if err != nil {
		cause := servErrors.ErrorAs(err)
		if cause != nil && cause.Code == servErrors.CACH_MISS_CODE {
			return echo.NewHTTPError(http.StatusNotFound, httpErrDescr.NO_SUCH_CODE_INFO)
		}
		return echo.NewHTTPError(http.StatusForbidden, httpErrDescr.WRONG_AUTH_CODE)
	}
	if userDataResp == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
	}
	return ctx.JSON(http.StatusOK, userDataResp)
}
