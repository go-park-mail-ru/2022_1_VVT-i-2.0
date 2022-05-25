package mock

import (
	"errors"

	"github.com/stretchr/testify/mock"
)

type Notificator struct {
	mock.Mock
}

func (n *Notificator) SendCode(phone string, code string) error {
	return nil
}

type NotificatorErr struct {
	mock.Mock
}

func (n *NotificatorErr) SendCode(phone string, code string) error {
	return errors.New("unknown error")
}
