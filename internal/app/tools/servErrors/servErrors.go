package servErrors

import (
	"fmt"

	"github.com/pkg/errors"
)

// ошибки, которые нужно залогировать
const (
	// -iota - 1
	CACH_ERROR = -iota - 1
)

// ошибки, которые не нужно логировать
const (
	CACH_MISS_CODE = iota

	WRONG_AUTH_CODE
	AUTH_REQUIRED
	NO_SUCH_RESTAURANT
	BAD_REQUEST_BODY
	SERVER_ERROR
	SENDING_AUTH_CODE

	// невалидные данные от пользователя
	INVALID_DATA
)

const (
	CACH_MISS_DESCR          = "в кэше отсутствует элемент по данному ключу"
	AUTH_REQUIRED_DESCR      = "для этого действия необходимо авторизоваться"
	BAD_AUTH_TOKEN_DESCR     = "не валидный токен авторизации"
	NO_SUCH_RESTAURANT_DESCR = "не существует такого ресторана"
	NO_SUCH_AUTH_DATA_DESCR  = "не найдены данные для авторизации пользователя"
	BAD_REQUEST_BODY_DESCR   = "не правильное тело запроса"
	SERVER_ERROR_DESCR       = "ошибка на сервера"
	WRONG_AUTH_CODE_DESCR    = "не верный код для входа"

	// невалидные данные от пользователя
	INVALID_DATA_DESCR = "переданы не валидные данные"
)

var codeDescr = map[int]string{
	AUTH_REQUIRED:      AUTH_REQUIRED_DESCR,
	NO_SUCH_RESTAURANT: NO_SUCH_RESTAURANT_DESCR,
	BAD_REQUEST_BODY:   BAD_REQUEST_BODY_DESCR,
	SERVER_ERROR:       SERVER_ERROR_DESCR,
	CACH_MISS_CODE:     CACH_MISS_DESCR,
	INVALID_DATA:       INVALID_DATA_DESCR,
}

type Error struct {
	Description string
	Code        int
}

func (e Error) Error() string {
	return fmt.Sprintf("error with code %d description: %s", e.Code, e.Description)
}

func (e *Error) Cause() error {
	return e
}

func NewError(eCode int, eDescr string) Error {
	if eDescr == "" {
		eDescr = codeDescr[eCode]
	}
	return Error{
		Code:        eCode,
		Description: eDescr,
	}
}

func ErrorAs(e error) *Error {
	var cause Error
	if ok := errors.As(e, &cause); ok {
		return &cause
	}
	return nil
}
