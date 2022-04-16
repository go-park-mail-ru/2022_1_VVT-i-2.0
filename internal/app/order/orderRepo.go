package order

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
)

type Repository interface {
	AddOrder(order *models.OrderRepoInput) (*models.OrderRepoAnsw, error)
}
