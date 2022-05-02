package order

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
)

type Usecase interface {
	CreateOrder(order *models.OrderUcaseReq) (*models.OrderUcaseResp, error)
	GetUserOrders(order *models.GetUserOrdersUcaseReq) (*models.GetUserOrdersUcaseResp, error)
	GetUserOrderStatuses(order *models.GetUserOrderStatusesUcaseReq) (*models.GetUserOrderStatusesUcaseResp, error)
}
