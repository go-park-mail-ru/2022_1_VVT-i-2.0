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
	OK = iota
	CACH_MISS_CODE
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
	NO_SUCH_CITY
	NO_SUCH_STREET
	CANT_CREATE_AVATAR_NAME
	CANT_SAVE_AVATAR
	DECODE_IMG
	UNKNOWN_STATIC_TYPE
	TEST_ERROR
	THIS_ORDER_DOESNOT_BELONG_USER
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
	DB_INSERT_DESCR                      = "неуспешная вставка в базу данных"
	NO_SUCH_CITY_DESCR                   = "не существует такого города"
	NO_SUCH_STREET_DESCR                 = "не существует такой улицы"
	CANT_CREATE_AVATAR_NAME_DESCR        = "не получается сгенерировать имя аватарки"
	CANT_SAVE_AVATAR_DESCR               = "не получается сохранить аватарку"
	UNKNOWN_STATIC_TYPE_DESCR            = "неизвестный тип статики, невозможно определить расположение"
	THIS_ORDER_DOESNOT_BELONG_USER_DESCR = "запрошенный заказ не принадлежит текущему пользователю"
)

var codeDescr = map[int]string{
	NO_SUCH_RESTAURANT:             NO_SUCH_RESTAURANT_DESCR,
	CACH_MISS_CODE:                 CACH_MISS_DESCR,
	NO_SUCH_USER:                   NO_SUCH_USER_DESCR,
	THIS_ORDER_DOESNOT_BELONG_USER: THIS_ORDER_DOESNOT_BELONG_USER_DESCR,
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
