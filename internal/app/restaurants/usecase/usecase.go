package usecase

import (
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/models"
	"github.com/go-park-mail-ru/2022_1_VVT-i-2.0/internal/app/restaurants"
	"github.com/pkg/errors"
)

type RestaurantsUsecase struct {
	RestaurantsRepo    	restaurants.Repository
}

func NewRestaurantsUsecase(restaurantsRepo restaurants.Repository) *RestaurantsUsecase {
	return &RestaurantsUsecase{
		RestaurantsRepo:    restaurantsRepo,
	}
}

func (u *RestaurantsUsecase) GetAllRestaurants() (*models.RestaurantsUsecase, error) {
	restaurantsData, err := u.RestaurantsRepo.GetRestaurants()
	if err != nil {
		return nil, errors.Wrapf(err, "error getting restaurants")
	}

	restaurantsUC := &models.RestaurantsUsecase{}

	for _, rest := range restaurantsData {
		item := &models.RestaurantUsecase{
			Id:         rest.Id,
			Name:       rest.Name,
			City:       rest.City,
			Address:    rest.Address,
			Image_path: rest.Image_path,
			Slug:       rest.Slug,
			Min_price:  rest.Min_price,
			Avg_price:  rest.Avg_price,
			Rating: rest.Rating,
			Count_rating: rest.Count_rating,
		}
		restaurantsUC.Restaurants = append(restaurantsUC.Restaurants, *item)
	}

	return restaurantsUC, nil
}

func (u *RestaurantsUsecase) GetRestaurantBySluf(slug string) (*models.RestaurantUsecase, error) {
	restaurantData, err := u.RestaurantsRepo.GetRestaurantBySlug(slug)
	if err != nil {
		return nil, errors.Wrapf(err, "error getting restaurants")
	}
	return &models.RestaurantUsecase{
		Id:         restaurantData.Id,
		Name:       restaurantData.Name,
		City:       restaurantData.City,
		Address:    restaurantData.Address,
		Image_path: restaurantData.Image_path,
		Slug:       restaurantData.Slug,
		Min_price:  restaurantData.Min_price,
		Avg_price:  restaurantData.Avg_price,
		Rating: restaurantData.Rating,
		Count_rating: restaurantData.Count_rating,
	}, nil
}

func (u *RestaurantsUsecase) GetDishByRestaurant(id int) (*models.DishesUseCase, error) {
	dishesData, err := u.RestaurantsRepo.GetDishByRestaurants(id)
	if err != nil {
		return nil, errors.Wrapf(err, "error getting restaurants")
	}

	dishesUC := &models.DishesUseCase{}

	for _, dish := range dishesData {
		item := &models.DishUseCase{
			Id: dish.Id,
			Restaurant: dish.Restaurant,
			Name: dish.Name,
			Description: dish.Description,
			Image_path: dish.Image_path,
			Calories: dish.Calories,
			Weight: dish.Weight,
			Price: dish.Price,
		}
		dishesUC.Dishes = append(dishesUC.Dishes, *item)
	}

	return dishesUC, nil
}

func (u *RestaurantsUsecase) GetCommentsRestaurantByRestaurants(id int) (*models.CommentsRestaurantUseCase, error) {
	commentsData, err := u.RestaurantsRepo.GetCommentsRestaurantByRestaurants(id)
	if err != nil {
		return nil, errors.Wrapf(err, "error getting restaurants")
	}

	commentsUC := &models.CommentsRestaurantUseCase{}

	for _, comment := range commentsData {
		item := &models.CommentRestaurantUseCase{
			Id: comment.Id,
			Restaurant: comment.Restaurant,
			User_id: comment.User_id,
			Comment_text: comment.Comment_text,
			Comment_rating: comment.Comment_rating,
		}
		commentsUC.Comment = append(commentsUC.Comment, *item)
	}

	return commentsUC, nil
}

func (u *RestaurantsUsecase) AddCommentsRestaurantByRestaurants(item *models.AddCommentRestaurantUseCase) (*models.CommentRestaurantUseCase, error) {
	comment, err := u.RestaurantsRepo.AddCommentsRestaurantByRestaurants(&models.AddCommentRestaurantDataStorage{
		Restaurant: item.Restaurant,
		User_id: item.User_id,
		Comment_text: item.Comment_text,
		Comment_rating: item.Comment_rating,
	})
	if err != nil {
		return nil, errors.Wrap(err, "error adding user to storage")
	}

	restaurant, err := u.RestaurantsRepo.GetRestaurantByID(comment.Restaurant)
	if err != nil {
		return nil, errors.Wrap(err, "error adding user to storage")
	}

	restaurant, err = u.RestaurantsRepo.UpdateRestaurantRating(comment.Restaurant, comment.Comment_rating + restaurant.Rating, restaurant.Count_rating + 1)
	if err != nil {
		return nil, errors.Wrap(err, "error adding user to storage")
	}

	return &models.CommentRestaurantUseCase{
		Id: comment.Id,
		Restaurant: comment.Restaurant,
		User_id: comment.User_id,
		Comment_text: comment.Comment_text,
		Comment_rating: comment.Comment_rating,
	}, nil
}

