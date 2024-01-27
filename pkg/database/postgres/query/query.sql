-- name: GetUser :one
SELECT id, name, email, password FROM users
WHERE email = $1 LIMIT 1;

-- name: UpdateUser :exec
UPDATE users
  set name = $1
WHERE id = $1;

-- name: DeleteUsers :exec
DELETE FROM users
WHERE id = $1;

-- name: CreateUser :one
INSERT INTO users (
  name, email, password
) VALUES (
  $1, $2, $3
) RETURNING *;