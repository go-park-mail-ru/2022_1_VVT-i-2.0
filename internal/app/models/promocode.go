package models

type PromocodeResp struct {
	Image          string `json:"img"`
	LogoImage      string `json:"logo"`
	Text           string `json:"text"`
	Promocode      string `json:"promocode"`
	RestaurantName string `json:"restName"`
	RestaurantSlug string `json:"restSlug"`
	Discount       int    `json:"discount"`
	PriceReduction int    `json:"priceReduction"`
	MinPrice       int    `json:"minPrice"`
}
type GetPromocodesResp struct {
	Promos []PromocodeResp `json:"promos"`
}

type PromocodeUcaseResp struct {
	Image          string
	LogoImage      string
	Text           string
	Promocode      string
	RestaurantName string
	RestaurantSlug string
	Discount       int
	PriceReduction int
	MinPrice       int
}

type GetPromocodesUcaseResp struct {
	Promos []PromocodeUcaseResp
}

type PromocodeRepoResp struct {
	Image          string `db:"image_path"`
	LogoImage      string `db:"logo_path"`
	Text           string `db:"text"`
	Promocode      string `db:"promocode"`
	RestaurantName string `db:"name"`
	RestaurantSlug string `db:"slug"`
	Discount       int    `db:"discount"`
	PriceReduction int    `db:"price_reduction"`
	MinPrice       int    `db:"min_price"`
}

type GetPromocodesRepoResp struct {
	Promos []PromocodeRepoResp
}
