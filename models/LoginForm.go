package models

type LoginForm struct {
	Email    string `form:"email" valid:"Required"`
	Password string `form:"password" valid:"Required; Min(8)"`
}
