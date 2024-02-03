-- name: GetUser :one
SELECT id, name, email, password FROM users
WHERE email = $1 LIMIT 1;

-- name: UpdateUser :exec
UPDATE users
  SET name = $2
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: CreateUser :one
INSERT INTO users (
  name, email, password
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetTodo :one
SELECT id, title, description, completed, created_at, user_id
FROM todos
WHERE user_id = $1 LIMIT 1;

-- name: GetTodos :many
SELECT id, title, description, completed, created_at
FROM todos
WHERE user_id = $1;

-- name: CreateTodo :one
INSERT INTO todos (
  title, description, completed, created_at, user_id
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: UpdateTodo :exec
UPDATE todos
  SET completed = $3
WHERE id = $1 AND user_id = $2;

-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = $1 AND user_id = $2;

