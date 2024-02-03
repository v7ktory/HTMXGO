// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package postgresdb

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createTodo = `-- name: CreateTodo :one
INSERT INTO todos (
  title, description, completed, created_at, user_id
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING id, title, description, completed, created_at, user_id
`

type CreateTodoParams struct {
	Title       string
	Description pgtype.Text
	Completed   pgtype.Bool
	CreatedAt   pgtype.Timestamp
	UserID      int32
}

func (q *Queries) CreateTodo(ctx context.Context, arg CreateTodoParams) (Todo, error) {
	row := q.db.QueryRow(ctx, createTodo,
		arg.Title,
		arg.Description,
		arg.Completed,
		arg.CreatedAt,
		arg.UserID,
	)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Completed,
		&i.CreatedAt,
		&i.UserID,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  name, email, password
) VALUES (
  $1, $2, $3
) RETURNING id, name, email, password
`

type CreateUserParams struct {
	Name     string
	Email    string
	Password string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser, arg.Name, arg.Email, arg.Password)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
	)
	return i, err
}

const deleteTodo = `-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = $1 AND user_id = $2
`

type DeleteTodoParams struct {
	ID     int32
	UserID int32
}

func (q *Queries) DeleteTodo(ctx context.Context, arg DeleteTodoParams) error {
	_, err := q.db.Exec(ctx, deleteTodo, arg.ID, arg.UserID)
	return err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const getTodo = `-- name: GetTodo :one
SELECT id, title, description, completed, created_at, user_id
FROM todos
WHERE user_id = $1 LIMIT 1
`

func (q *Queries) GetTodo(ctx context.Context, userID int32) (Todo, error) {
	row := q.db.QueryRow(ctx, getTodo, userID)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Completed,
		&i.CreatedAt,
		&i.UserID,
	)
	return i, err
}

const getTodos = `-- name: GetTodos :many
SELECT id, title, description, completed, created_at
FROM todos
WHERE user_id = $1
`

type GetTodosRow struct {
	ID          int32
	Title       string
	Description pgtype.Text
	Completed   pgtype.Bool
	CreatedAt   pgtype.Timestamp
}

func (q *Queries) GetTodos(ctx context.Context, userID int32) ([]GetTodosRow, error) {
	rows, err := q.db.Query(ctx, getTodos, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetTodosRow
	for rows.Next() {
		var i GetTodosRow
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.Completed,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUser = `-- name: GetUser :one
SELECT id, name, email, password FROM users
WHERE email = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUser, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
	)
	return i, err
}

const updateTodo = `-- name: UpdateTodo :exec
UPDATE todos
  SET title = $2, description = $3, completed = $4
WHERE id = $1
`

type UpdateTodoParams struct {
	ID          int32
	Title       string
	Description pgtype.Text
	Completed   pgtype.Bool
}

func (q *Queries) UpdateTodo(ctx context.Context, arg UpdateTodoParams) error {
	_, err := q.db.Exec(ctx, updateTodo,
		arg.ID,
		arg.Title,
		arg.Description,
		arg.Completed,
	)
	return err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users
  SET name = $2
WHERE id = $1
`

type UpdateUserParams struct {
	ID   int32
	Name string
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.Exec(ctx, updateUser, arg.ID, arg.Name)
	return err
}
