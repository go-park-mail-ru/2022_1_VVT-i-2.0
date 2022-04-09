package userHandler

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/authManager"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/delivery/http/httpErrDescr"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/delivery/http/middleware"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/user"
	"github.com/labstack/echo/v4"
)

const (
	tokenCookieKey = "token"
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

func createTokenCookie(token string, domen string, exp time.Duration) *http.Cookie {
	return &http.Cookie{
		Name:     tokenCookieKey,
		Value:    token,
		HttpOnly: true,
		Expires:  time.Now().Add(exp),
		Domain:   domen,
		Path:     "/",
	}
}

func (h UserHandler) Login(ctx echo.Context) error {
	if middleware.GetUserFromCtx(ctx) != nil {
		return echo.NewHTTPError(http.StatusConflict, httpErrDescr.ALREADY_AUTHORIZED)
	}

	logger := middleware.GetLoggerFromCtx(ctx)
	requestId := middleware.GetRequestIdFromCtx(ctx)

	var loginReq models.LoginRequest
	if err := ctx.Bind(&loginReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, httpErrDescr.BAD_REQUEST_BODY)
	}
	if _, err := govalidator.ValidateStruct(loginReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, httpErrDescr.INVALID_DATA)
	}

	userDataUcase, err := h.Usecase.Login(&loginReq)
	if err != nil {
		cause := servErrors.ErrorAs(err)
		if cause != nil && cause.Code == servErrors.CACH_MISS_CODE {
			return echo.NewHTTPError(http.StatusNotFound, httpErrDescr.NO_SUCH_CODE_INFO)
		}
		return echo.NewHTTPError(http.StatusForbidden, httpErrDescr.WRONG_AUTH_CODE)
	}

	if userDataUcase == nil {
		logger.Error(requestId, "from user-usecase-register returned userData==nil and err==nil, unknown error")
		return echo.NewHTTPError(http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
	}

	token, err := h.AuthManager.CreateToken(*authManager.NewTokenPayload(userDataUcase.Id))
	if err != nil {
		logger.Error(requestId, "error creating token: "+err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
	}

	host, _, _ := net.SplitHostPort(ctx.Request().Host)
	tokenCookie := createTokenCookie(token, host, h.AuthManager.GetEpiryTime())

	ctx.SetCookie(tokenCookie)
	return ctx.JSON(http.StatusOK, models.UserDataResp{Phone: userDataUcase.Phone, Email: userDataUcase.Email, Name: userDataUcase.Name})
}

func (h UserHandler) Register(ctx echo.Context) error {
	if middleware.GetUserFromCtx(ctx) != nil {
		return echo.NewHTTPError(http.StatusConflict, httpErrDescr.ALREADY_AUTHORIZED)
	}

	logger := middleware.GetLoggerFromCtx(ctx)
	requestId := middleware.GetRequestIdFromCtx(ctx)

	var registerReq models.RegisterRequest
	if err := ctx.Bind(&registerReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, httpErrDescr.BAD_REQUEST_BODY)
	}
	fmt.Println(registerReq)
	if _, err := govalidator.ValidateStruct(registerReq); err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, httpErrDescr.INVALID_DATA)
	}

	userDataUcase, err := h.Usecase.Register(&registerReq)
	if err != nil {
		cause := servErrors.ErrorAs(err)
		if cause != nil && cause.Code == servErrors.CACH_MISS_CODE {
			return echo.NewHTTPError(http.StatusNotFound, httpErrDescr.NO_SUCH_CODE_INFO)
		}
		logger.Error(requestId, err.Error())
		return echo.NewHTTPError(http.StatusForbidden, httpErrDescr.WRONG_AUTH_CODE)
	}
	if userDataUcase == nil {
		logger.Error(requestId, "from user-usecase-register returned userData==nil and err==nil, unknown error")
		return echo.NewHTTPError(http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
	}

	token, err := h.AuthManager.CreateToken(*authManager.NewTokenPayload(userDataUcase.Id))
	if err != nil {
		logger.Error(requestId, "error creating token: "+err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
	}

	host, _, _ := net.SplitHostPort(ctx.Request().Host)
	tokenCookie := createTokenCookie(token, host, h.AuthManager.GetEpiryTime())

	ctx.SetCookie(tokenCookie)
	return ctx.JSON(http.StatusOK, models.UserDataResp{Phone: userDataUcase.Phone, Email: userDataUcase.Email, Name: userDataUcase.Name})
}

func (h UserHandler) Logout(ctx echo.Context) error {
	if middleware.GetUserFromCtx(ctx) == nil {
		return echo.NewHTTPError(http.StatusUnauthorized, httpErrDescr.AUTH_REQUIRED)
	}
	host, _, _ := net.SplitHostPort(ctx.Request().Host)
	tokenCookie := createTokenCookie("", host, -time.Hour)

	ctx.SetCookie(tokenCookie)
	return ctx.NoContent(http.StatusOK)
}

func (h UserHandler) SendCode(ctx echo.Context) error {
	if middleware.GetUserFromCtx(ctx) != nil {
		return echo.NewHTTPError(http.StatusConflict, httpErrDescr.ALREADY_AUTHORIZED)
	}

	logger := middleware.GetLoggerFromCtx(ctx)
	requestId := middleware.GetRequestIdFromCtx(ctx)

	var sendCodeReq models.SendCodeReq
	if err := ctx.Bind(&sendCodeReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, httpErrDescr.BAD_REQUEST_BODY)
	}
	fmt.Println(sendCodeReq)
	if _, err := govalidator.ValidateStruct(sendCodeReq); err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, httpErrDescr.INVALID_DATA)
	}
	isRegistered, err := h.Usecase.SendCode(&sendCodeReq)
	if err != nil {
		logger.Error(requestId, "error sending code: "+err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
	}
	return ctx.JSON(http.StatusOK, models.SendCodeResp{IsRegistered: isRegistered})
}
