package reviews

import "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"

type Repository interface {
	GetRestaurantReviews(req *models.GetRestaurantReviewsRepoReq) (*models.GetRestaurantReviewsRepoResp, error)
	AddRestaurantReview(req *models.AddRestaurantReviewRepoReq) (*models.RestaurantReviewRepo, error)
	HasReviewToOrder(req *models.CanReviewedRepoReq) (*models.CanReviewedRepoResp, error)
}
