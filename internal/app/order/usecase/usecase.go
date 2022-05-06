package usecase

import (
	"context"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	orderProto "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/order/proto"
	"google.golang.org/grpc/status"
)

type OrderUsecase struct {
	OrderCli orderProto.OrderServiceClient
}

func NewUsecase(orderCli orderProto.OrderServiceClient) *OrderUsecase {
	return &OrderUsecase{
		OrderCli: orderCli,
	}
}

func (u *OrderUsecase) CreateOrder(order *models.OrderUcaseReq) (*models.OrderUcaseResp, error) {
	cart := make([]*orderProto.OrderPositionReq, len(order.Cart))
	for i, position := range order.Cart {
		cart[i] = &orderProto.OrderPositionReq{Id: position.Id, Count: position.Count}
	}
	orderResp, err := u.OrderCli.CreateOrder(context.Background(), &orderProto.CreateOrderReq{UserId: int64(order.UserId), Address: order.Address, Comment: order.Comment, Cart: cart})
	if err != nil {
		return nil, servErrors.NewError(int(status.Code(err)), err.Error())
	}
	return &models.OrderUcaseResp{OrderId: orderResp.OrderId}, err
}

func (u *OrderUsecase) GetUserOrders(order *models.GetUserOrdersUcaseReq) (*models.GetUserOrdersUcaseResp, error) {
	orders, err := u.OrderCli.GetUserOrders(context.Background(), &orderProto.GetUserOrdersReq{UserId: int64(order.UserId)})

	if err != nil {
		return nil, servErrors.NewError(int(status.Code(err)), err.Error())
	}

	ordersResp := make([]models.ShortOrderUcase, len(orders.Orders))
	for i, order := range orders.Orders {
		ordersResp[i] = models.ShortOrderUcase{OrderId: order.OrderId, Date: order.Date, Status: order.Status, RestaurantName: order.RestaurantName, TotalPrice: order.TotalPrice}
	}
	return &models.GetUserOrdersUcaseResp{Orders: ordersResp}, err
}

func (u *OrderUsecase) GetUserOrderStatuses(order *models.GetUserOrderStatusesUcaseReq) (*models.GetUserOrderStatusesUcaseResp, error) {
	orders, err := u.OrderCli.GetUserOrderStatuses(context.Background(), &orderProto.GetUserOrderStatusesReq{UserId: int64(order.UserId)})

	if err != nil {
		return nil, servErrors.NewError(int(status.Code(err)), err.Error())
	}

	ordersResp := make([]models.OrderStatusUcase, len(orders.OrderStatuses))
	for i, order := range orders.OrderStatuses {
		ordersResp[i] = models.OrderStatusUcase{OrderId: order.OrderId, Status: order.Status}
	}
	return &models.GetUserOrderStatusesUcaseResp{OrderStatuses: ordersResp}, err
}

func (u *OrderUsecase) GetUserOrder(req *models.GetUserOrderUcaseReq) (*models.GetUserOrderUcaseResp, error) {
	order, err := u.OrderCli.GetUserOrder(context.Background(), &orderProto.GetUserOrderReq{UserId: int64(req.UserId), OrderId: req.OrderId})

	if err != nil {
		return nil, servErrors.NewError(int(status.Code(err)), err.Error())
	}

	resp := models.GetUserOrderUcaseResp{OrderId: order.OrderId, Address: order.Address, Date: order.Date, RestaurantName: order.RestaurantName, RestaurantSlug: order.RestaurantSlug, TotalPrice: order.TotalPrice, Status: order.Status, Cart: make([]models.OrderPositionUcaseResp, len(order.Cart))}
	for i, poz := range order.Cart {
		resp.Cart[i] = models.OrderPositionUcaseResp{Name: poz.Name, Description: poz.Description, ImagePath: poz.ImagePath, Calories: poz.Calories, Count: poz.Count, Price: poz.Price, Weigth: poz.Weigth}
	}
	return &resp, err
}
