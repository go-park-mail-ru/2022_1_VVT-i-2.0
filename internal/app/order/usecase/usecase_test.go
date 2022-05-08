package ucase

import (
	"reflect"
	"testing"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	mockOrderCli "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/order/mock"
	"github.com/stretchr/testify/assert"
)

func TestOrderUcase_GetOrder(t *testing.T) {
	mockOrderCli := new(mockOrderCli.OrderGrpcCli)
	ucase := NewUcase(mockOrderCli)

	resp, err := ucase.GetUserOrder(&models.GetUserOrderUcaseReq{OrderId: 1})

	assert.NoError(t, err)

	expectResp := &models.GetUserOrderUcaseResp{
		OrderId:        1,
		RestaurantName: "RestName",
		TotalPrice:     100,
		Date:           "01.01.2021",
		Status:         "Получен",
		Cart: []models.OrderPositionUcaseResp{
			{
				Name:        "name",
				ImagePath:   "img.png",
				Price:       10,
				Count:       10,
				Calories:    100,
				Weigth:      50,
				Description: "description"},
		},
	}

	if !reflect.DeepEqual(resp, expectResp) {
		t.Errorf("results not match, want %v, have %v", resp, expectResp)
		return
	}
}

func TestOrderUcase_CreateOrder(t *testing.T) {
	mockOrderCli := new(mockOrderCli.OrderGrpcCli)
	ucase := NewUcase(mockOrderCli)

	resp, err := ucase.CreateOrder(&models.OrderUcaseReq{
		Address: "Москва, Измайловский проспект, 73/2",
		Comment: "comment",
		Cart:    []models.OrderPosition{{Id: 1, Count: 2}}})

	assert.NoError(t, err)

	expectResp := &models.OrderUcaseResp{OrderId: 1}

	if !reflect.DeepEqual(resp, expectResp) {
		t.Errorf("results not match, want %v, have %v", resp, expectResp)
		return
	}
}

func TestOrderUcase_GetOrders(t *testing.T) {
	mockOrderCli := new(mockOrderCli.OrderGrpcCli)
	ucase := NewUcase(mockOrderCli)

	resp, err := ucase.GetUserOrders(&models.GetUserOrdersUcaseReq{UserId: 1})

	assert.NoError(t, err)

	expectResp := &models.GetUserOrdersUcaseResp{Orders: []models.ShortOrderUcase{{OrderId: 1, RestaurantName: "RestName", TotalPrice: 100, Date: "01.01.2021", Status: "Получен"}}}

	if !reflect.DeepEqual(resp, expectResp) {
		t.Errorf("results not match, want %v, have %v", resp, expectResp)
		return
	}
}

func TestOrderUcase_GetOrderStatuses(t *testing.T) {
	mockOrderCli := new(mockOrderCli.OrderGrpcCli)
	ucase := NewUcase(mockOrderCli)

	resp, err := ucase.GetUserOrderStatuses(&models.GetUserOrderStatusesUcaseReq{UserId: 1})

	assert.NoError(t, err)

	expectResp := &models.GetUserOrderStatusesUcaseResp{OrderStatuses: []models.OrderStatusUcase{{OrderId: 1, Status: "Получен"}}}

	if !reflect.DeepEqual(resp, expectResp) {
		t.Errorf("results not match, want %v, have %v", resp, expectResp)
		return
	}
}
