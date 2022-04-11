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
	FLASHCALL_RESPONSE_ERR

	WRONG_AUTH_CODE
	NO_SUCH_RESTAURANT
	NO_SUCH_USER
	SENDING_AUTH_CODE
	CREATE_TOKEN
	PARSE_TOKEN
	WRONG_TOKEN_CLAIMS
	// невалидные данные от пользователя
	NO_SUCH_ENTITY_IN_DB
	INVALID_DATA
	DB_ERROR
	DB_INSERT
	DB_UPDATE
)

const (
	DB_ERROR_DESCR             = "ошибка при работе с базой данных"
	CACH_MISS_DESCR            = "в кэше отсутствует элемент по данному ключу"
	BAD_AUTH_TOKEN_DESCR       = "не валидный токен авторизации"
	NO_SUCH_RESTAURANT_DESCR   = "не существует такого ресторана"
	NO_SUCH_AUTH_DATA_DESCR    = "не найдены данные для авторизации пользователя"
	BAD_REQUEST_BODY_DESCR     = "не правильное тело запроса"
	WRONG_AUTH_CODE_DESCR      = "не верный код для входа"
	NO_SUCH_ENTITY_IN_DB_DESCR = "нет такой сущности в базе данных"
	NO_SUCH_USER_DESCR         = "не существует такого пользователя"
	// невалидные данные от пользователя
	DB_INSERT_DESCR = "неуспешная вставка в базу данных"
)

var codeDescr = map[int]string{
	NO_SUCH_RESTAURANT: NO_SUCH_RESTAURANT_DESCR,
	CACH_MISS_CODE:     CACH_MISS_DESCR,
	NO_SUCH_USER:       NO_SUCH_USER_DESCR,
}

type Error struct {
	Description string
	Code        int
}

func (e Error) Error() string {
	return fmt.Sprintf("error with code %d description: %s", e.Code, e.Description)
}

func (e Error) Cause() error {
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

// TODO потестить норм ли работает ас
