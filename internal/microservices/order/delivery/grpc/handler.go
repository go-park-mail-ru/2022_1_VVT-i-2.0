package grpc

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/order"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/order/models"
	proto "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/order/proto"
)

type grpcOrderHandler struct {
	Usecase order.Usecase
	proto.UnimplementedOrderServiceServer
}

func NewOrderHandler(usecase order.Usecase) *grpcOrderHandler {
	return &grpcOrderHandler{
		Usecase: usecase,
	}
}

func (h grpcOrderHandler) CreateOrder(ctx context.Context, req *proto.CreateOrderReq) (*proto.CreateOrderResp, error) {
	fmt.Println("grpc-h")
	cart := make([]models.OrderPositionUcase, len(req.Cart))
	for i, position := range req.Cart {
		cart[i] = models.OrderPositionUcase{Id: position.Id, Count: position.Count}
	}
	orderResp, err := h.Usecase.CreateOrder(&models.CreateOrderUcaseReq{Address: req.Address, Comment: req.Comment, UserId: req.UserId, Cart: cart})
	fmt.Println("grpc-eh")
	fmt.Println(orderResp, err)
	if err != nil {
		cause := servErrors.ErrorAs(err)
		if cause == nil {
			return &proto.CreateOrderResp{}, status.Error(codes.Internal, err.Error())
		}
		return &proto.CreateOrderResp{}, status.Error(codes.Code(cause.Code), err.Error())
	}
	return &proto.CreateOrderResp{OrderId: orderResp.OrderId}, nil
}

// GetUserOrder(context.Context, *GetUserOrderReq) (*GetUserOrderResp, error)
// GetUserOrderStatuses(context.Context, *GetUserOrderStatusesReq) (*GetUserOrderStatusesResp, error)

func (h grpcOrderHandler) GetUserOrders(ctx context.Context, req *proto.GetUserOrdersReq) (*proto.GetUserOrdersResp, error) {
	fmt.Println("grpc-h")
	ordersUcaseResp, err := h.Usecase.GetUserOrders(&models.GetUserOrdersUcaseReq{UserId: req.UserId})
	fmt.Println("grpc-eh")
	fmt.Println(ordersUcaseResp, err)
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
	fmt.Println("grpc-h")
	ordersUcaseResp, err := h.Usecase.GetUserOrders(&models.GetUserOrdersUcaseReq{UserId: req.UserId})
	fmt.Println("grpc-eh")
	fmt.Println(ordersUcaseResp, err)
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
	fmt.Println("------------------------------")
	fmt.Println(ordersStatusesResp)
	return &proto.GetUserOrderStatusesResp{OrderStatuses: ordersStatusesResp}, nil
}

func (h grpcOrderHandler) GetUserOrder(ctx context.Context, req *proto.GetUserOrderReq) (*proto.GetUserOrderResp, error) {
	fmt.Println("grpc-h")
	order, err := h.Usecase.GetUserOrder(&models.GetUserOrderUcaseReq{UserId: req.UserId, OrderId: req.OrderId})
	fmt.Println("grpc-eh")
	fmt.Println(order, err)
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

	return &proto.GetUserOrderResp{OrderId: order.OrderId, Address: order.Address, Date: order.Date, RestaurantName: order.RestaurantName, TotalPrice: order.TotalPrice, Status: order.Status, Cart: cart}, nil
}
