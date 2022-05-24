package mock

import (
	"database/sql"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/stretchr/testify/mock"
)

type CommentsRepository struct {
	mock.Mock
}

func (r *CommentsRepository) GetRestaurantBySlug(req models.GetRestaurantBySlugRepoReq) (*models.RestaurantRepo, error) {
	if req.Slug == "" {
		return nil, nil
	}
	return &models.RestaurantRepo{
		Id:                   1,
		Name:                 "Name",
		ImagePath:            "imgPath",
		Slug:                 "slug",
		MinPrice:             1,
		AggRating:            1,
		ReviewCount:          1,
		UpMinutsToDelivery:   1,
		DownMinutsToDelivery: 1}, nil
}

func (r *CommentsRepository) GetUserById(id models.UserId) (*models.UserDataRepo, error) {
	if id == 0 {
		return nil, nil
	}
	user := &models.UserDataRepo{
		Id:     1,
		Name:   "name",
		Phone:  "89166152595",
		Email:  "seregey.golubev@mail.ru",
		Avatar: sql.NullString{String: "avatar", Valid: true},
	}
	return user, nil
}

func (r *CommentsRepository) GetRestaurantByID(req models.GetRestaurantByIdRepoReq) (*models.RestaurantRepo, error) {
	if req.Id == 0 {
		return nil, nil
	}
	return &models.RestaurantRepo{
		Id:                   1,
		Name:                 "Name",
		ImagePath:            "imgPath",
		Slug:                 "slug",
		MinPrice:             1,
		AggRating:            1,
		ReviewCount:          1,
		UpMinutsToDelivery:   1,
		DownMinutsToDelivery: 1}, nil
}

func (r *CommentsRepository) GetRestaurantComments(req models.GetRestaurantCommentsRepoReq) (*models.CommentsRestaurantDataStorage, error) {
	return &models.CommentsRestaurantDataStorage{Comments: []models.CommentRestaurantDataStorage{{
		RestaurantId: 1,
		Author:       "author",
		Text:         "text",
		Stars:        4,
		Date:         "date",
	}}}, nil
}

func (r *CommentsRepository) AddRestaurantComment(req models.AddRestaurantCommentRepoReq) (*models.CommentRestaurantDataStorage, error) {
	return &models.CommentRestaurantDataStorage{
		RestaurantId: 1,
		Author:       "author",
		Text:         "text",
		Stars:        4,
		Date:         "date"}, nil
}

func (r *CommentsRepository) UpdateRestaurantRating(req models.UpdateRestaurantRatingRepoReq) (*models.RestaurantRepo, error) {
	return &models.RestaurantRepo{
		Id:                   1,
		Name:                 "Name",
		ImagePath:            "imgPath",
		Slug:                 "slug",
		MinPrice:             1,
		AggRating:            1,
		ReviewCount:          1,
		UpMinutsToDelivery:   1,
		DownMinutsToDelivery: 1}, nil
}
