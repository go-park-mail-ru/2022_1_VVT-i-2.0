package models

type CommentRestaurantDataStorage struct {
	Restaurant_id	int
	Author			string
	Text 			string
	Stars			int
	Date 			string	`db:"get_ru_date"`
}

type AddCommentRestaurantDataStorage struct {
	Restaurant_id	int
	User		string
	Comment_text string
	Comment_rating int
}

type AddCommentRestaurantUseCase struct {
	Restaurant	int
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
	Restaurant_id	int		`json:"restaurants_id"`
	Author			string	`json:"author"`
	Text 			string	`json:"text"`
	Stars			int		`json:"starts"`
	Date 			string	`json:"date"`
	//Restaurant	int `json:"restaurants"`
	//User		string `json:"user"`
	//Comment_text string `json:"commentText"`
	//Comment_rating int `json:"commentRating"`
}

type CommentsRestaurantUseCase struct {
	Comment []CommentRestaurantUseCase
}

type GetCommentDataDelivery struct {
	Author			string	`json:"author"`
	Text 			string	`json:"text"`
	Stars			int		`json:"stars"`
	Date 			string	`json:"date"`
}

type GetCommentsDataDelivery struct {
	Comment []GetCommentDataDelivery `json:"comment"`
}

type CommentDataDelivery struct {
	Restaurant_id	int		`json:"restaurants_id"`
	Author			string	`json:"author"`
	Text 			string	`json:"text"`
	Stars			int		`json:"stars"`
	Date 			string	`json:"date"`
}

type CommentsDataDelivery struct {
	Comment []CommentDataDelivery `json:"comment"`
}

type AddCommentRestaurant struct {
	Restaurant		int		`json:"restaurants_id"`
	Comment_text 	string	`json:"text"`
	Comment_rating 	int		`json:"stars"`
}

type Comment struct {
	Id 			int
	Restaurant	int
	User_id		int
	Comment_text string
	Comment_rating int
}

//create table comments
//(
//author varchar(50) NOT NULL,
//restaurant_id integer REFERENCES restaurants NOT NULL,
//text varchar(1024) NOT NULL,
//stars integer NOT NULL CHECK(1 <= stars and stars <= 5),
//date TIMESTAMP DEFAULT now() NOT NULL
//);

