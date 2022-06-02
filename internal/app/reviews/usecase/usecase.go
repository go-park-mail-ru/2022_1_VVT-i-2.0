package usecase

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/reviews"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/pkg/errors"
)

type RestaurantReviewsUcase struct {
	Repo reviews.Repository
}

func NewRestaurantReviewsUcase(repo reviews.Repository) *RestaurantReviewsUcase {
	return &RestaurantReviewsUcase{
		Repo: repo,
	}
}

func (u *RestaurantReviewsUcase) GetRestaurantReviews(req *models.GetRestaurantReviewsUcaseReq) (*models.GetRestaurantReviewsUcaseResp, error) {
	reviewsRepoResp, err := u.Repo.GetRestaurantReviews(&models.GetRestaurantReviewsRepoReq{Slug: req.Slug})
	if err != nil {
		return nil, errors.Wrap(err, "error getting comments from storage")
	}

	reviewsResp := &models.GetRestaurantReviewsUcaseResp{Reviews: make([]models.RestaurantReviewUcase, len(reviewsRepoResp.Reviews))}

	for i, review := range reviewsRepoResp.Reviews {
		reviewsResp.Reviews[i] = models.RestaurantReviewUcase(review)
	}
	return reviewsResp, nil
}

func (u *RestaurantReviewsUcase) AddRestaurantReview(req *models.AddRestaurantReviewUcaseReq) (*models.AddRestaurantReviewUcaseResp, error) {
	canReviewed, err := u.Repo.HasReviewToOrder(&models.CanReviewedRepoReq{UserId: req.UserId, OrderId: req.OrderId})
	if err != nil {
		return nil, errors.Wrap(err, "error chacking can such user review order")
	}
	if !canReviewed.Can {
		return nil, servErrors.NewError(servErrors.ORDER_ALREADY_REVIEWED, "")
	}
	comment, err := u.Repo.AddRestaurantReview(&models.AddRestaurantReviewRepoReq{UserId: int64(req.UserId), Slug: req.Slug, Rating: req.Rating, Text: req.Text, OrderId: req.OrderId})

	if err != nil {
		return nil, errors.Wrap(err, "error adding review")
	}
	return (*models.AddRestaurantReviewUcaseResp)(comment), nil
}
