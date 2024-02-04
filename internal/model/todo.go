package model

import "time"

type Todo struct {
	ID          int32
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
	UserID      int32
}
type TodoReq struct {
	Title       string `form:"title"`
	Description string `form:"description"`
}

type TodoInfo struct {
	ID        int32
	Title     string
	Completed bool
}
