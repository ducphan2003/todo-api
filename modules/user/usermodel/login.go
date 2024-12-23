package usermodel

type UserLogin struct {
	Name     string `json:"name" form:"name"`
	Password string `json:"password" form:"password"`
}
