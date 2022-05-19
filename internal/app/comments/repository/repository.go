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

func (r *CommentsRepo) GetRestaurantByID(req models.GetRestaurantByIdRepoReq) (*models.RestaurantRepo, error) {
	restaurant := &models.RestaurantRepo{}
	err := r.DB.Get(restaurant, "SELECT id, name, image_path, slug, min_price, up_time_to_delivery, down_time_to_delivery, review_count, agg_rating FROM restaurants WHERE id = $1", req.Id)
	switch err {
	case nil:
		return restaurant, nil
	case sql.ErrNoRows:
		return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, err.Error())
	default:
		return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
	}
}

func (r *CommentsRepo) GetRestaurantBySlug(req models.GetRestaurantBySlugRepoReq) (*models.RestaurantRepo, error) {
	restaurant := &models.RestaurantRepo{}
	err := r.DB.Get(restaurant, "SELECT id, name,  image_path, slug, min_price, agg_rating, review_count, up_time_to_delivery, down_time_to_delivery FROM restaurants WHERE slug = $1", req.Slug)
	switch err {
	case nil:
		return restaurant, nil
	case sql.ErrNoRows:
		return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, err.Error())
	default:
		return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
	}
}

func (r *CommentsRepo) GetRestaurantComments(req models.GetRestaurantCommentsRepoReq) (*models.CommentsRestaurantDataStorage, error) {
	comments := make([]*models.CommentRestaurantDataStorage, 0)
	err := r.DB.Select(&comments, `SELECT restaurant_id, author, text, stars, get_ru_date(date) FROM comments WHERE restaurant_id = $1 ORDER BY date DESC`, req.Id)

	switch err {
	case nil:
		resp := &models.CommentsRestaurantDataStorage{Comments: make([]models.CommentRestaurantDataStorage, len(comments))}
		for i, dish := range comments {
			resp.Comments[i] = *dish
		}
		return resp, nil
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

func (r *CommentsRepo) AddRestaurantComment(req models.AddRestaurantCommentRepoReq) (*models.CommentRestaurantDataStorage, error) {
	comment := &models.CommentRestaurantDataStorage{}
	err := r.DB.Get(comment, `INSERT INTO comments (restaurant_id, author, text, stars) VALUES ($1,$2,$3,$4) RETURNING restaurant_id, author, text, stars, get_ru_date(date)`, req.RestaurantId, req.User, req.CommentText, req.CommentRating)
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

func (r *CommentsRepo) UpdateRestaurantRating(req models.UpdateRestaurantRatingRepoReq) (*models.RestaurantRepo, error) {
	restaurant := &models.RestaurantRepo{}
	err := r.DB.Get(restaurant, `UPDATE restaurants SET agg_rating=$1, review_count=$2 WHERE id=$3 RETURNING id, name, image_path, slug, image_path, slug, min_price, up_time_to_delivery, down_time_to_delivery, review_count, agg_rating`, req.NewRestRating, req.CountRating, req.RestId)
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
