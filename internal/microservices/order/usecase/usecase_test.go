package ucase

import (
	"reflect"
	"testing"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/order/mock"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/order/models"
	"github.com/stretchr/testify/require"
)

func TestOrderUcase_CreateOrder(t *testing.T) {
	mockUserRepo := new(mock.OrderRepo)
	ucase := NewOrderUcase(mockUserRepo)

	req := &models.CreateOrderUcaseReq{
		Address: "Москва, Измайловский проспект, 73/2",
		Comment: "comment",
		Cart:    []models.OrderPositionUcase{{Id: 1, Count: 2}},
	}

	userResp, err := ucase.CreateOrder(req)
	require.NoError(t, err)

	expectResp := &models.CreateOrderUcaseResp{
		OrderId: 1,
	}

	if !reflect.DeepEqual(userResp, expectResp) {
		t.Errorf("results not match, want %v, have %v", userResp, expectResp)
		return
	}
}

func TestOrderUcase_GetUserOrders(t *testing.T) {
	mockUserRepo := new(mock.OrderRepo)
	ucase := NewOrderUcase(mockUserRepo)

	req := &models.GetUserOrdersUcaseReq{
		UserId: 1,
	}

	userResp, err := ucase.GetUserOrders(req)
	require.NoError(t, err)
	require.NotNil(t, userResp)

}

func TestOrderUcase_GetUserOrderStatuses(t *testing.T) {
	mockUserRepo := new(mock.OrderRepo)
	ucase := NewOrderUcase(mockUserRepo)

	req := &models.GetUserOrderStatusesUcaseReq{
		UserId: 1,
	}

	userResp, err := ucase.GetUserOrderStatuses(req)
	require.NoError(t, err)
	require.NotNil(t, userResp)

}

func TestOrderUcase_GetUserOrder_WhongUser(t *testing.T) {
	mockUserRepo := new(mock.OrderRepo)
	ucase := NewOrderUcase(mockUserRepo)

	req := &models.GetUserOrderUcaseReq{
		UserId:  2,
		OrderId: 1,
	}

	userResp, err := ucase.GetUserOrder(req)
	require.Error(t, err)
	require.Nil(t, userResp)
}

func TestOrderUcase_GetUserOrder(t *testing.T) {
	mockUserRepo := new(mock.OrderRepo)
	ucase := NewOrderUcase(mockUserRepo)

	req := &models.GetUserOrderUcaseReq{
		UserId:  1,
		OrderId: 1,
	}

	userResp, err := ucase.GetUserOrder(req)
	require.NoError(t, err)
	require.NotNil(t, userResp)
}
