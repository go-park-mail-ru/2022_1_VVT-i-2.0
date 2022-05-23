package grpc

import (
	"context"
	"reflect"
	"testing"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/order/mock"
	proto "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/order/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOrderHandler_CreateOrder(t *testing.T) {
	mockUcase := new(mock.OrderUcase)
	handler := NewOrderHandler(mockUcase)

	req := proto.CreateOrderReq{
		Address: "Москва, Измайловский проспект, 73/2",
		Comment: "comment",
		Cart:    []*proto.OrderPositionReq{{Id: 1, Count: 2}},
	}

	userResp, err := handler.CreateOrder(context.Background(), &req)
	require.NoError(t, err)

	expectResp := &proto.CreateOrderResp{
		OrderId: 1,
	}

	if !reflect.DeepEqual(userResp, expectResp) {
		t.Errorf("results not match, want %v, have %v", userResp, expectResp)
		return
	}
}

func TestOrderHandler_GetOrders(t *testing.T) {
	mockUcase := new(mock.OrderUcase)
	handler := NewOrderHandler(mockUcase)

	req := proto.GetUserOrdersReq{
		UserId: 1,
	}

	userResp, err := handler.GetUserOrders(context.Background(), &req)
	require.NoError(t, err)

	require.NotNil(t, userResp)

}

func TestOrderHandler_GetOrderStatuses(t *testing.T) {
	mockUcase := new(mock.OrderUcase)
	handler := NewOrderHandler(mockUcase)

	req := proto.GetUserOrderStatusesReq{
		UserId: 1,
	}

	userResp, err := handler.GetUserOrderStatuses(context.Background(), &req)
	require.NoError(t, err)

	require.NotNil(t, userResp)
}

func TestOrderHandler_GetOrder(t *testing.T) {
	mockUcase := new(mock.OrderUcase)
	handler := NewOrderHandler(mockUcase)

	req := proto.GetUserOrderReq{
		UserId:  1,
		OrderId: 1,
	}

	userResp, err := handler.GetUserOrder(context.Background(), &req)
	require.NoError(t, err)

	require.NotNil(t, userResp)
}

func TestOrderHandler_CreateOrder_Err(t *testing.T) {
	mockUcase := new(mock.OrderUcaseErr)
	handler := NewOrderHandler(mockUcase)

	req := proto.CreateOrderReq{
		Address: "Москва, Измайловский проспект, 73/2",
		Comment: "comment",
		Cart:    []*proto.OrderPositionReq{{Id: 1, Count: 2}},
	}

	_, err := handler.CreateOrder(context.Background(), &req)
	assert.Error(t, err)
}

func TestOrderHandler_GetOrders_Err(t *testing.T) {
	mockUcase := new(mock.OrderUcaseErr)
	handler := NewOrderHandler(mockUcase)

	req := proto.GetUserOrdersReq{
		UserId: 1,
	}

	_, err := handler.GetUserOrders(context.Background(), &req)
	assert.Error(t, err)
}

func TestOrderHandler_GetOrderStatuses_Err(t *testing.T) {
	mockUcase := new(mock.OrderUcaseErr)
	handler := NewOrderHandler(mockUcase)

	req := proto.GetUserOrderStatusesReq{
		UserId: 1,
	}

	_, err := handler.GetUserOrderStatuses(context.Background(), &req)
	assert.Error(t, err)
}

func TestOrderHandler_GetOrder_Err(t *testing.T) {
	mockUcase := new(mock.OrderUcaseErr)
	handler := NewOrderHandler(mockUcase)

	req := proto.GetUserOrderReq{
		UserId:  1,
		OrderId: 1,
	}

	_, err := handler.GetUserOrder(context.Background(), &req)
	assert.Error(t, err)
}
