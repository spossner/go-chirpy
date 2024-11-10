-- name: CreateUser :one
INSERT INTO users (email, hashed_password)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET email=$1, hashed_password=$2
WHERE id=$3
RETURNING *;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1;

-- name: GetUserById :one
SELECT *
FROM users
WHERE id = $1;


-- name: DeleteUsers :exec
DELETE FROM users;