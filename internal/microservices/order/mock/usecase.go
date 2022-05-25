package mock

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/order/models"
	"github.com/stretchr/testify/mock"
)

type OrderUcase struct {
	mock.Mock
}

func (r *OrderUcase) CreateOrder(order *models.CreateOrderUcaseReq) (*models.CreateOrderUcaseResp, error) {
	return &models.CreateOrderUcaseResp{OrderId: 1}, nil
}

func (r *OrderUcase) GetUserOrders(user *models.GetUserOrdersUcaseReq) (*models.GetUserOrdersUcaseResp, error) {
	return &models.GetUserOrdersUcaseResp{Orders: []models.ShortOrderUcase{{OrderId: 1, RestaurantName: "RestName", TotalPrice: 100, Date: "01.01.2021", Status: "Получен"}}}, nil
}

func (r *OrderUcase) GetUserOrder(user *models.GetUserOrderUcaseReq) (*models.GetUserOrderUcaseResp, error) {
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

func (r *OrderUcase) GetUserOrderStatuses(user *models.GetUserOrderStatusesUcaseReq) (*models.GetUserOrderStatusesUcaseResp, error) {
	return &models.GetUserOrderStatusesUcaseResp{OrderStatuses: []models.OrderStatusUcase{{OrderId: 1, Status: "Получен"}}}, nil
}

type OrderUcaseErr struct {
	mock.Mock
}

func (r *OrderUcaseErr) CreateOrder(order *models.CreateOrderUcaseReq) (*models.CreateOrderUcaseResp, error) {
	return nil, servErrors.NewError(servErrors.DB_ERROR, "")
}

func (r *OrderUcaseErr) GetUserOrders(user *models.GetUserOrdersUcaseReq) (*models.GetUserOrdersUcaseResp, error) {
	return nil, servErrors.NewError(servErrors.DB_ERROR, "")
}

func (r *OrderUcaseErr) GetUserOrder(user *models.GetUserOrderUcaseReq) (*models.GetUserOrderUcaseResp, error) {
	return nil, servErrors.NewError(servErrors.DB_ERROR, "")
}

func (r *OrderUcaseErr) GetUserOrderStatuses(user *models.GetUserOrderStatusesUcaseReq) (*models.GetUserOrderStatusesUcaseResp, error) {
	return nil, servErrors.NewError(servErrors.DB_ERROR, "")
}
