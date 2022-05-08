package ucase

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/order"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/order/models"
	"github.com/pkg/errors"
)

type OrderUcase struct {
	OrderRepo order.Repository
}

func NewOrderUcase(orderRepo order.Repository) *OrderUcase {
	return &OrderUcase{
		OrderRepo: orderRepo,
	}
}

func (u *OrderUcase) CreateOrder(order *models.CreateOrderUcaseReq) (*models.CreateOrderUcaseResp, error) {
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

func (u *OrderUcase) GetUserOrders(user *models.GetUserOrdersUcaseReq) (*models.GetUserOrdersUcaseResp, error) {

	orders, err := u.OrderRepo.GetUserOrders(&models.GetUserOrdersRepoReq{UserId: user.UserId})

	if err != nil || orders == nil {
		return nil, errors.Wrap(err, "error getting orders from storage")
	}
	ordersResp := make([]models.ShortOrderUcase, len(orders.Orders))
	for i, position := range orders.Orders {
		ordersResp[i] = models.ShortOrderUcase(position)
	}
	return &models.GetUserOrdersUcaseResp{Orders: ordersResp}, nil
}

func (u *OrderUcase) GetUserOrderStatuses(user *models.GetUserOrderStatusesUcaseReq) (*models.GetUserOrderStatusesUcaseResp, error) {

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

func (u *OrderUcase) GetUserOrder(req *models.GetUserOrderUcaseReq) (*models.GetUserOrderUcaseResp, error) {

	order, err := u.OrderRepo.GetUserOrder(&models.GetUserOrderRepoReq{OrderId: req.OrderId})

	if err != nil {
		return nil, errors.Wrap(err, "error getting orders from storage")
	}

	if order.UserId != req.UserId {
		return nil, servErrors.NewError(servErrors.THIS_ORDER_DOESNOT_BELONG_USER, "")
	}
	cart := make([]models.OrderPositionUcaseResp, len(order.Cart))
	for i, poz := range order.Cart {
		cart[i] = models.OrderPositionUcaseResp{Name: poz.Name, Description: poz.Description, ImagePath: poz.ImagePath, Calories: poz.Calories, Count: poz.Count, Price: poz.Price, Weigth: poz.Weight}
	}

	return &models.GetUserOrderUcaseResp{OrderId: order.OrderId, Address: order.Address, Date: order.Date, RestaurantName: order.RestaurantName, RestaurantSlug: order.RestaurantSlug, TotalPrice: order.TotalPrice, Status: order.Status, Cart: cart}, nil
}
