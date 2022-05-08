package mock

import (
	"database/sql"

	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/stretchr/testify/mock"
)

type CommentsRepository struct {
	mock.Mock
}

func (r *CommentsRepository) GetRestaurantBySlug(slug string) (*models.RestaurantRepo, error) {
	if slug == "" {
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

func (r *CommentsRepository) GetRestaurantByID(id int) (*models.RestaurantRepo, error) {
	if id == 0 {
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

func (r *CommentsRepository) GetRestaurantComments(id int) ([]*models.CommentRestaurantDataStorage, error) {
	return []*models.CommentRestaurantDataStorage{{
		RestaurantId: 1,
		Author:       "author",
		Text:         "text",
		Stars:        4,
		Date:         "date"}}, nil
}

func (r *CommentsRepository) AddRestaurantComment(item *models.AddCommentRestaurantDataStorage) (*models.CommentRestaurantDataStorage, error) {
	if item == nil {
		return nil, nil
	}
	return &models.CommentRestaurantDataStorage{
		RestaurantId: 1,
		Author:       "author",
		Text:         "text",
		Stars:        4,
		Date:         "date"}, nil
}

func (r *CommentsRepository) UpdateRestaurantRating(restId int, newRestRating int, countRating int) (*models.RestaurantRepo, error) {
	if restId == 0 && newRestRating == 0 && countRating == 0 {
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
