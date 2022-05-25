package userHandler

import (
	"net"
	"net/http"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/delivery/http/httpErrDescr"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/delivery/http/middleware"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/authManager"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/staticManager"
	_ "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/validator"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/user"
	"github.com/labstack/echo/v4"
)

const (
	tokenCookieKey    = "token"
	CSRFCookieName    = "_csrf"
	avatarMaxSize     = 4000000
	updateUserMaxSize = 1000
)

type UserHandler struct {
	Ucase         user.Ucase
	AuthManager   authManager.AuthManager
	StaticManager staticManager.FileManager
}

func NewUserHandler(ucase user.Ucase, authManager authManager.AuthManager, staticManager staticManager.FileManager) *UserHandler {
	return &UserHandler{
		Ucase:         ucase,
		AuthManager:   authManager,
		StaticManager: staticManager,
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
		return httpErrDescr.NewHTTPError(ctx, http.StatusConflict, httpErrDescr.ALREADY_AUTHORIZED)
	}

	logger := middleware.GetLoggerFromCtx(ctx)
	requestId := middleware.GetRequestIdFromCtx(ctx)

	var loginReq models.LoginReq
	if err := ctx.Bind(&loginReq); err != nil {
		return httpErrDescr.NewHTTPError(ctx, http.StatusBadRequest, httpErrDescr.BAD_REQUEST_BODY)
	}
	if _, err := govalidator.ValidateStruct(loginReq); err != nil {
		return httpErrDescr.NewHTTPError(ctx, http.StatusBadRequest, httpErrDescr.INVALID_DATA)
	}

	userDataUcase, err := h.Ucase.Login(&models.LoginUcaseReq{Phone: loginReq.Phone, Code: loginReq.Code})
	if err != nil {
		cause := servErrors.ErrorAs(err)
		if cause == nil {
			logger.Error(requestId, err.Error())
			return httpErrDescr.NewHTTPError(ctx, http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
		}
		switch cause.Code {
		case servErrors.WRONG_AUTH_CODE:
			return httpErrDescr.NewHTTPError(ctx, http.StatusForbidden, httpErrDescr.WRONG_AUTH_CODE)
		case servErrors.CACH_MISS_CODE:
			return httpErrDescr.NewHTTPError(ctx, http.StatusNotFound, httpErrDescr.NO_SUCH_CODE_INFO)
		case servErrors.NO_SUCH_ENTITY_IN_DB:
			return httpErrDescr.NewHTTPError(ctx, http.StatusNotFound, httpErrDescr.NO_SUCH_USER)
		default:
			logger.Error(requestId, err.Error())
			return httpErrDescr.NewHTTPError(ctx, http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
		}
	}

	token, err := h.AuthManager.CreateToken(authManager.NewTokenPayload(userDataUcase.Id))
	if err != nil {
		logger.Error(requestId, "error creating token: "+err.Error())
		return httpErrDescr.NewHTTPError(ctx, http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
	}

	host, _, _ := net.SplitHostPort(ctx.Request().Host)
	tokenCookie := createTokenCookie(token, host, h.AuthManager.GetEpiryTime())

	ctx.SetCookie(tokenCookie)
	csrfToken := middleware.GetCSRFTokenromCtx(ctx)
	if csrfToken != "" {
		ctx.Response().Header().Add(echo.HeaderXCSRFToken, csrfToken)
	}
	return ctx.JSON(http.StatusOK, models.LoginResp{Phone: userDataUcase.Phone, Email: userDataUcase.Email, Name: userDataUcase.Name, Avatar: h.StaticManager.GetAvatarUrl(userDataUcase.Avatar), Address: userDataUcase.Address})
}

func (h UserHandler) Register(ctx echo.Context) error {

	if middleware.GetUserFromCtx(ctx) != nil {
		return httpErrDescr.NewHTTPError(ctx, http.StatusConflict, httpErrDescr.ALREADY_AUTHORIZED)
	}

	logger := middleware.GetLoggerFromCtx(ctx)
	requestId := middleware.GetRequestIdFromCtx(ctx)

	var registerReq models.RegisterReq

	if err := ctx.Bind(&registerReq); err != nil {
		return httpErrDescr.NewHTTPError(ctx, http.StatusBadRequest, httpErrDescr.BAD_REQUEST_BODY)
	}
	if _, err := govalidator.ValidateStruct(registerReq); err != nil {
		return httpErrDescr.NewHTTPError(ctx, http.StatusBadRequest, httpErrDescr.INVALID_DATA)
	}

	userDataUcase, err := h.Ucase.Register(&models.RegisterUcaseReq{Phone: registerReq.Phone, Code: registerReq.Code, Name: registerReq.Name, Email: registerReq.Email})
	if err != nil {
		cause := servErrors.ErrorAs(err)
		if cause == nil {
			logger.Error(requestId, err.Error())
			return httpErrDescr.NewHTTPError(ctx, http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
		}
		switch cause.Code {
		case servErrors.WRONG_AUTH_CODE:
			return httpErrDescr.NewHTTPError(ctx, http.StatusForbidden, httpErrDescr.WRONG_AUTH_CODE)
		case servErrors.CACH_MISS_CODE:
			return httpErrDescr.NewHTTPError(ctx, http.StatusNotFound, httpErrDescr.NO_SUCH_CODE_INFO)
		case servErrors.DB_INSERT:
			return httpErrDescr.NewHTTPError(ctx, http.StatusConflict, httpErrDescr.SUCH_USER_ALREADY_EXISTS)
		default:
			logger.Error(requestId, err.Error())
			return httpErrDescr.NewHTTPError(ctx, http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
		}
	}

	token, err := h.AuthManager.CreateToken(authManager.NewTokenPayload(userDataUcase.Id))
	if err != nil {
		logger.Error(requestId, "error creating token: "+err.Error())
		return httpErrDescr.NewHTTPError(ctx, http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
	}

	host, _, _ := net.SplitHostPort(ctx.Request().Host)
	tokenCookie := createTokenCookie(token, host, h.AuthManager.GetEpiryTime())

	ctx.SetCookie(tokenCookie)
	csrfToken := middleware.GetCSRFTokenromCtx(ctx)
	if csrfToken != "" {
		ctx.Response().Header().Add(echo.HeaderXCSRFToken, csrfToken)
	}
	return ctx.JSON(http.StatusOK, models.UserDataResp{Phone: userDataUcase.Phone, Email: userDataUcase.Email, Name: userDataUcase.Name, Avatar: h.StaticManager.GetAvatarUrl(userDataUcase.Avatar)})
}

func (h UserHandler) Logout(ctx echo.Context) error {
	if middleware.GetUserFromCtx(ctx) == nil {
		return httpErrDescr.NewHTTPError(ctx, http.StatusUnauthorized, httpErrDescr.AUTH_REQUIRED)
	}
	host, _, _ := net.SplitHostPort(ctx.Request().Host)
	resetTokenCookie := createTokenCookie("", host, -time.Hour)

	resetCsrfCookie := &http.Cookie{
		Name:    CSRFCookieName,
		Expires: time.Now().Add(-time.Hour),
		Domain:  host,
		Path:    "/",
	}

	ctx.SetCookie(resetTokenCookie)
	ctx.SetCookie(resetCsrfCookie)
	return ctx.NoContent(http.StatusOK)
}

func (h UserHandler) SendCode(ctx echo.Context) error {
	if middleware.GetUserFromCtx(ctx) != nil {
		return httpErrDescr.NewHTTPError(ctx, http.StatusConflict, httpErrDescr.ALREADY_AUTHORIZED)
	}

	logger := middleware.GetLoggerFromCtx(ctx)
	requestId := middleware.GetRequestIdFromCtx(ctx)

	var sendCodeReq models.SendCodeReq
	if err := ctx.Bind(&sendCodeReq); err != nil {
		return httpErrDescr.NewHTTPError(ctx, http.StatusBadRequest, httpErrDescr.BAD_REQUEST_BODY)
	}
	if _, err := govalidator.ValidateStruct(sendCodeReq); err != nil {
		return httpErrDescr.NewHTTPError(ctx, http.StatusBadRequest, httpErrDescr.INVALID_DATA)
	}
	isRegistered, err := h.Ucase.SendCode(&models.SendCodeUcaseReq{Phone: sendCodeReq.Phone})
	if err != nil {
		logger.Error(requestId, "error sending code: "+err.Error())
		return httpErrDescr.NewHTTPError(ctx, http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
	}
	return ctx.JSON(http.StatusOK, models.SendCodeResp(isRegistered))
}

func (h UserHandler) GetUser(ctx echo.Context) error {
	user := middleware.GetUserFromCtx(ctx)
	if user == nil {
		return ctx.JSON(http.StatusOK, ``)
	}

	logger := middleware.GetLoggerFromCtx(ctx)
	requestId := middleware.GetRequestIdFromCtx(ctx)

	userDataUcase, err := h.Ucase.GetUser(user.Id)

	if err != nil {
		cause := servErrors.ErrorAs(err)
		if cause != nil && cause.Code == servErrors.NO_SUCH_ENTITY_IN_DB {
			return httpErrDescr.NewHTTPError(ctx, http.StatusForbidden, httpErrDescr.NO_SUCH_USER)
		}
		logger.Error(requestId, err.Error())
		return httpErrDescr.NewHTTPError(ctx, http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
	}

	return ctx.JSON(http.StatusOK, models.UserDataResp{Phone: userDataUcase.Phone, Email: userDataUcase.Email, Name: userDataUcase.Name, Avatar: h.StaticManager.GetAvatarUrl(userDataUcase.Avatar)})
}

func (h UserHandler) UpdateUser(ctx echo.Context) error {
	user := middleware.GetUserFromCtx(ctx)
	if user == nil {
		return httpErrDescr.NewHTTPError(ctx, http.StatusUnauthorized, httpErrDescr.AUTH_REQUIRED)
	}

	err := ctx.Request().ParseMultipartForm(avatarMaxSize + updateUserMaxSize)
	if err != nil {
		return httpErrDescr.NewHTTPError(ctx, http.StatusBadRequest, httpErrDescr.BAD_REQUEST_BODY)
	}
	logger := middleware.GetLoggerFromCtx(ctx)
	requestId := middleware.GetRequestIdFromCtx(ctx)

	updateReq := models.UpdateUserReq{
		Name:  ctx.Request().FormValue("name"),
		Email: ctx.Request().FormValue("email"),
	}

	if _, err := govalidator.ValidateStruct(updateReq); err != nil || (updateReq.Email == "" && updateReq.Name == "") {
		return httpErrDescr.NewHTTPError(ctx, http.StatusBadRequest, httpErrDescr.INVALID_DATA)
	}

	avatarImage, _, _ := ctx.Request().FormFile("avatar")
	if avatarImage != nil {
		defer avatarImage.Close()
	}

	userDataUcase, err := h.Ucase.UpdateUser(&models.UpdateUserUcase{Id: user.Id, Email: updateReq.Email, Name: updateReq.Name, AvatarImg: avatarImage})

	if err != nil {
		cause := servErrors.ErrorAs(err)
		if cause == nil {
			logger.Error(requestId, err.Error())
			return httpErrDescr.NewHTTPError(ctx, http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
		}
		switch cause.Code {
		case servErrors.DB_UPDATE:
			return httpErrDescr.NewHTTPError(ctx, http.StatusConflict, httpErrDescr.SUCH_USER_ALREADY_EXISTS)
		case servErrors.DECODE_IMG:
			return httpErrDescr.NewHTTPError(ctx, http.StatusBadRequest, httpErrDescr.BAD_IMAGE)
		}
		logger.Error(requestId, err.Error())
		return httpErrDescr.NewHTTPError(ctx, http.StatusInternalServerError, httpErrDescr.SERVER_ERROR)
	}

	csrfToken := middleware.GetCSRFTokenromCtx(ctx)
	if csrfToken != "" {
		ctx.Response().Header().Add(echo.HeaderXCSRFToken, csrfToken)
	}

	return ctx.JSON(http.StatusOK, models.UserDataResp{Phone: userDataUcase.Phone, Email: userDataUcase.Email, Name: userDataUcase.Name, Avatar: h.StaticManager.GetAvatarUrl(userDataUcase.Avatar)})
}
