package model

type User struct {
	ID       int32
	Name     string
	Email    string
	Password string
}
type UserInfo struct {
	Sex      string
	Name     string
	Email    string
	Birthday string
}

type UserSession struct {
	ID    int32
	Name  string
	Email string
}
