package comments

import "github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"

type Repository interface {
	GetRestaurantByID(req models.GetRestaurantByIdRepoReq) (*models.RestaurantRepo, error)
	GetRestaurantBySlug(req models.GetRestaurantBySlugRepoReq) (*models.RestaurantRepo, error)
	GetRestaurantComments(req models.GetRestaurantCommentsRepoReq) (*models.CommentsRestaurantDataStorage, error)
	GetUserById(id models.UserId) (*models.UserDataRepo, error)
	AddRestaurantComment(req models.AddRestaurantCommentRepoReq) (*models.CommentRestaurantDataStorage, error)
	UpdateRestaurantRating(req models.UpdateRestaurantRatingRepoReq) (*models.RestaurantRepo, error)
}
