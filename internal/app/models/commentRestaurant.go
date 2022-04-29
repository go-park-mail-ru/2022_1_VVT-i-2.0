package models

type CommentRestaurantDataStorage struct {
	Id 			int
	Restaurant	int
	User_id		int
	Comment_text string
	Comment_rating int
}

type AddCommentRestaurantDataStorage struct {
	Restaurant	int
	User_id		int
	Comment_text string
	Comment_rating int
}

type AddCommentRestaurantUseCase struct {
	Restaurant	int
	User_id		int
	Comment_text string
	Comment_rating int
}

type AddCommentsRestaurantDataStorage struct {
	Comment []AddCommentRestaurantDataStorage
}

type CommentRestaurantId int64

//type CommentRestaurantUseCase struct {
//	Id 			int
//	Restaurant	int
//	User_id		int
//	Comment_text string
//	Comment_rating int
//}

type CommentRestaurantUseCase struct {
	Id 			int `json:"id"`
	Restaurant	int `json:"restaurants"`
	User_id		int `json:"userId"`
	Comment_text string `json:"commentText"`
	Comment_rating int `json:"commentRating"`
}

type CommentsRestaurantUseCase struct {
	Comment []CommentRestaurantUseCase
}

type CommentDataDelivery struct {
	Id 			int `json:"id"`
	Restaurant	int `json:"restaurants"`
	User_id		int `json:"userId"`
	Comment_text string `json:"commentText"`
	Comment_rating int `json:"commentRating"`
}

type CommentsDataDelivery struct {
	Comment []CommentDataDelivery `json:"comment"`
}

type AddCommentRestaurant struct {
	Restaurant		int		`json:"restaurants"`
	User_id			int		`json:"userId"`
	Comment_text 	string	`json:"commentText"`
	Comment_rating 	int		`json:"commentRating"`
}

type Comment struct {
	Id 			int
	Restaurant	int
	User_id		int
	Comment_text string
	Comment_rating int
}

