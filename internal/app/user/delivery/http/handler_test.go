package userHandler

import (
	"encoding/json"
	"fmt"
	conf "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/config"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/domain/interfaces"
	data "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/domain/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/staticManager/localStaticManager"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUserHandler_Login(t *testing.T) {
	config := conf.NewConfig()
	err := conf.ReadConfigFile("../config/serv.toml", config)

	staticManager := localStaticManager.NewLocalFileManager(config.ServConfig.StaticUrl, config.ServConfig.StaticPath)

	mockUCase := new(interfaces.UserUseCase)
	mockAuthManager := new(interfaces.AuthManager)
	e := echo.New()

	userRegister := &models.RegisterReq{
		Phone: 	data.User.Phone,
		Code: "1234",
		Name:   data.User.Name,
		Email: 	data.User.Email,
	}

	j, err := json.Marshal(userRegister)

	req, err := http.NewRequest(echo.GET, "/login", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := UserHandler{
		Usecase:       mockUCase,
		AuthManager: mockAuthManager,
		StaticManager: staticManager,
	}

	json := `{"name":"Name","phone":"79166152595","email":"email@email.com","avatar":"avatar/avatar"}
`

	err = handler.Login(c)

	jj := rec.Body.String()
	fmt.Println(jj)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, json, rec.Body.String())
}

func TestUserHandler_Register(t *testing.T) {
	config := conf.NewConfig()
	err := conf.ReadConfigFile("../config/serv.toml", config)

	staticManager := localStaticManager.NewLocalFileManager(config.ServConfig.StaticUrl, config.ServConfig.StaticPath)

	mockUCase := new(interfaces.UserUseCase)
	mockAuthManager := new(interfaces.AuthManager)
	e := echo.New()

	userRegister := &models.RegisterReq{
		Phone: 	data.User.Phone,
		Code: "1234",
		Name:   data.User.Name,
		Email: 	data.User.Email,
	}

	j, err := json.Marshal(userRegister)

	req, err := http.NewRequest(echo.GET, "/login", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := UserHandler{
		Usecase:       mockUCase,
		AuthManager: mockAuthManager,
		StaticManager: staticManager,
	}

	json := `{"name":"Name","phone":"79166152595","email":"email@email.com","avatar":"avatar/avatar"}
`

	err = handler.Register(c)

	jj := rec.Body.String()
	fmt.Println(jj)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, json, rec.Body.String())
}

func TestUserHandler_SendCode(t *testing.T) {
	config := conf.NewConfig()
	err := conf.ReadConfigFile("../config/serv.toml", config)

	staticManager := localStaticManager.NewLocalFileManager(config.ServConfig.StaticUrl, config.ServConfig.StaticPath)

	mockUCase := new(interfaces.UserUseCase)
	mockAuthManager := new(interfaces.AuthManager)
	e := echo.New()

	codeReq := &models.SendCodeReq{
		Phone: "79166152595",
	}

	j, err := json.Marshal(codeReq)

	req, err := http.NewRequest(echo.GET, "/send_code", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := UserHandler{
		Usecase:       mockUCase,
		AuthManager: mockAuthManager,
		StaticManager: staticManager,
	}

	json := `{"registered":true}
`

	err = handler.SendCode(c)

	jj := rec.Body.String()
	fmt.Println(jj)

	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, json, rec.Body.String())
}

//func TestUserHandler_UpdateUser(t *testing.T) {
//	config := conf.NewConfig()
//	err := conf.ReadConfigFile("../config/serv.toml", config)
//
//	staticManager := localStaticManager.NewLocalFileManager(config.ServConfig.StaticUrl, config.ServConfig.StaticPath)
//
//	mockUCase := new(interfaces.UserUseCase)
//	mockAuthManager := new(interfaces.AuthManager)
//	e := echo.New()
//
//	codeReq := &models.UpdateUserReq{
//		Name:  "Name",
//		Email: "Email",
//	}
//
//	j, err := json.Marshal(codeReq)
//
//	req, err := http.NewRequest(echo.GET, "/send_code", strings.NewReader(string(j)))
//	assert.NoError(t, err)
//	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
//
//	rec := httptest.NewRecorder()
//	c := e.NewContext(req, rec)
//	handler := UserHandler{
//		Usecase:       mockUCase,
//		AuthManager: mockAuthManager,
//		StaticManager: staticManager,
//	}
//
//	json := `{"registered":true}
//`
//
//	err = handler.UpdateUser(c)
//
//	jj := rec.Body.String()
//	fmt.Println(jj)
//
//	require.NoError(t, err)
//	assert.Equal(t, http.StatusOK, rec.Code)
//	assert.Equal(t, json, rec.Body.String())
//}