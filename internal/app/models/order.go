package models

import "database/sql"

type OrderPosition struct {
	Id    int64 `json:"id,string"`
	Count int64 `json:"count,string"`
}

type OrderRepo struct {
	OrderId int64
	User_id UserId
	Address string
	Comment sql.NullString
	Cost    string
	Cart    []OrderPosition
}

type OrderRepoInput struct {
	User_id UserId
	Address string
	Comment string
	Cart    []OrderPosition
}

// type OrderRepoInput struct {
// 	UserId  int64
// 	Address string
// 	Comment string
// 	Cart    []OrderPosition
// }

type OrderRepoAnsw struct {
	OrderId int64 `json:"id"`
}

type OrderRepoResp struct {
	OrderId int64 `json:"id"`
}

type OrderReq struct {
	Address string          `json:"address" valid:"address,required"`
	Comment string          `json:"comment" valid:"comment"` //TODO валидатор сделать
	Cart    []OrderPosition `json:"cart" valid:"required"`
}

type OrderUcaseInput struct {
	UserId  UserId
	Address string
	Comment string
	Cart    []OrderPosition
}

type OrderUcaseAnsw struct {
	OrderId int64
}
