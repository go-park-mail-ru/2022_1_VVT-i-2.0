package mock

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/stretchr/testify/mock"
)

type PromoUcase struct {
	mock.Mock
}

func (u *PromoUcase) GetAllPromocodes() (*models.GetPromocodesUcaseResp, error) {
	return &models.GetPromocodesUcaseResp{Promos: []models.PromocodeUcaseResp{
		{Promocode: "promo1", RestaurantSlug: "slug1", RestaurantName: "name1", MinPrice: 101, Discount: 1, PriceReduction: 2, Text: "text1", LogoImage: "logo1", Image: "image1"},
		{Promocode: "promo2", RestaurantSlug: "slug2", RestaurantName: "name2", MinPrice: 102, Discount: 1, PriceReduction: 2, Text: "text2", LogoImage: "logo2", Image: "image2"},
	}}, nil
}
