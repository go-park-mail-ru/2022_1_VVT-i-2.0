package repository

import (
	"database/sql"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/tools/servErrors"
	"github.com/jmoiron/sqlx"
)

type RecommendationsRepo struct {
	DB *sqlx.DB
}

func NewRecommendationsRepo(db *sqlx.DB) *RecommendationsRepo {
	return &RecommendationsRepo{DB: db}
}

func (r *RecommendationsRepo) GetRestaurantDishes(req models.GetRestaurantDishesRepoReq) (*models.GetRestaurantDishesCategoriesRepoResp, error) {
	dishes := make([]*models.DishCategoriesRepo, 0)
	err := r.DB.Select(&dishes, "SELECT id, restaurant_id, category, name, description, image_path, calories, price, weight FROM dishes WHERE restaurant_id = $1", req.Id)
	switch err {
	case nil:
		resp := &models.GetRestaurantDishesCategoriesRepoResp{Dishes: make([]models.DishCategoriesRepo, len(dishes))}
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
