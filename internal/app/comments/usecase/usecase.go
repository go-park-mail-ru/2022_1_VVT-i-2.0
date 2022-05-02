package usecase

import (
	"fmt"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/comments"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/pkg/errors"
)

type CommentsUsecase struct {
	Repository comments.Repository
}

func NewCommentsUsecase(commentsRepo comments.Repository) *CommentsUsecase {
	return &CommentsUsecase{
		Repository: commentsRepo,
	}
}

func (u *CommentsUsecase) GetRestaurantComments(slug string) (*models.CommentsRestaurantUseCase, error) {
	restaurant := &models.RestaurantDataStorage{}
	restaurant, err := u.Repository.GetRestaurantBySlug(slug)
	if err != nil {
		fmt.Println("сломалась тут2")
		return nil, errors.Wrapf(err, "error getting restaurant")
	}

	commentsData, err := u.Repository.GetRestaurantComments(restaurant.Id)
	if err != nil {
		fmt.Println("сломалась тут3")
		return nil, errors.Wrapf(err, "error getting comments")
	}

	commentsUC := &models.CommentsRestaurantUseCase{}

	for _, comment := range commentsData {
		item := &models.CommentRestaurantUseCase{
			Author:        	comment.Author,
			Text:   		comment.Text,
			Stars: 			comment.Stars,
			Date: 			comment.Date,
		}
		commentsUC.Comment = append(commentsUC.Comment, *item)
	}
	return commentsUC, nil
}

func (u *CommentsUsecase) AddRestaurantComment(id models.UserId, item *models.AddCommentRestaurantUseCase) (*models.CommentRestaurantUseCase, error) {
	userData, err := u.Repository.GetUserById(id)
	if err != nil {
		return nil, errors.Wrapf(err, "error getting user by id %d", id)
	}

	comment, err := u.Repository.AddRestaurantComment(&models.AddCommentRestaurantDataStorage{
		Restaurant_id:	item.Restaurant,
		User: 			userData.Name,
		Comment_text:   item.Comment_text,
		Comment_rating: item.Comment_rating,
	})
	if err != nil {
		return nil, errors.Wrap(err, "error adding user to storage")
	}

	restaurant, err := u.Repository.GetRestaurantByID(comment.Restaurant_id)
	if err != nil {
		return nil, errors.Wrap(err, "error adding user to storage")
	}

	_, err = u.Repository.UpdateRestaurantRating(comment.Restaurant_id, comment.Stars+restaurant.AggRating, restaurant.ReviewCount+1)
	if err != nil {
		return nil, errors.Wrap(err, "error adding user to storage")
	}

	return &models.CommentRestaurantUseCase{
		Restaurant_id:	comment.Restaurant_id,
		Author:        	comment.Author,
		Text:   		comment.Text,
		Stars: 			comment.Stars,
		Date: 			comment.Date,
	}, nil
}

//func (u *CommentsUsecase) AddRestaurantComment(item *models.AddCommentRestaurantUseCase) (*models.CommentRestaurantUseCase, error) {
//	comment, err := u.Repository.AddRestaurantComment(&models.AddCommentRestaurantDataStorage{
//		Restaurant:     item.Restaurant,
//		User_id:        item.User_id,
//		Comment_text:   item.Comment_text,
//		Comment_rating: item.Comment_rating,
//	})
//	if err != nil {
//		return nil, errors.Wrap(err, "error adding user to storage")
//	}
//
//	restaurant, err := u.Repository.GetRestaurantByID(comment.Restaurant)
//	if err != nil {
//		return nil, errors.Wrap(err, "error adding user to storage")
//	}
//
//	_, err = u.Repository.UpdateRestaurantRating(comment.Restaurant, comment.Comment_rating+restaurant.AggRating, restaurant.ReviewCount+1)
//	if err != nil {
//		return nil, errors.Wrap(err, "error adding user to storage")
//	}
//
//	return &models.CommentRestaurantUseCase{
//		Id:             comment.Id,
//		Restaurant:     comment.Restaurant,
//		User_id:        comment.User_id,
//		Comment_text:   comment.Comment_text,
//		Comment_rating: comment.Comment_rating,
//	}, nil
//}
