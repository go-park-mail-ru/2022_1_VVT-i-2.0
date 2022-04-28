package models

type mockRestaurant struct {
	Id   		int
	Name 		string
	City 		string
	Address 	string
	Image_path 	string
	Slug 		string
	Min_price 	int
	Avg_price 	int
	Rating     int
	Count_rating int
}

var Rest = &mockRestaurant{
	Id: 1,
	Name: "Name",
	City: "City",
	Address: "Address",
	Image_path: "imgPath",
	Slug: "slug",
	Min_price: 1,
	Avg_price: 1,
	Rating: 1,
	Count_rating: 1,
}

type mockDish struct {
	Id 			int
	Restaurant 	int
	Name 		string
	Description string
	Image_path  string
	Calories    int
	Price       int
	Weight		int
}

var Dish = &mockDish{
	Id: 1,
	Restaurant: 1,
	Name: "Name",
	Description: "Description",
	Image_path: "imgPath",
	Calories: 1,
	Price: 1,
	Weight: 1,
}

type mockCommentRestaurant struct {
	Id 			int
	Restaurant	int
	User_id		int
	Comment_text string
	Comment_rating int
}

var CommentRestaurant = &mockCommentRestaurant{
	Id: 1,
	Restaurant: 1,
	User_id: 1,
	Comment_text: "comment",
	Comment_rating: 5,
}

