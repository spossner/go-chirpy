-- name: CreateChirp :one
INSERT INTO chirps (body, user_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetChirpById :one
SELECT *
FROM chirps
WHERE id = $1;

-- name: GetChirps :many
SELECT *
from chirps
order by created_at;

-- name: GetChirpsByUserId :many
SELECT *
from chirps
where user_id = $1;