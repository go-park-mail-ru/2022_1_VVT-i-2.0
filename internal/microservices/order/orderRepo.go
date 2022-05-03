package order

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/order/models"
)

type Repository interface {
	CreateOrder(order *models.CreateOrderRepoReq) (*models.CreateOrderRepoResp, error)
	GetUserOrders(user *models.GetUserOrdersRepoReq) (*models.GetUserOrdersRepoResp, error)
	GetUserOrder(user *models.GetUserOrderRepoReq) (*models.GetUserOrderRepoResp, error)
	GetUserOrderStatuses(user *models.GetUserOrderStatusesRepoReq) (*models.GetUserOrderStatusesRepoResp, error)
}
