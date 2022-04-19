package repository

import (
	"database/sql"
	"fmt"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/jmoiron/sqlx"
)

type RestaurantsRepo struct {
	DB *sqlx.DB
}

func NewRestaurantsRepo(db *sqlx.DB) *RestaurantsRepo {
	return &RestaurantsRepo{DB: db}
}

func (r *RestaurantsRepo) GetRestaurants() ([]*models.RestaurantDataStorage, error) {
	restaurantsDS := make([]*models.RestaurantDataStorage, 0, 21)
	err := r.DB.Select(&restaurantsDS, "SELECT id, name, city, address, image_path, slug, min_price, avg_price, rating FROM restaurants")
	switch err {
	case nil:
		return restaurantsDS, nil
	case sql.ErrNoRows:
		return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, err.Error())
	default:
		return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
	}
}

func (r *RestaurantsRepo) GetRestaurantsBySlug(slug string) (*models.RestaurantDataStorage, error) {
	restaurant := &models.RestaurantDataStorage{}
	err := r.DB.Get(restaurant, "SELECT id, name, city, address, image_path, slug, min_price, avg_price, rating FROM restaurants WHERE slug = $1", slug)
	switch err {
	case nil:
		return restaurant, nil
	case sql.ErrNoRows:
		return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, err.Error())
	default:
		return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
	}
}

func (r *RestaurantsRepo) GetDishByRestaurants(id int) ([]*models.DishDataStorage, error) {
	dishes := make([]*models.DishDataStorage, 0, 21)
	err := r.DB.Select(&dishes, "SELECT id, restaurant, name, description, image_path, calories, price FROM dish WHERE restaurant = $1", id)
	switch err {
	case nil:
		return dishes, nil
	case sql.ErrNoRows:
		return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, err.Error())
	default:
		return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
	}
}

func (r *RestaurantsRepo) GetCommentsRestaurantByRestaurants(id int) ([]*models.CommentRestaurantDataStorage, error) {
	fmt.Println(id)
	comments := make([]*models.CommentRestaurantDataStorage, 0, 2)
	err := r.DB.Select(&comments, `SELECT id, restaurant, user_id, comment_text FROM comment_restaurants WHERE id = $1`, id)

	switch err {
	case nil:
		return comments, nil
	case sql.ErrNoRows:
		return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, err.Error())
	default:
		return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
	}
}

func (r *RestaurantsRepo) AddCommentsRestaurantByRestaurants(newUser *models.AddCommentRestaurantDataStorage) (bool, error) {
	var newUserId int64
	err := r.DB.QueryRow(`INSERT INTO comment_restaurants (restaurant, user_id, comment_text) VALUES ($1,$2,$3) RETURNING id`, newUser.Restaurant, newUser.User_id, newUser.Comment_text).Scan(&newUserId)
	fmt.Println(err)
	if err != nil {
		if err == sql.ErrConnDone || err == sql.ErrTxDone {
			return true, servErrors.NewError(servErrors.DB_ERROR, err.Error())
		}
		return true, servErrors.NewError(servErrors.DB_INSERT, err.Error())
	}
	if newUserId == 0 {
		return true, servErrors.NewError(servErrors.DB_INSERT, "")
	}
	return true, nil
}