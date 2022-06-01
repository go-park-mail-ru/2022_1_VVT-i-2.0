package mock

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/stretchr/testify/mock"
)

type CommentsUsecase struct {
	mock.Mock
}

func (a *CommentsUsecase) GetRestaurantReviews(req *models.GetRestaurantReviewsUcaseReq) (*models.GetRestaurantReviewsUcaseResp, error) {
	if req.Slug == "" {
		return nil, nil
	}
	return &models.GetRestaurantReviewsUcaseResp{Reviews: []models.RestaurantReviewUcase{{
		Author: "author",
		Text:   "text",
		Stars:  4,
		Date:   "date"}}}, nil
}

func (a *CommentsUsecase) AddRestaurantReview(req *models.AddRestaurantReviewUcaseReq) (*models.AddRestaurantReviewUcaseResp, error) {
	return &models.AddRestaurantReviewUcaseResp{
		Author: "author",
		Text:   "text",
		Stars:  4,
		Date:   "date"}, nil
}

type CommentsUsecaseErr struct {
	mock.Mock
}

func (a *CommentsUsecaseErr) GetRestaurantReviews(req *models.GetRestaurantReviewsUcaseReq) (*models.GetRestaurantReviewsUcaseResp, error) {
	return nil, servErrors.NewError(servErrors.DB_ERROR, "")
}

func (a *CommentsUsecaseErr) AddRestaurantReview(req *models.AddRestaurantReviewUcaseReq) (*models.AddRestaurantReviewUcaseResp, error) {
	return nil, servErrors.NewError(servErrors.DB_ERROR, "")
}
