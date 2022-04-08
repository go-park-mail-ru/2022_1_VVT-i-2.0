package models

type UserId uint64

// // type User struct {
// }

type UserDataStorage struct {
	Id    UserId `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type UserAddDataStorage struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type UserDataResp struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type UserDataUpdateReq struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type LoginRequest struct {
	Phone string `json:"phone" valid:"phone, required"`
	Code  string `json:"code" valid:"required"`
}

type RegisterRequest struct {
	Phone string `json:"phone" valid:"phone, required"`
	Code  string `json:"code" valid:"required"`
	Name  string `json:"name" valid:"required"`
	Email string `json:"email" valid:"email,required"`
}

type SendCodeReq struct {
	Phone string
}
