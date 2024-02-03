package model

type User struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
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
