package mock

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/stretchr/testify/mock"
)

type CommentsRepository struct {
	mock.Mock
}

func (r *CommentsRepository) GetRestaurantReviews(req *models.GetRestaurantReviewsRepoReq) (*models.GetRestaurantReviewsRepoResp, error) {
	return &models.GetRestaurantReviewsRepoResp{Reviews: []models.RestaurantReviewRepo{{
		Author: "author",
		Text:   "text",
		Stars:  4,
		Date:   "date",
	}}}, nil
}

func (r *CommentsRepository) AddRestaurantReview(req *models.AddRestaurantReviewRepoReq) (*models.RestaurantReviewRepo, error) {
	return &models.RestaurantReviewRepo{
		Author: "author",
		Text:   "text",
		Stars:  4,
		Date:   "date"}, nil
}

type CommentsRepositoryErr struct {
	mock.Mock
}

func (c *CommentsRepositoryErr) GetRestaurantReviews(req *models.GetRestaurantReviewsRepoReq) (*models.GetRestaurantReviewsRepoResp, error) {
	return nil, servErrors.NewError(servErrors.DB_ERROR, "")
}

func (c *CommentsRepositoryErr) AddRestaurantReview(req *models.AddRestaurantReviewRepoReq) (*models.RestaurantReviewRepo, error) {
	return nil, servErrors.NewError(servErrors.DB_ERROR, "")
}
