package reviews

import "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"

type Usecase interface {
	GetRestaurantReviews(req *models.GetRestaurantReviewsUcaseReq) (*models.GetRestaurantReviewsUcaseResp, error)
	AddRestaurantReview(req *models.AddRestaurantReviewUcaseReq) (*models.AddRestaurantReviewUcaseResp, error)
}
