package repository

import (
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

func (r *RestaurantsRepo) GetRestaurants() (*models.RestaurantsRepo, error) {
	restaurants := make([]*models.RestaurantRepo, 0)
	err := r.DB.Select(&restaurants, "SELECT id, name,  image_path, slug, min_price, agg_rating, review_count, up_time_to_delivery, down_time_to_delivery FROM restaurants ORDER BY CASE WHEN review_count=0 THEN 0 ELSE agg_rating::float/review_count END DESC")
	switch err {
	case nil:
		resp := models.RestaurantsRepo{Restaurants: make([]models.RestaurantRepo, len(restaurants))}
		for i, restaurant := range restaurants {
			resp.Restaurants[i] = *restaurant
		}
		return &resp, nil
	default:
		return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
	}
}

func (r *RestaurantsRepo) GetRestaurantsByCategory(category models.GetRestaurantByCategoryRepoReq) (*models.RestaurantsRepo, error) {
	restaurants := make([]*models.RestaurantRepo, 0)
	err := r.DB.Select(&restaurants, `SELECT r.id id, r.name, r.image_path image_path, r.slug slug, r.min_price min_price, r.agg_rating agg_rating, r.review_count review_count, up_time_to_delivery, down_time_to_delivery FROM restaurants r JOIN categori_restaurant cr ON r.id=cr.restaurant_id JOIN categories c ON cr.categori_id=c.id WHERE c.name=$1 ORDER BY CASE WHEN review_count=0 THEN 0 ELSE agg_rating::float/review_count END DESC`, category.Name)
	switch err {
	case nil:
		resp := models.RestaurantsRepo{Restaurants: make([]models.RestaurantRepo, len(restaurants))}
		for i, restaurant := range restaurants {
			resp.Restaurants[i] = *restaurant
		}
		return &resp, nil
	default:
		return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
	}
}

func (r *RestaurantsRepo) GetRestaurantsBySearchQuery(query models.GetRestaurantBySearchQueryRepoReq) (*models.RestaurantsRepo, error) {
	restaurants := make([]*models.RestaurantRepo, 0)
	err := r.DB.Select(&restaurants, `SELECT r.id id, r.name, r.image_path image_path, r.slug slug, r.min_price min_price, r.agg_rating agg_rating, r.review_count review_count, up_time_to_delivery, down_time_to_delivery FROM restaurants r JOIN categori_restaurant cr ON r.id=cr.restaurant_id JOIN categories c ON cr.categori_id=c.id WHERE c.name ILIKE $1 ORDER BY CASE WHEN review_count=0 THEN 0 ELSE agg_rating::float/review_count END DESC`, "%"+query.Query+"%")
	if len(restaurants) == 0 {
		err = r.DB.Select(&restaurants, `SELECT id id, name, image_path image_path, slug slug, min_price min_price, agg_rating agg_rating, review_count review_count, up_time_to_delivery, down_time_to_delivery FROM restaurants WHERE name ILIKE $1 ORDER BY CASE WHEN review_count=0 THEN 0 ELSE agg_rating::float/review_count END DESC`, "%"+query.Query+"%")
	}
	switch err {
	case nil:
		resp := models.RestaurantsRepo{Restaurants: make([]models.RestaurantRepo, len(restaurants))}
		for i, restaurant := range restaurants {
			resp.Restaurants[i] = *restaurant
		}
		return &resp, nil
	default:
		return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
	}
}
