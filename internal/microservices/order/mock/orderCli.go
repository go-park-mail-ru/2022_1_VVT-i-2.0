package mock

import (
	"context"

	proto "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/order/proto"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type OrderGrpcCli struct {
	mock.Mock
}

func (c *OrderGrpcCli) CreateOrder(ctx context.Context, in *proto.CreateOrderReq, opts ...grpc.CallOption) (*proto.CreateOrderResp, error) {
	return &proto.CreateOrderResp{OrderId: 1}, nil
}

func (c *OrderGrpcCli) GetUserOrders(ctx context.Context, in *proto.GetUserOrdersReq, opts ...grpc.CallOption) (*proto.GetUserOrdersResp, error) {
	return &proto.GetUserOrdersResp{Orders: []*proto.ShortOrderResp{{OrderId: 1, RestaurantName: "RestName", TotalPrice: 100, Date: "01.01.2021", Status: "Получен"}}}, nil
}

func (c *OrderGrpcCli) GetUserOrder(ctx context.Context, in *proto.GetUserOrderReq, opts ...grpc.CallOption) (*proto.GetUserOrderResp, error) {
	return &proto.GetUserOrderResp{
		OrderId:        1,
		RestaurantName: "RestName",
		TotalPrice:     100,
		Date:           "01.01.2021",
		Status:         "Получен",
		Cart: []*proto.OrderPositionResp{
			{
				Name:        "name",
				ImagePath:   "img.png",
				Price:       10,
				Count:       10,
				Calories:    100,
				Weigth:      50,
				Description: "description"},
		},
	}, nil
}

func (c *OrderGrpcCli) GetUserOrderStatuses(ctx context.Context, in *proto.GetUserOrderStatusesReq, opts ...grpc.CallOption) (*proto.GetUserOrderStatusesResp, error) {
	return &proto.GetUserOrderStatusesResp{OrderStatuses: []*proto.OrderStatus{{OrderId: 1, Status: "Получен"}}}, nil
}
