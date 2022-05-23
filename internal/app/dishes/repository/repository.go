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

func (r *DishesRepo) GetRestaurantBySlug(req models.GetRestaurantBySlugRepoReq) (*models.RestaurantRepo, error) {
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

func (r *DishesRepo) GetCategories(req models.GetCategoriesByIdRepoReq) (*models.Categories, error) {
	var tags []string
	if err := r.DB.QueryRow(`SELECT categories FROM restaurants WHERE id = $1`, req.Id).Scan(pq.Array(&tags)); err != nil {
		return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, err.Error())
	}
	categories := &models.Categories{}
	categories.Categories = tags
	return categories, nil
}

func (r *DishesRepo) GetRestaurantDishes(req models.GetRestaurantDishesRepoReq) (*models.GetRestaurantDishesCategoriesRepoResp, error) {
	dishes := make([]*models.DishCategoriesRepo, 0)
	err := r.DB.Select(&dishes, "SELECT id, restaurant_id, categori, name, description, image_path, calories, price, weight FROM dishes WHERE restaurant_id = $1", req.Id)
	switch err {
	case nil:
		resp := &models.GetRestaurantDishesRepoResp{Dishes: make([]models.DishRepo, len(dishes))}
		for i, dish := range dishes {
			resp.Dishes[i] = *dish
		}
		return resp, nil
	case sql.ErrNoRows:
		return nil, servErrors.NewError(servErrors.NO_SUCH_ENTITY_IN_DB, err.Error())
	default:
		return nil, servErrors.NewError(servErrors.DB_ERROR, err.Error())
	}
}
