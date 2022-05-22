package grpc

import (
	"context"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/order"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/order/models"
	proto "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/order/proto"
)

type grpcOrderHandler struct {
	Ucase order.Ucase
	proto.UnimplementedOrderServiceServer
}

func NewOrderHandler(ucase order.Ucase) *grpcOrderHandler {
	return &grpcOrderHandler{
		Ucase: ucase,
	}
}

func (h grpcOrderHandler) CreateOrder(ctx context.Context, req *proto.CreateOrderReq) (*proto.CreateOrderResp, error) {
	cart := make([]models.OrderPositionUcase, len(req.Cart))
	for i, position := range req.Cart {
		cart[i] = models.OrderPositionUcase{Id: position.Id, Count: position.Count}
	}
	orderResp, err := h.Ucase.CreateOrder(&models.CreateOrderUcaseReq{Address: req.Address, Comment: req.Comment, UserId: req.UserId, Cart: cart})
	if err != nil {
		cause := servErrors.ErrorAs(err)
		if cause == nil {
			return &proto.CreateOrderResp{}, status.Error(codes.Internal, err.Error())
		}
		return &proto.CreateOrderResp{}, status.Error(codes.Code(cause.Code), err.Error())
	}
	return &proto.CreateOrderResp{OrderId: orderResp.OrderId}, nil
}

func (h grpcOrderHandler) GetUserOrders(ctx context.Context, req *proto.GetUserOrdersReq) (*proto.GetUserOrdersResp, error) {
	ordersUcaseResp, err := h.Ucase.GetUserOrders(&models.GetUserOrdersUcaseReq{UserId: req.UserId})
	if err != nil {
		cause := servErrors.ErrorAs(err)
		if cause == nil {
			return &proto.GetUserOrdersResp{}, status.Error(codes.Internal, err.Error())
		}
		return &proto.GetUserOrdersResp{}, status.Error(codes.Code(cause.Code), err.Error())
	}
	ordersResp := make([]*proto.ShortOrderResp, len(ordersUcaseResp.Orders))
	for i, position := range ordersUcaseResp.Orders {
		ordersResp[i] = &proto.ShortOrderResp{OrderId: position.OrderId, Date: position.Date, RestaurantName: position.RestaurantName, Status: position.Status, TotalPrice: position.TotalPrice}
	}
	return &proto.GetUserOrdersResp{Orders: ordersResp}, nil
}

func (h grpcOrderHandler) GetUserOrderStatuses(ctx context.Context, req *proto.GetUserOrderStatusesReq) (*proto.GetUserOrderStatusesResp, error) {
	ordersUcaseResp, err := h.Ucase.GetUserOrders(&models.GetUserOrdersUcaseReq{UserId: req.UserId})
	if err != nil {
		cause := servErrors.ErrorAs(err)
		if cause == nil {
			return &proto.GetUserOrderStatusesResp{}, status.Error(codes.Internal, err.Error())
		}
		return &proto.GetUserOrderStatusesResp{}, status.Error(codes.Code(cause.Code), err.Error())
	}
	ordersStatusesResp := make([]*proto.OrderStatus, len(ordersUcaseResp.Orders))
	for i, position := range ordersUcaseResp.Orders {
		ordersStatusesResp[i] = &proto.OrderStatus{OrderId: position.OrderId, Status: position.Status}
	}
	return &proto.GetUserOrderStatusesResp{OrderStatuses: ordersStatusesResp}, nil
}

func (h grpcOrderHandler) GetUserOrder(ctx context.Context, req *proto.GetUserOrderReq) (*proto.GetUserOrderResp, error) {
	order, err := h.Ucase.GetUserOrder(&models.GetUserOrderUcaseReq{UserId: req.UserId, OrderId: req.OrderId})
	if err != nil {
		cause := servErrors.ErrorAs(err)
		if cause == nil {
			return &proto.GetUserOrderResp{}, status.Error(codes.Internal, err.Error())
		}
		return &proto.GetUserOrderResp{}, status.Error(codes.Code(cause.Code), err.Error())
	}

	cart := make([]*proto.OrderPositionResp, len(order.Cart))
	for i, poz := range order.Cart {
		cart[i] = &proto.OrderPositionResp{Name: poz.Name, Description: poz.Description, ImagePath: poz.ImagePath, Calories: poz.Calories, Count: poz.Count, Price: poz.Price, Weigth: poz.Weigth}
	}

	return &proto.GetUserOrderResp{OrderId: order.OrderId, Address: order.Address, Date: order.Date, RestaurantName: order.RestaurantName, RestaurantSlug: order.RestaurantSlug, TotalPrice: order.TotalPrice, Status: order.Status, Cart: cart}, nil
}