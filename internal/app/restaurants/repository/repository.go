package repository

import (
	"database/sql"
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

func (r *RestaurantsRepo) GetRestaurants() ([]*models.Restaurant, error) {
	restaurants := make([]*models.Restaurant, 0, 21)
	err := r.DB.Select(restaurants, `SELECT * FROM restaurants`)
	switch err {
	case nil:
		return restaurants, nil
	case sql.ErrNoRows:
		return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, err.Error())
	default:
		return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
	}
}

func (r *RestaurantsRepo) GetRestaurantsBySlug(slug string) (*models.Restaurant, error) {
	restaurant := &models.Restaurant{}
	err := r.DB.Get(restaurant, `select * from restaurants where slug = $1`, slug)
	switch err {
	case nil:
		return restaurant, nil
	case sql.ErrNoRows:
		return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, err.Error())
	default:
		return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
	}
}

func (r *RestaurantsRepo) GetDishByRestaurants(id int) ([]*models.Dish, error) {
	dishes := make([]*models.Dish, 0, 21)
	err := r.DB.Select(dishes, `select * from dish where restaurant = $1`, id)
	switch err {
	case nil:
		return dishes, nil
	case sql.ErrNoRows:
		return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, err.Error())
	default:
		return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
	}
}

//func (r *RestaurantsRepo) GetDishByRestaurants(slug string) (*models.Restaurant, *models.Dish, error) {
//	restaurants := &models.Restaurant{}
//	err := r.DB.Get(restaurants, `select * from restaurants where slug = $1`, slug)
//	if err == sql.ErrNoRows {
//		return nil, nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, err.Error())
//	}
//	dish := &models.Dish{}
//	err = r.DB.Get(dish, `select * from dish where restaurant = $1`, restaurants.Id)
//	if err == sql.ErrNoRows {
//		return nil, nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, err.Error())
//	}
//	return restaurants, dish, nil
//}