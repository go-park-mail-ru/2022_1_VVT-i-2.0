package usecase

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/comments"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/pkg/errors"
)

type CommentsUsecase struct {
	Repository comments.Repository
}

func NewRestaurantsUsecase(commensRepo comments.Repository) *CommentsUsecase {
	return &CommentsUsecase{
		Repository: commensRepo,
	}
}

func (u *CommentsUsecase) GetRestaurantComments(id int) (*models.CommentsRestaurantUseCase, error) {
	commentsData, err := u.Repository.GetRestaurantComments(id)
	if err != nil {
		return nil, errors.Wrapf(err, "error getting restaurants")
	}

	commentsUC := &models.CommentsRestaurantUseCase{}

	for _, comment := range commentsData {
		item := &models.CommentRestaurantUseCase{
			Id:             comment.Id,
			Restaurant:     comment.Restaurant,
			User_id:        comment.User_id,
			Comment_text:   comment.Comment_text,
			Comment_rating: comment.Comment_rating,
		}
		commentsUC.Comment = append(commentsUC.Comment, *item)
	}
	return commentsUC, nil
}

func (u *CommentsUsecase) AddRestaurantComment(item *models.AddCommentRestaurantUseCase) (*models.CommentRestaurantUseCase, error) {
	comment, err := u.Repository.AddRestaurantComment(&models.AddCommentRestaurantDataStorage{
		Restaurant:     item.Restaurant,
		User_id:        item.User_id,
		Comment_text:   item.Comment_text,
		Comment_rating: item.Comment_rating,
	})
	if err != nil {
		return nil, errors.Wrap(err, "error adding user to storage")
	}

	restaurant, err := u.Repository.GetRestaurantByID(comment.Restaurant)
	if err != nil {
		return nil, errors.Wrap(err, "error adding user to storage")
	}

	_, err = u.Repository.UpdateRestaurantRating(comment.Restaurant, comment.Comment_rating+restaurant.AggRating, restaurant.ReviewCount+1)
	if err != nil {
		return nil, errors.Wrap(err, "error adding user to storage")
	}

	return &models.CommentRestaurantUseCase{
		Id:             comment.Id,
		Restaurant:     comment.Restaurant,
		User_id:        comment.User_id,
		Comment_text:   comment.Comment_text,
		Comment_rating: comment.Comment_rating,
	}, nil
}
