package model

type UserInfo struct {
	Sex      string `json:"sex"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Birthday string `json:"birthday"`
}

type UserSession struct {
	ID    int32  `form:"id"`
	Name  string `form:"name"`
	Email string `form:"email"`
}

type SignupReq struct {
	Name     string `form:"name"`
	Email    string `form:"email"`
	Password string `form:"password"`
}

type LoginReq struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}
