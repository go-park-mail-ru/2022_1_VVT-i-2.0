package order

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
)

type Usecase interface {
	Order(order *models.OrderUcaseInput) (*models.OrderUcaseAnsw, error)
}
