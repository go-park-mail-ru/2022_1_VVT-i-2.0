package usecase

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/order"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/order/models"
	"github.com/pkg/errors"
)

type OrderUsecase struct {
	OrderRepo order.Repository
}

func NewOrderUcase(orderRepo order.Repository) *OrderUsecase {
	return &OrderUsecase{
		OrderRepo: orderRepo,
	}
}

func (u *OrderUsecase) CreateOrder(order *models.CreateOrderUcaseReq) (*models.CreateOrderUcaseResp, error) {
	// TODO: сделать проверку, есть ли такой адрес
	cart := make([]models.OrderPositionRepo, len(order.Cart))
	for i, position := range order.Cart {
		cart[i] = models.OrderPositionRepo(position)
	}
	orderId, err := u.OrderRepo.CreateOrder(&models.CreateOrderRepoReq{UserId: order.UserId, Address: order.Address, Comment: order.Comment, Cart: cart})

	if err != nil || orderId.OrderId <= 0 {
		return nil, errors.Wrap(err, "error adding order to storage")
	}
	return &models.CreateOrderUcaseResp{OrderId: orderId.OrderId}, nil

}

func (u *OrderUsecase) GetUserOrders(user *models.GetUserOrdersUcaseReq) (*models.GetUserOrdersUcaseResp, error) {

	orders, err := u.OrderRepo.GetUserOrders(&models.GetUserOrdersRepoReq{UserId: user.UserId})

	if err != nil || orders == nil {
		return nil, errors.Wrap(err, "error getting orders from storage")
	}
	ordersResp := make([]models.ShortOrderUcase, len(orders.OrderStatuses))
	for i, position := range orders.OrderStatuses {
		ordersResp[i] = models.ShortOrderUcase(position)
	}
	return &models.GetUserOrdersUcaseResp{Orders: ordersResp}, nil
}

func (u *OrderUsecase) GetUserOrderStatuses(user *models.GetUserOrderStatusesUcaseReq) (*models.GetUserOrderStatusesUcaseResp, error) {

	orderStatuses, err := u.OrderRepo.GetUserOrderStatuses(&models.GetUserOrderStatusesRepoReq{UserId: user.UserId})

	if err != nil || orderStatuses == nil {
		return nil, errors.Wrap(err, "error getting orders from storage")
	}
	orderStatusesResp := make([]models.OrderStatusUcase, len(orderStatuses.OrderStatuses))
	for i, position := range orderStatuses.OrderStatuses {
		orderStatusesResp[i] = models.OrderStatusUcase(position)
	}
	return &models.GetUserOrderStatusesUcaseResp{OrderStatuses: orderStatusesResp}, nil
}
