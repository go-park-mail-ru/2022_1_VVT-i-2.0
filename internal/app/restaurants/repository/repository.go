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

//func (r *RestaurantsRepo) GetRestaurants() ([]*models.RestaurantDataStorage, error) {
//	restaurants := make([]*models.RestaurantDataStorage, 0, 21)
//	rows, err := r.DB.Query(`SELECT * FROM restaurants`)
//	defer rows.Close()
//
//	for rows.Next() {
//		restaurant := &models.RestaurantDataStorage{}
//		err = rows.Scan(
//			&restaurant.Id,
//			&restaurant.Name,
//			&restaurant.City,
//			&restaurant.Address,
//			&restaurant.Image_path,
//			&restaurant.Slug,
//			&restaurant.Min_price,
//			&restaurant.Avg_price,
//			&restaurant.Rating)
//		if err != nil {
//			return nil, err
//		}
//		restaurants = append(restaurants, restaurant)
//	}
//
//	return restaurants, err
//	switch err {
//	case nil:
//		return restaurants, nil
//	case sql.ErrNoRows:
//		return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, err.Error())
//	default:
//		return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
//	}
//}

func (r *RestaurantsRepo) GetRestaurants() (*models.RestaurantsDataStorage, error) {
	restaurants := &models.RestaurantsDataStorage{}
	rows, err := r.DB.Query(`SELECT * FROM restaurants`)
	defer rows.Close()

	for rows.Next() {
		restaurant := &models.RestaurantDataStorage{}
		err = rows.Scan(
			&restaurant.Id,
			&restaurant.Name,
			&restaurant.City,
			&restaurant.Address,
			&restaurant.Image_path,
			&restaurant.Slug,
			&restaurant.Min_price,
			&restaurant.Avg_price,
			&restaurant.Rating)
		if err != nil {
			return nil, err
		}
		restaurants.Restaurants = append(restaurants.Restaurants, restaurant)
	}

	return restaurants, err
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

func (r *RestaurantsRepo) GetDishByRestaurants(id int) (*models.DishesDataStorage, error) {

	dishes := &models.DishesDataStorage{}
	rows, err := r.DB.Query(`select * from dish where restaurant = $1`, id)
	defer rows.Close()

	for rows.Next() {
		dish := &models.Dish{}
		err = rows.Scan(
			&dish.Id,
			&dish.Restaurant,
			&dish.Name,
			&dish.Description,
			&dish.Image_path,
			&dish.Image_path,
			&dish.Calories,
			&dish.Price)
		if err != nil {
			return nil, err
		}
		dishes.Dishes = append(dishes.Dishes, dish)
	}

	switch err {
	case nil:
		return dishes, nil
	case sql.ErrNoRows:
		return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, err.Error())
	default:
		return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
	}
}