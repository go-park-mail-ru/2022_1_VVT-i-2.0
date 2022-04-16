package usecase

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/order"
	"github.com/pkg/errors"
)

type OrderUsecase struct {
	OrderRepo order.Repository
}

func NewUsecase(orderRepo order.Repository) *OrderUsecase {
	return &OrderUsecase{
		OrderRepo: orderRepo,
	}
}

func (u *OrderUsecase) Order(order *models.OrderUcaseInput) (*models.OrderUcaseAnsw, error) {
	// TODO: сделать проверку, есть ли такой адрес
	orderId, err := u.OrderRepo.AddOrder(&models.OrderRepoInput{User_id: order.UserId,
		Address: order.Address,
		Comment: order.Comment,
		Cart:    order.Cart})
	if err != nil || orderId.OrderId <= 0 {
		return nil, errors.Wrap(err, "error adding order to storage")
	}
	return &models.OrderUcaseAnsw{OrderId: orderId.OrderId}, nil

}
