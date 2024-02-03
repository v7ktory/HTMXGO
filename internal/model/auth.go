package model

type SignupReq struct {
	Name     string `form:"name"`
	Email    string `form:"email"`
	Password string `form:"password"`
}

type LoginReq struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}
