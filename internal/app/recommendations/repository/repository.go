package repository

import "github.com/jmoiron/sqlx"

type RecommendationsRepo struct {
	DB *sqlx.DB
}

func NewRecommendationsRepo(db *sqlx.DB) *RecommendationsRepo {
	return &RecommendationsRepo{DB: db}
}
