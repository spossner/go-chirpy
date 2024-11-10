-- name: CreateChirp :one
INSERT INTO chirps (body, user_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetChirpById :one
SELECT *
FROM chirps
WHERE id = $1;

-- name: DeleteChirpById :exec
DELETE
FROM chirps
WHERE id = $1;

-- name: GetChirps :many
SELECT *
FROM chirps
ORDER BY created_at;

-- name: GetChirpsByUserId :many
SELECT *
FROM chirps
WHERE user_id = $1;