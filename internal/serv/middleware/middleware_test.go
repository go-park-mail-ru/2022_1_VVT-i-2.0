package middleware

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func testHandlerFuncPanic(w http.ResponseWriter, r *http.Request) {
	panic("test panic")
}

func TestPanicMiddleware(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("the panic was not canceled by middleware")
		}
	}()

	logger := zap.NewExample()
	defer logger.Sync()
	sugarLogger := logger.Sugar()

	r := mux.NewRouter()
	r.HandleFunc("/test", testHandlerFuncPanic).Methods("GET")
	r.Use(Logger{Logger: sugarLogger}.PanicMiddleware)
	w := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/test", bytes.NewBuffer(nil))
	require.NoError(t, err)

	r.ServeHTTP(w, req)
}
