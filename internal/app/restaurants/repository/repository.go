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

func (r *RestaurantsRepo) GetRestaurants() ([]*models.RestaurantDataStorage, error) {
	restaurantsDS := make([]*models.RestaurantDataStorage, 0)
	err := r.DB.Select(&restaurantsDS, "SELECT id, name,  image_path, slug, min_price, agg_rating, review_count, up_time_to_delivery, down_time_to_delivery FROM restaurants")
	switch err {
	case nil:
		return restaurantsDS, nil
	case sql.ErrNoRows:
		return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, err.Error())
	default:
		return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
	}
}

func (r *RestaurantsRepo) GetRestaurantsByCategory(category models.GetRestaurantByCategoryRepoReq) ([]*models.RestaurantDataStorage, error) {
	restaurantsDS := make([]*models.RestaurantDataStorage, 0)
	err := r.DB.Select(&restaurantsDS, `SELECT r.id id, r.name, r.image_path image_path, r.slug slug, r.min_price min_price, r.agg_rating agg_rating, r.review_count review_count, up_time_to_delivery, down_time_to_delivery FROM restaurants r JOIN categori_restaurant cr ON r.id=cr.restaurant_id JOIN categories c ON cr.categori_id=c.id WHERE c.name=$1`, category.Name)
	switch err {
	case nil:
		return restaurantsDS, nil
	case sql.ErrNoRows:
		return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, err.Error())
	default:
		return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
	}
}
