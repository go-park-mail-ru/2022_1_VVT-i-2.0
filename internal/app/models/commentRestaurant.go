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

// Repository models

type CommentRestaurantDataStorage struct {
	RestaurantId	int		`db:"restaurant_id"`
	Author        	string	`db:"author"`
	Text          	string	`db:"text"`
	Stars         	int		`db:"stars"`
	Date          	string	`db:"get_ru_date"`
}

type CommentsRestaurantDataStorage struct {
	Comments []CommentRestaurantDataStorage
}

//// UseCase

// UseCase Request

type GetRestaurantCommentsUcaseReq struct {
	Slug string
}

// UseCase models

type AddCommentRestaurantDataStorage struct {
	RestaurantId	int		`db:"restaurant_id"`
	User           	string	`db:"author"`
	CommentText   	string	`db:"text"`
	CommentRating 	int		`db:"stars"`
}

type AddCommentRestaurantUseCase struct {
	Slug           string
	CommentText   string
	CommentRating int
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
	RestaurantId	int    `json:"restaurants_id"`
	Author       	string `json:"author"`
	Text          	string `json:"text"`
	Stars         	int    `json:"starts"`
	Date          	string `json:"date"`
	//Restaurant	int `json:"restaurants"`
	//User		string `json:"user"`
	//Comment_text string `json:"commentText"`
	//Comment_rating int `json:"commentRating"`
}

type CommentsRestaurantUseCase struct {
	Comment []CommentRestaurantUseCase
}

type GetCommentDataDelivery struct {
	Author string `json:"author"`
	Text   string `json:"text"`
	Stars  int    `json:"stars"`
	Date   string `json:"date"`
}

type GetCommentsDataDelivery struct {
	Comment []GetCommentDataDelivery `json:"comment"`
}

type CommentDataDelivery struct {
	RestaurantId	int    `json:"restaurants_id"`
	Author       	 string `json:"author"`
	Text         	 string `json:"text"`
	Stars        	 int    `json:"stars"`
	Date         	 string `json:"date"`
}

type CommentsDataDelivery struct {
	Comment []CommentDataDelivery `json:"comment"`
}

type AddCommentRestaurant struct {
	Slug     		string	`json:"slug"`
	CommentText  	string 	`json:"text"`
	CommentRating	int    	`json:"stars"`
}

type Comment struct {
	Id             int
	Restaurant     int
	UserId        int
	CommentText   string
	CommentRating int
}

//create table comments
//(
//author varchar(50) NOT NULL,
//restaurant_id integer REFERENCES restaurants NOT NULL,
//text varchar(1024) NOT NULL,
//stars integer NOT NULL CHECK(1 <= stars and stars <= 5),
//date TIMESTAMP DEFAULT now() NOT NULL
//);
