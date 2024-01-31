package model

import "time"

type TodoReq struct {
	Title       string `form:"title"`
	Description string `form:"description"`
}

type Todo struct {
	ID          int32     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	UserID      int32     `json:"user_id"`
}
