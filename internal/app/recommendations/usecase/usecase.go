package usecase

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/recommendations"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/restaurants"
)

type RecommendationsUcase struct {
	RecommendationsRepo recommendations.Repository
}

func NewRecommendationsUcase(recommendationsRepo restaurants.Repository) *RecommendationsUcase {
	return &RecommendationsUcase{
		RecommendationsRepo: recommendationsRepo,
	}
}
