package usecase

import (
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

func (u *CommentsUsecase) GetRestaurantComments(req models.GetRestaurantCommentsUcaseReq) (*models.CommentsRestaurantUseCase, error) {
	restaurant, err := u.Repository.GetRestaurantBySlug(models.GetRestaurantBySlugRepoReq{
		Slug: req.Slug,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "error getting restaurant")
	}

	commentsData, err := u.Repository.GetRestaurantComments(models.GetRestaurantCommentsRepoReq{
		Id: models.Id(restaurant.Id),
	})
	if err != nil {
		return nil, errors.Wrapf(err, "error getting comments")
	}

	commentsUC := &models.CommentsRestaurantUseCase{}

	for _, comment := range commentsData.Comments {
		item := &models.CommentRestaurantUseCase{
			RestaurantId: comment.RestaurantId,
			Author:       comment.Author,
			Text:         comment.Text,
			Stars:        comment.Stars,
			Date:         comment.Date,
		}
		commentsUC.Comment = append(commentsUC.Comment, *item)
	}
	return commentsUC, nil
}

func (u *CommentsUsecase) AddRestaurantComment(req models.AddCommentRestaurantUcaseReq) (*models.CommentRestaurantUseCase, error) {
	userData, err := u.Repository.GetUserById(req.UserId)
	if err != nil {
		return nil, errors.Wrapf(err, "error getting user by id %d", req.UserId)
	}

	RestaurantData, err := u.Repository.GetRestaurantBySlug(models.GetRestaurantBySlugRepoReq{
		Slug: req.Slug,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "error getting restaurants by slug %s", req.Slug)
	}

	comment, err := u.Repository.AddRestaurantComment(models.AddRestaurantCommentRepoReq{
		RestaurantId:  RestaurantData.Id,
		User:          userData.Name,
		CommentText:   req.CommentText,
		CommentRating: req.CommentRating,
	})
	if err != nil {
		return nil, errors.Wrap(err, "error adding user to storage")
	}

	restaurant, err := u.Repository.GetRestaurantByID(models.GetRestaurantByIdRepoReq{
		Id: models.Id(comment.RestaurantId),
	})
	if err != nil {
		return nil, errors.Wrap(err, "error adding user to storage")
	}

	_, err = u.Repository.UpdateRestaurantRating(models.UpdateRestaurantRatingRepoReq{
		RestId: comment.RestaurantId,
		NewRestRating: comment.Stars+restaurant.AggRating,
		CountRating: restaurant.ReviewCount+1,
	})
	if err != nil {
		return nil, errors.Wrap(err, "error adding user to storage")
	}

	return &models.CommentRestaurantUseCase{
		RestaurantId: comment.RestaurantId,
		Author:       comment.Author,
		Text:         comment.Text,
		Stars:        comment.Stars,
		Date:         comment.Date,
	}, nil
}