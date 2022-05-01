package models

type mockRestaurant struct {
	Id                   int
	Name                 string
	ImagePath            string
	Slug                 string
	MinPrice             int
	AggRating            int
	ReviewCount          int
	UpMinutsToDelivery   int
	DownMinutsToDelivery int
}

var Rest = &mockRestaurant{
	Id:                   1,
	Name:                 "Name",
	ImagePath:            "imgPath",
	Slug:                 "slug",
	MinPrice:             1,
	AggRating:            1,
	ReviewCount:          1,
	UpMinutsToDelivery:   1,
	DownMinutsToDelivery: 1,
}

type mockDish struct {
	Id           int
	RestaurantId int
	Name         string
	Description  string
	ImagePath    string
	Calories     int
	Price        int
	Weight       int
}

var Dish = &mockDish{
	Id:           1,
	RestaurantId: 1,
	Name:         "Name",
	Description:  "Description",
	ImagePath:    "imgPath",
	Calories:     1,
	Price:        1,
	Weight:       1,
}

type mockCommentRestaurant struct {
	Id             int
	Restaurant     int
	User_id        int
	Comment_text   string
	Comment_rating int
}

var CommentRestaurant = &mockCommentRestaurant{
	Id:             1,
	Restaurant:     1,
	User_id:        1,
	Comment_text:   "comment",
	Comment_rating: 5,
}
