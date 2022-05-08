package mock

import "github.com/stretchr/testify/mock"

type Notificator struct {
	mock.Mock
}

func (n *Notificator) SendCode(phone string, code string) error {
	return nil
}
