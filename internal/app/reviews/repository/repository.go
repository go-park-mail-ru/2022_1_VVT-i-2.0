package repository

import (
	"database/sql"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/jmoiron/sqlx"
)

type ReviewsRepo struct {
	DB *sqlx.DB
}

func NewReviewsRepo(db *sqlx.DB) *ReviewsRepo {
	return &ReviewsRepo{DB: db}
}

func (r *ReviewsRepo) GetRestaurantReviews(req *models.GetRestaurantReviewsRepoReq) (*models.GetRestaurantReviewsRepoResp, error) {
	reviews := make([]*models.RestaurantReviewRepo, 0)
	err := r.DB.Select(&reviews, `SELECT author, text, stars, get_ru_date(date) FROM reviews c JOIN restaurants r ON c.restaurant_id=r.id WHERE r.slug = $1 ORDER BY date DESC`, req.Slug)

	switch err {
	case nil:
		resp := &models.GetRestaurantReviewsRepoResp{Reviews: make([]models.RestaurantReviewRepo, len(reviews))}
		for i, review := range reviews {
			resp.Reviews[i] = *review
		}
		return resp, nil
	default:
		return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
	}
}

func (r *ReviewsRepo) AddRestaurantReview(req *models.AddRestaurantReviewRepoReq) (*models.RestaurantReviewRepo, error) {
	comment := models.RestaurantReviewRepo{}
	err := r.DB.Get(&comment, `INSERT INTO reviews (restaurant_id, author, text, stars, order_id) VALUES ((SELECT id FROM restaurants WHERE slug=$1),(SELECT name FROM users WHERE id=$2),$3,$4,$5) RETURNING author, text, stars, get_ru_date(date)`, req.Slug, req.UserId, req.Text, req.Rating, req.OrderId)
	if err != nil {
		if err == sql.ErrConnDone || err == sql.ErrTxDone {
			return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
		}
		return nil, servErrors.NewError(servErrors.DB_INSERT, err.Error())
	}
	return &comment, nil
}

func (r *ReviewsRepo) HasReviewToOrder(req *models.CanReviewedRepoReq) (*models.CanReviewedRepoResp, error) {
	canReviewed := models.CanReviewedRepoResp{}
	err := r.DB.Get(&canReviewed, `SELECT ((NOT reviewed) AND user_id=$1) can_reviewed FROM orders_internal WHERE id=$2`, req.UserId, req.OrderId)
	if err != nil {
		if err == sql.ErrConnDone || err == sql.ErrTxDone {
			return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
		}
		return &models.CanReviewedRepoResp{Can: false}, nil
	}
	return &canReviewed, nil
}
