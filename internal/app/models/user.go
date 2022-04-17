package models

type UserId uint64

type UserDataStorage struct {
	Id    UserId
	Name  string
	Phone string
	Email string
}

type UserAddDataStorage struct {
	Name  string
	Phone string
	Email string
}

type UserDataUsecase struct {
	Id    UserId
	Name  string
	Phone string
	Email string
}

type UserDataResp struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type LoginReq struct {
	Phone string `json:"phone" valid:"phone, required"`
	Code  string `json:"code" valid:"required"`
}

type RegisterReq struct {
	Phone string `json:"phone" valid:"phone, required"`
	Code  string `json:"code" valid:"required"`
	Name  string `json:"name" valid:"name,required"`
	Email string `json:"email" valid:"email,required"`
}

// type RegisterReq struct {
// Phone string `json:"phone" valid:"phone, required"`
// Code  string `json:"code" valid:"required"`
// Name  string `json:"name" valid:"name,required"`
// 	Email string `json:"email" valid:"email,required"`
// }

type SendCodeReq struct {
	Phone string `json:"phone" valid:"phone, required"`
}

type SendCodeResp struct {
	IsRegistered bool `json:"registered"`
}

type UpdateUserReq struct {
	Name  string `json:"name" valid:"name"`
	Email string `json:"email" valid:"email"`
}

type UpdateUser struct {
	Id    UserId
	Name  string
	Email string
}

// validate name
// экранировать код
