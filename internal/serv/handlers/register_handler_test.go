package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	models "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/serv/models"
)

func TestRegister(t *testing.T) {
	type testCase struct {
		Request         models.RegisterRequest
		Response        string
		StatusCode      int
		CaseDescription string
	}

	testCases := []testCase{
		{
			Request:         models.RegisterRequest{Phone: "+7(999)999-99-99", Password: "Qw111111", Username: "Наташа"},
			Response:        `{"username":"Наташа"}`,
			StatusCode:      http.StatusOK,
			CaseDescription: "valid",
		},
		{
			Request:         models.RegisterRequest{Phone: "+7(999)999-99-99", Password: "Qw111111", Username: "Наташа"},
			Response:        `{"error":"such user already exists"}`,
			StatusCode:      http.StatusConflict,
			CaseDescription: "such phone already exists",
		},
		{
			Request:         models.RegisterRequest{Phone: "+7(999)999-99-99", Password: "Qw111111", Username: "Наташ7"},
			Response:        `{"error":"not valid username"}`,
			StatusCode:      http.StatusBadRequest,
			CaseDescription: "invalid username",
		},
		{
			Request:         models.RegisterRequest{Phone: "+7(999)999-99-999", Password: "Qw111111", Username: "Наташа"},
			Response:        `{"error":"not valid phone"}`,
			StatusCode:      http.StatusBadRequest,
			CaseDescription: "invalid phone",
		},
		{
			Request:         models.RegisterRequest{Phone: "+7(999)999-99-99", Password: "Qw111 111", Username: "Наташа"},
			Response:        `{"error":"not valid password"}`,
			StatusCode:      http.StatusBadRequest,
			CaseDescription: "invalid password",
		},
	}

	for _, testCase := range testCases {
		url := "http://example.com/api/v1/register"
		payloadBuf, _ := json.Marshal(testCase.Request)

		req := httptest.NewRequest(http.MethodPost, url, strings.NewReader(string(payloadBuf)))
		w := httptest.NewRecorder()

		RegisterHandler(w, req)

		if w.Code != testCase.StatusCode {
			t.Errorf("%s wrong StatusCode: got %d, expected %d",
				testCase.CaseDescription, w.Code, testCase.StatusCode)
		}

		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)

		bodyStr := string(body)
		if bodyStr != testCase.Response+"\n" {
			t.Errorf("[%s] : wrong Response: got %+v, expected %+v",
				testCase.CaseDescription, bodyStr, testCase.Response)
		}
	}
}

func TestLogin(t *testing.T) {
	type testCase struct {
		Request         models.LoginRequest
		Response        string
		StatusCode      int
		CaseDescription string
	}

	testCases := []testCase{
		{
			Request:         models.LoginRequest{Phone: "+7(900)555-35-35", Password: "qw12qqqq"},
			Response:        `{"username":"Наташа","userAddress":"Москва, Петровка 38"}`,
			StatusCode:      http.StatusOK,
			CaseDescription: "valid",
		},
		{
			Request:         models.LoginRequest{Phone: "+7(000)000-00-00", Password: "Qw111111"},
			Response:        `{"error":"no such user"}` + "\n" + `{}`,
			StatusCode:      http.StatusForbidden,
			CaseDescription: "no such user",
		},
		{
			Request:         models.LoginRequest{Phone: "+7(999)999-99-999", Password: "Qw111111"},
			Response:        `{"error":"not valid phone"}`,
			StatusCode:      http.StatusBadRequest,
			CaseDescription: "invalid phone",
		},
		{
			Request:         models.LoginRequest{Phone: "+7(999)999-99-99", Password: "Qw111 111"},
			Response:        `{"error":"not valid password"}`,
			StatusCode:      http.StatusBadRequest,
			CaseDescription: "invalid password",
		},
	}

	for _, testCase := range testCases {
		url := "http://example.com/api/v1/login"
		payloadBuf, _ := json.Marshal(testCase.Request)

		req := httptest.NewRequest(http.MethodPost, url, strings.NewReader(string(payloadBuf)))
		w := httptest.NewRecorder()

		LoginHandler(w, req)

		if w.Code != testCase.StatusCode {
			t.Errorf("%s wrong StatusCode: got %d, expected %d",
				testCase.CaseDescription, w.Code, testCase.StatusCode)
		}

		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)

		bodyStr := string(body)
		if bodyStr != testCase.Response+"\n" {
			t.Errorf("[%s] : wrong Response: got %+v, expected %+v",
				testCase.CaseDescription, bodyStr, testCase.Response)
		}
	}
}
