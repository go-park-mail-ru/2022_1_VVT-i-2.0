package userHandler

// import (
// 	"bytes"
// 	"encoding/json"
// 	"mime/multipart"
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"

// 	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/delivery/http/middleware"
// 	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
// 	mockAuth "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/authManager/mock"
// 	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/logger"
// 	mockLogger "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/logger/mock"
// 	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/staticManager/localStaticManager"
// 	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/user/mock"
// 	"github.com/labstack/echo/v4"
// 	"github.com/stretchr/testify/assert"
// )

// func TestUserHandler_Login(t *testing.T) {
// 	staticManager := localStaticManager.NewLocalFileManager("", "")

// 	mockUCase := new(mock.UserUcase)
// 	mockAuthManager := new(mockAuth.AuthManager)
// 	e := echo.New()

// 	userRegister := &models.RegisterReq{
// 		Phone: "79999999999",
// 		Code:  "1234",
// 		Name:  "Name",
// 		Email: "email@mail.com",
// 	}

// 	j, _ := json.Marshal(userRegister)

// 	req, err := http.NewRequest(echo.POST, "/login", strings.NewReader(string(j)))
// 	assert.NoError(t, err)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)
// 	handler := UserHandler{
// 		Ucase:         mockUCase,
// 		AuthManager:   mockAuthManager,
// 		StaticManager: staticManager,
// 	}

// 	json := `{"name":"Name","phone":"79999999999","email":"email@mail.com","avatar":"avatar/avatar.png"}
// `
// 	err = handler.Login(c)

// 	assert.NoError(t, err)
// 	assert.Equal(t, http.StatusOK, rec.Code)
// 	assert.Equal(t, json, rec.Body.String())
// }

// func TestUserHandler_Register(t *testing.T) {
// 	staticManager := localStaticManager.NewLocalFileManager("", "")

// 	mockUCase := new(mock.UserUcase)
// 	mockAuthManager := new(mockAuth.AuthManager)
// 	e := echo.New()

// 	userRegister := &models.RegisterReq{
// 		Phone: "79999999999",
// 		Code:  "1234",
// 		Name:  "Name",
// 		Email: "email@mail.com",
// 	}

// 	j, _ := json.Marshal(userRegister)

// 	req, err := http.NewRequest(echo.POST, "/login", strings.NewReader(string(j)))
// 	assert.NoError(t, err)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)
// 	handler := NewUserHandler(mockUCase, mockAuthManager, staticManager)

// 	json := `{"name":"Name","phone":"79999999999","email":"email@mail.com","avatar":"avatar/avatar.png"}
// `

// 	err = handler.Register(c)

// 	assert.NoError(t, err)
// 	assert.Equal(t, http.StatusOK, rec.Code)
// 	assert.Equal(t, json, rec.Body.String())
// }

// func TestUserHandler_GetUser(t *testing.T) {
// 	staticManager := localStaticManager.NewLocalFileManager("", "")

// 	mockUCase := new(mock.UserUcase)
// 	mockAuthManager := new(mockAuth.AuthManager)
// 	e := echo.New()

// 	req, err := http.NewRequest(echo.GET, "/user", strings.NewReader(""))
// 	assert.NoError(t, err)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)
// 	c.Set(middleware.UserCtxKey, middleware.UserCtx{Id: 1})
// 	handler := NewUserHandler(mockUCase, mockAuthManager, staticManager)

// 	json := `{"name":"Name","phone":"79999999999","email":"email@mail.com","avatar":"avatar/avatar.png"}
// `

// 	err = handler.GetUser(c)

// 	assert.NoError(t, err)
// 	assert.Equal(t, http.StatusOK, rec.Code)
// 	assert.Equal(t, json, rec.Body.String())
// }

// func TestUserHandler_SendCode(t *testing.T) {
// 	staticManager := localStaticManager.NewLocalFileManager("", "")

// 	mockUCase := new(mock.UserUcase)
// 	mockAuthManager := new(mockAuth.AuthManager)
// 	e := echo.New()

// 	codeReq := &models.SendCodeReq{
// 		Phone: "79999999999",
// 	}

// 	j, _ := json.Marshal(codeReq)

// 	req, err := http.NewRequest(echo.POST, "/send_code", strings.NewReader(string(j)))
// 	assert.NoError(t, err)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)
// 	handler := UserHandler{
// 		Ucase:         mockUCase,
// 		AuthManager:   mockAuthManager,
// 		StaticManager: staticManager,
// 	}

// 	json := `{"registered":true}
// `
// 	err = handler.SendCode(c)

// 	assert.NoError(t, err)
// 	assert.Equal(t, http.StatusOK, rec.Code)
// 	assert.Equal(t, json, rec.Body.String())
// }

// func TestUserHandler_Logout(t *testing.T) {
// 	staticManager := localStaticManager.NewLocalFileManager("", "")

// 	mockUCase := new(mock.UserUcase)
// 	mockAuthManager := new(mockAuth.AuthManager)
// 	e := echo.New()

// 	req, err := http.NewRequest(echo.GET, "/logout", strings.NewReader(""))
// 	assert.NoError(t, err)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)
// 	c.Set(middleware.UserCtxKey, middleware.UserCtx{Id: 1})
// 	handler := UserHandler{
// 		Ucase:         mockUCase,
// 		AuthManager:   mockAuthManager,
// 		StaticManager: staticManager,
// 	}

// 	err = handler.Logout(c)

// 	assert.NoError(t, err)
// 	assert.Equal(t, http.StatusOK, rec.Code)
// }

// func TestUserHandler_UpdateUser(t *testing.T) {
// 	staticManager := localStaticManager.NewLocalFileManager("", "")

// 	mockUCase := new(mock.UserUcase)
// 	mockAuthManager := new(mockAuth.AuthManager)
// 	e := echo.New()

// 	body := new(bytes.Buffer)
// 	writer := multipart.NewWriter(body)
// 	part, err := writer.CreateFormField("name")
// 	assert.NoError(t, err)

// 	part.Write([]byte("Name"))
// 	assert.NoError(t, err)
// 	assert.NoError(t, writer.Close())

// 	req, err := http.NewRequest(echo.POST, "/update", body)
// 	req.Header.Set(echo.HeaderContentType, writer.FormDataContentType())

// 	assert.NoError(t, err)

// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)
// 	handler := UserHandler{
// 		Ucase:         mockUCase,
// 		AuthManager:   mockAuthManager,
// 		StaticManager: staticManager,
// 	}

// 	json := `{"name":"Name","phone":"79999999999","email":"email@mail.com","avatar":"avatar/avatar.png"}
// `
// 	c.Set(middleware.UserCtxKey, middleware.UserCtx{Id: 1})
// 	err = handler.UpdateUser(c)

// 	assert.NoError(t, err)
// 	assert.Equal(t, http.StatusOK, rec.Code)
// 	assert.Equal(t, json, rec.Body.String())
// }

// func TestUserHandler_Login_Err(t *testing.T) {
// 	staticManager := localStaticManager.NewLocalFileManager("", "")

// 	mockUCase := new(mock.UserUcaseErr)
// 	mockAuthManager := new(mockAuth.AuthManager)
// 	mockLogger := new(mockLogger.Logger)
// 	e := echo.New()

// 	userRegister := &models.RegisterReq{
// 		Phone: "79999999999",
// 		Code:  "1234",
// 		Name:  "Name",
// 		Email: "email@mail.com",
// 	}

// 	j, _ := json.Marshal(userRegister)

// 	req, err := http.NewRequest(echo.GET, "/login", strings.NewReader(string(j)))
// 	assert.NoError(t, err)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)

// 	c.Set(middleware.LoggerCtxKey, &logger.ServLogger{Logger: mockLogger})
// 	handler := UserHandler{
// 		Ucase:         mockUCase,
// 		AuthManager:   mockAuthManager,
// 		StaticManager: staticManager,
// 	}

// 	err = handler.Login(c)

// 	assert.Equal(t, http.StatusInternalServerError, c.Response().Status)
// 	assert.Error(t, err)
// }

// func TestUserHandler_Register_Err(t *testing.T) {
// 	staticManager := localStaticManager.NewLocalFileManager("", "")

// 	mockUCase := new(mock.UserUcaseErr)
// 	mockAuthManager := new(mockAuth.AuthManager)
// 	mockLogger := new(mockLogger.Logger)
// 	e := echo.New()

// 	userRegister := &models.RegisterReq{
// 		Phone: "79999999999",
// 		Code:  "1234",
// 		Name:  "Name",
// 		Email: "email@mail.com",
// 	}

// 	j, _ := json.Marshal(userRegister)

// 	req, err := http.NewRequest(echo.GET, "/login", strings.NewReader(string(j)))
// 	assert.NoError(t, err)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)

// 	c.Set(middleware.LoggerCtxKey, &logger.ServLogger{Logger: mockLogger})
// 	handler := UserHandler{
// 		Ucase:         mockUCase,
// 		AuthManager:   mockAuthManager,
// 		StaticManager: staticManager,
// 	}

// 	err = handler.Register(c)

// 	assert.Error(t, err)
// 	assert.Equal(t, http.StatusInternalServerError, c.Response().Status)
// }

// func TestUserHandler_SendCode_Err(t *testing.T) {
// 	staticManager := localStaticManager.NewLocalFileManager("", "")

// 	mockUCase := new(mock.UserUcaseErr)
// 	mockAuthManager := new(mockAuth.AuthManager)
// 	mockLogger := new(mockLogger.Logger)
// 	e := echo.New()

// 	codeReq := &models.SendCodeReq{
// 		Phone: "79999999999",
// 	}

// 	j, _ := json.Marshal(codeReq)

// 	req, err := http.NewRequest(echo.GET, "/send_code", strings.NewReader(string(j)))
// 	assert.NoError(t, err)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)
// 	c.Set(middleware.LoggerCtxKey, &logger.ServLogger{Logger: mockLogger})
// 	handler := UserHandler{
// 		Ucase:         mockUCase,
// 		AuthManager:   mockAuthManager,
// 		StaticManager: staticManager,
// 	}

// 	err = handler.SendCode(c)

// 	assert.Error(t, err)
// 	assert.Equal(t, http.StatusInternalServerError, c.Response().Status)
// }

// func TestUserHandler_UpdateUser_Err(t *testing.T) {
// 	staticManager := localStaticManager.NewLocalFileManager("", "")

// 	mockUCase := new(mock.UserUcaseErr)
// 	mockAuthManager := new(mockAuth.AuthManager)
// 	mockLogger := new(mockLogger.Logger)
// 	e := echo.New()

// 	body := new(bytes.Buffer)
// 	writer := multipart.NewWriter(body)
// 	part, err := writer.CreateFormField("name")
// 	assert.NoError(t, err)

// 	part.Write([]byte("Name"))
// 	assert.NoError(t, err)
// 	assert.NoError(t, writer.Close())

// 	req, err := http.NewRequest(echo.POST, "/update", body)
// 	req.Header.Set(echo.HeaderContentType, writer.FormDataContentType())

// 	assert.NoError(t, err)

// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)
// 	c.Set(middleware.LoggerCtxKey, &logger.ServLogger{Logger: mockLogger})
// 	handler := UserHandler{
// 		Ucase:         mockUCase,
// 		AuthManager:   mockAuthManager,
// 		StaticManager: staticManager,
// 	}

// 	c.Set(middleware.UserCtxKey, middleware.UserCtx{Id: 1})
// 	err = handler.UpdateUser(c)

// 	assert.Error(t, err)
// 	assert.Equal(t, http.StatusInternalServerError, c.Response().Status)
// }

// func TestUserHandler_GetUser_Err(t *testing.T) {
// 	staticManager := localStaticManager.NewLocalFileManager("", "")

// 	mockUCase := new(mock.UserUcaseErr)
// 	mockAuthManager := new(mockAuth.AuthManager)
// 	mockLogger := new(mockLogger.Logger)
// 	e := echo.New()

// 	req, err := http.NewRequest(echo.GET, "/user", strings.NewReader(""))
// 	assert.NoError(t, err)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)
// 	c.Set(middleware.UserCtxKey, middleware.UserCtx{Id: 1})
// 	c.Set(middleware.LoggerCtxKey, &logger.ServLogger{Logger: mockLogger})
// 	handler := NewUserHandler(mockUCase, mockAuthManager, staticManager)

// 	err = handler.GetUser(c)

// 	assert.Error(t, err)
// 	assert.Equal(t, http.StatusInternalServerError, c.Response().Status)
// }
