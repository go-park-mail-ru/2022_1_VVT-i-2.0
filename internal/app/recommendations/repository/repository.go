package repository

import (
	"database/sql"
	"strconv"

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

func generateRecommendationsQuery(req *models.RecommendationsRepoReq) string {
	if len(req.DishesId) == 0 {
		return `SELECT DISTINCT ON (image_path) image_path, id, restaurant_id, category, name, description, calories, price, weight FROM dishes WHERE restaurant_id = $1 LIMIT $2`
	}
	query := `SELECT DISTINCT ON (image_path) image_path, id, restaurant_id, category, name, description, calories, price, weight FROM dishes WHERE restaurant_id = $1 AND category NOT IN(SELECT category FROM dishes WHERE id IN ($2`

	nextPlaceholderNum := 3
	i := 0
	for ; i < len(req.DishesId)-1; i++ {
		query += `,$` + strconv.Itoa(nextPlaceholderNum+i)
	}

	query += `)) LIMIT $` + strconv.Itoa(nextPlaceholderNum+i)
	return query
}

func expandRecommendationReqArgs(req *models.RecommendationsRepoReq) []interface{} {
	args := make([]interface{}, 0, len(req.DishesId)+2)
	args = append(args, req.RestId)
	for i := 1; i <= len(req.DishesId); i++ {
		args = append(args, req.DishesId[i-1])
	}
	args = append(args, req.Limit)
	return args
}

func (r *RecommendationsRepo) GetRecommendations(req *models.RecommendationsRepoReq) (*models.RecommendationsRepoResp, error) {
	dishes := make([]*models.RecommendationRepo, 0, req.Limit)
	err := r.DB.Select(&dishes, generateRecommendationsQuery(req), expandRecommendationReqArgs(req)...)
	switch err {
	case nil:
		resp := &models.RecommendationsRepoResp{Dishes: make([]models.RecommendationRepo, len(dishes))}
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
