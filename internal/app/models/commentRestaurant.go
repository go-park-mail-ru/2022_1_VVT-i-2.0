package models

////Repository

// Repository Request

type GetRestaurantByIdRepoReq struct {
	Id Id
}

//type GetRestaurantBySlugRepoReq struct {
//	Slug string
//}

type GetRestaurantCommentsRepoReq struct {
	Id Id
}

type UpdateRestaurantRatingRepoReq struct {
	RestId        int
	NewRestRating int
	CountRating   int
}

type AddRestaurantCommentRepoReq struct {
	RestaurantId  int    `db:"restaurant_id"`
	User          string `db:"author"`
	CommentText   string `db:"text"`
	CommentRating int    `db:"stars"`
}

// Repository models

type CommentRestaurantDataStorage struct {
	RestaurantId int    `db:"restaurant_id"`
	Author       string `db:"author"`
	Text         string `db:"text"`
	Stars        int    `db:"stars"`
	Date         string `db:"get_ru_date"`
}

type CommentsRestaurantDataStorage struct {
	Comments []CommentRestaurantDataStorage
}

//// UseCase

// UseCase Request

type GetRestaurantCommentsUcaseReq struct {
	Slug string
}

type AddCommentRestaurantUcaseReq struct {
	UserId        UserId
	Slug          string
	CommentText   string
	CommentRating int
}

// UseCase models

type CommentRestaurantUseCase struct {
	RestaurantId int    `json:"restaurants_id"`
	Author       string `json:"author"`
	Text         string `json:"text"`
	Stars        int    `json:"starts"`
	Date         string `json:"date"`
}

type CommentsRestaurantUseCase struct {
	Comment []CommentRestaurantUseCase
}

//// Handler

// handler request

type AddCommentRestaurantReq struct {
	Slug          string `json:"slug"`
	CommentText   string `json:"text"`
	CommentRating int    `json:"stars"`
}

// handler models

type GetCommentDataDelivery struct {
	Author string `json:"author"`
	Text   string `json:"text" valid:"comment"`
	Stars  int    `json:"stars"`
	Date   string `json:"date"`
}

type GetCommentsDataDelivery struct {
	Comment []GetCommentDataDelivery `json:"comment"`
}

type CommentDataDelivery struct {
	RestaurantId int    `json:"restaurants_id"`
	Author       string `json:"author"`
	Text         string `json:"text"`
	Stars        int    `json:"stars"`
	Date         string `json:"date"`
}
