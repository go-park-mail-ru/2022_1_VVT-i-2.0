package models

//easyjson:json
type ReviewResp struct {
	Author string `json:"author"`
	Text   string `json:"text"`
	Stars  int    `json:"stars"`
	Date   string `json:"date"`
}

//easyjson:json
type GetRestaurantReviews struct {
	Reviews []ReviewResp `json:"comment"`
}

type GetRestaurantReviewsUcaseReq struct {
	Slug string
}

type RestaurantReviewUcase struct {
	Author string
	Text   string
	Stars  int
	Date   string
}

type GetRestaurantReviewsUcaseResp struct {
	Reviews []RestaurantReviewUcase
}

type GetRestaurantReviewsRepoReq struct {
	Slug string
}

type RestaurantReviewRepo struct {
	Author string `db:"author"`
	Text   string `db:"text"`
	Stars  int    `db:"stars"`
	Date   string `db:"get_ru_date"`
}

type GetRestaurantReviewsRepoResp struct {
	Reviews []RestaurantReviewRepo
}

//easyjson:json
type AddRestaurantReviewReq struct {
	Slug    string `json:"slug"`
	Text    string `json:"text"`
	Rating  int    `json:"stars"`
	OrderId int64  `json:"orderId"`
}

//easyjson:json
type AddRestaurantReviewResp struct {
	Author string `json:"author"`
	Text   string `json:"text"`
	Stars  int    `json:"stars"`
	Date   string `json:"date"`
}

type AddRestaurantReviewUcaseReq struct {
	UserId  int64
	Slug    string
	Text    string
	Rating  int
	OrderId int64
}

type AddRestaurantReviewUcaseResp struct {
	Author string
	Text   string
	Stars  int
	Date   string
}

type AddRestaurantReviewRepoReq struct {
	UserId  int64
	Slug    string
	Text    string
	Rating  int
	OrderId int64
}

type CanReviewedRepoReq struct {
	OrderId int64
	UserId  int64
}

type CanReviewedRepoResp struct {
	Can bool `db:"can_reviewed"`
}
