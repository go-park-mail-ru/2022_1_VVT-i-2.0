package repository

import (
	"database/sql"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/jmoiron/sqlx"
)

type CommentsRepo struct {
	DB *sqlx.DB
}

func NewRestaurantsRepo(db *sqlx.DB) *CommentsRepo {
	return &CommentsRepo{DB: db}
}

func (r *CommentsRepo) GetRestaurantByID(id int) (*models.RestaurantDataStorage, error) {
	restaurant := &models.RestaurantDataStorage{}
	err := r.DB.Get(restaurant, "SELECT id, name, image_path, slug, min_price,  rating, count_rating FROM restaurants WHERE id = $1", id)
	switch err {
	case nil:
		return restaurant, nil
	case sql.ErrNoRows:
		return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, err.Error())
	default:
		return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
	}
}

func (r *CommentsRepo) GetRestaurantComments(id int) ([]*models.CommentRestaurantDataStorage, error) {
	comments := make([]*models.CommentRestaurantDataStorage, 0, 3)
	err := r.DB.Select(&comments, `SELECT id, restaurant, user_id, comment_text, comment_rating FROM comment_restaurants WHERE restaurant = $1`, id)

	switch err {
	case nil:
		return comments, nil
	case sql.ErrNoRows:
		return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, err.Error())
	default:
		return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
	}
}

func (r *CommentsRepo) AddRestaurantComment(newComment *models.AddCommentRestaurantDataStorage) (*models.CommentRestaurantDataStorage, error) {
	comment := &models.CommentRestaurantDataStorage{}
	err := r.DB.Get(comment, `INSERT INTO comment_restaurants (restaurant, user_id, comment_text, comment_rating) VALUES ($1,$2,$3,$4) RETURNING id, restaurant, user_id, comment_text, comment_rating`, newComment.Restaurant, newComment.User_id, newComment.Comment_text, newComment.Comment_rating)
	if err != nil {
		if err == sql.ErrConnDone || err == sql.ErrTxDone {
			return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
		}
		return nil, servErrors.NewError(servErrors.DB_INSERT, err.Error())
	}
	if comment == nil {
		return nil, servErrors.NewError(servErrors.DB_INSERT, "")
	}
	return comment, nil
}

func (r *CommentsRepo) UpdateRestaurantRating(restId int, newRestRating int, countRating int) (*models.RestaurantDataStorage, error) {
	restaurant := &models.RestaurantDataStorage{}
	err := r.DB.Get(restaurant, `UPDATE restaurants SET rating=$1, count_rating=$2 WHERE id=$3 RETURNING id, name, city, address, image_path, slug, min_price, avg_price, rating, count_rating`, newRestRating, countRating, restId)
	if err != nil {
		if err == sql.ErrConnDone || err == sql.ErrTxDone {
			return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
		}
		return nil, servErrors.NewError(servErrors.DB_INSERT, err.Error())
	}
	if restaurant == nil {
		return nil, servErrors.NewError(servErrors.DB_INSERT, "")
	}
	return restaurant, nil
}
