package usecase

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/recommendations"
	"github.com/pkg/errors"
)

type RecommendationsUcase struct {
	Repo recommendations.Repository
}

func NewRecommendationsUcase(recommendationsRepo recommendations.Repository) *RecommendationsUcase {
	return &RecommendationsUcase{
		Repo: recommendationsRepo,
	}
}

func (u *RecommendationsUcase) GetRecommendations(req *models.RecommendationsUcaseReq) (*models.RecommendationsUcaseResp, error) {
	recommendationsRepoResp, err := u.Repo.GetRecommendations(&models.RecommendationsRepoReq{RestId: req.RestId, Limit: req.Limit, DishesId: req.DishesId})
	if err != nil {
		return nil, errors.Wrap(err, "error getting restaurant dishes")
	}

	if len(recommendationsRepoResp.Dishes) == 0 {
		return nil, nil
	}

	RecommendationsResp := &models.RecommendationsUcaseResp{Dishes: make([]models.RecommendationUcase, 0, req.Limit)}

	for _, recommendation := range recommendationsRepoResp.Dishes {
		RecommendationsResp.Dishes = append(RecommendationsResp.Dishes, models.RecommendationUcase(recommendation))
	}

	return RecommendationsResp, nil
}
