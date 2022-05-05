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

func NewCommentsRepo(db *sqlx.DB) *CommentsRepo {
	return &CommentsRepo{DB: db}
}

func (r *CommentsRepo) GetRestaurantByID(id int) (*models.RestaurantDataStorage, error) {
	restaurant := &models.RestaurantDataStorage{}
	err := r.DB.Get(restaurant, "SELECT id, name, image_path, slug, min_price, up_time_to_delivery, down_time_to_delivery, review_count, agg_rating FROM restaurants WHERE id = $1", id)
	switch err {
	case nil:
		return restaurant, nil
	case sql.ErrNoRows:
		return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, err.Error())
	default:
		return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
	}
}

func (r *CommentsRepo) GetRestaurantBySlug(slug string) (*models.RestaurantDataStorage, error) {
	restaurant := &models.RestaurantDataStorage{}
	err := r.DB.Get(restaurant, "SELECT id, name,  image_path, slug, min_price, agg_rating, review_count, up_time_to_delivery, down_time_to_delivery FROM restaurants WHERE slug = $1", slug)
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
	comments := make([]*models.CommentRestaurantDataStorage, 0)
	err := r.DB.Select(&comments, `SELECT restaurant_id, author, text, stars, get_ru_date(date) date_creating FROM comments WHERE restaurant_id = $1 ORDER BY date DESC`, id)

	switch err {
	case nil:
		return comments, nil
	case sql.ErrNoRows:
		return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, err.Error())
	default:
		return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
	}
}

func (r *CommentsRepo) GetUserById(id models.UserId) (*models.UserDataRepo, error) {
	user := &models.UserDataRepo{}
	err := r.DB.Get(user, `SELECT id, phone, email, name, avatar FROM users WHERE id = $1`, id)

	switch err {
	case nil:
		return user, nil
	case sql.ErrNoRows:
		return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, err.Error())
	default:
		return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
	}
}

func (r *CommentsRepo) AddRestaurantComment(newComment *models.AddCommentRestaurantDataStorage) (*models.CommentRestaurantDataStorage, error) {
	comment := &models.CommentRestaurantDataStorage{}
	err := r.DB.Get(comment, `INSERT INTO comments (restaurant_id, author, text, stars) VALUES ($1,$2,$3,$4) RETURNING restaurant_id, author, text, stars, date_creating`, newComment.Restaurant_id, newComment.User, newComment.Comment_text, newComment.Comment_rating)
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
	err := r.DB.Get(restaurant, `UPDATE restaurants SET agg_rating=$1, review_count=$2 WHERE id=$3 RETURNING id, name, image_path, slug, image_path, slug, min_price, up_time_to_delivery, down_time_to_delivery, review_count, agg_rating`, newRestRating, countRating, restId)
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

//token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVkIjoiMjAyMi0wNS0wOFQxODozNDo0OS42ODYzMzgrMDM6MDAiLCJpZCI6NH0.XhuLHltOesExgjuC1bKFo5zl5D37rLXR2JNzTyqhtJw
// _csrf=ZhSKtRBGuTj24Pgc4tdtqNHL37NUHWux
