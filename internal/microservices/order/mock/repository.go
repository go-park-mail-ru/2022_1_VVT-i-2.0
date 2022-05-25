package mock

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/microservices/order/models"
	"github.com/stretchr/testify/mock"
)

type OrderRepo struct {
	mock.Mock
}

func (r *OrderRepo) CreateOrder(order *models.CreateOrderRepoReq) (*models.CreateOrderRepoResp, error) {
	return &models.CreateOrderRepoResp{OrderId: 1}, nil
}

func (r *OrderRepo) GetUserOrders(user *models.GetUserOrdersRepoReq) (*models.GetUserOrdersRepoResp, error) {
	return &models.GetUserOrdersRepoResp{Orders: []models.ShortOrderRepo{{OrderId: 1, RestaurantName: "RestName", TotalPrice: 100, Date: "01.01.2021", Status: "Получен"}}}, nil
}

func (r *OrderRepo) GetUserOrder(user *models.GetUserOrderRepoReq) (*models.GetUserOrderRepoResp, error) {
	return &models.GetUserOrderRepoResp{
		UserId:         1,
		OrderId:        1,
		RestaurantName: "RestName",
		TotalPrice:     100,
		Date:           "01.01.2021",
		Status:         "Получен",
		Cart: []models.OrderPositionRepoResp{
			{
				Name:        "name",
				ImagePath:   "img.png",
				Price:       10,
				Count:       10,
				Calories:    100,
				Weight:      50,
				Description: "description"},
		},
	}, nil
}

func (r *OrderRepo) GetUserOrderStatuses(user *models.GetUserOrderStatusesRepoReq) (*models.GetUserOrderStatusesRepoResp, error) {
	return &models.GetUserOrderStatusesRepoResp{OrderStatuses: []models.OrderStatusRepo{{OrderId: 1, Status: "Получен"}}}, nil
}

func (r *OrderRepo) GetAddress(req *models.GetAddressRepoReq) (*models.GetAddressRepoResp, error) {
	return &models.GetAddressRepoResp{City: "Москва", Street: "Петровка", House: "38"}, nil
}
