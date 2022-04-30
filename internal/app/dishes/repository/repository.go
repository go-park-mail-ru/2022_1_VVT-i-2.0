package repository

import (
	"database/sql"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/jmoiron/sqlx"
)

type DishesRepo struct {
	DB *sqlx.DB
}

func NewDishesRepo(db *sqlx.DB) *DishesRepo {
	return &DishesRepo{DB: db}
}

func (r *DishesRepo) GetRestaurantBySlug(slug string) (*models.RestaurantDataStorage, error) {
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

func (r *DishesRepo) GetDishesByRestaurant(id int) ([]*models.DishDataStorage, error) {
	dishes := make([]*models.DishDataStorage, 0, 21)
	err := r.DB.Select(&dishes, "SELECT id, restaurant_id, name, description, image_path, calories, price, weight FROM dishes WHERE restaurant_id = $1", id)
	switch err {
	case nil:
		return dishes, nil
	case sql.ErrNoRows:
		return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, err.Error())
	default:
		return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
	}
}
