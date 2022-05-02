package order

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/order/models"
)

type Usecase interface {
	CreateOrder(order *models.CreateOrderUcaseReq) (*models.CreateOrderUcaseResp, error)
	GetUserOrders(user *models.GetUserOrdersUcaseReq) (*models.GetUserOrdersUcaseResp, error)
	// GetUserOrder(user *models.GetUserOrderUcaseReq) (*models.GetUserOrderUcaseResp, error)
	GetUserOrderStatuses(user *models.GetUserOrderStatusesUcaseReq) (*models.GetUserOrderStatusesUcaseResp, error)
}
