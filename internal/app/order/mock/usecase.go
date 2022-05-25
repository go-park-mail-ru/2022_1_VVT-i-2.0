package mock

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/stretchr/testify/mock"
)

type OrderUcase struct {
	mock.Mock
}

func (u *OrderUcase) CreateOrder(order *models.OrderUcaseReq) (*models.OrderUcaseResp, error) {
	return &models.OrderUcaseResp{OrderId: 1}, nil
}

func (u *OrderUcase) GetUserOrders(order *models.GetUserOrdersUcaseReq) (*models.GetUserOrdersUcaseResp, error) {
	return &models.GetUserOrdersUcaseResp{Orders: []models.ShortOrderUcase{
		{OrderId: 1, RestaurantName: "RestName", TotalPrice: 100, Date: "01.01.2021", Status: "Получен"},
	}}, nil
}

func (u *OrderUcase) GetUserOrderStatuses(order *models.GetUserOrderStatusesUcaseReq) (*models.GetUserOrderStatusesUcaseResp, error) {
	return &models.GetUserOrderStatusesUcaseResp{OrderStatuses: []models.OrderStatusUcase{
		{OrderId: 1, Status: "Получен"},
	}}, nil
}

func (u *OrderUcase) GetUserOrder(req *models.GetUserOrderUcaseReq) (*models.GetUserOrderUcaseResp, error) {
	return &models.GetUserOrderUcaseResp{
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
	}, nil
}

type OrderUcaseErr struct {
	mock.Mock
}

func (u *OrderUcaseErr) CreateOrder(order *models.OrderUcaseReq) (*models.OrderUcaseResp, error) {
	return nil, servErrors.NewError(servErrors.DB_ERROR, "")
}

func (u *OrderUcaseErr) GetUserOrders(order *models.GetUserOrdersUcaseReq) (*models.GetUserOrdersUcaseResp, error) {
	return nil, servErrors.NewError(servErrors.DB_ERROR, "")
}

func (u *OrderUcaseErr) GetUserOrderStatuses(order *models.GetUserOrderStatusesUcaseReq) (*models.GetUserOrderStatusesUcaseResp, error) {
	return nil, servErrors.NewError(servErrors.DB_ERROR, "")
}

func (u *OrderUcaseErr) GetUserOrder(req *models.GetUserOrderUcaseReq) (*models.GetUserOrderUcaseResp, error) {
	return nil, servErrors.NewError(servErrors.DB_ERROR, "")
}
