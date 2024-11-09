-- name: CreateRefreshToken :one
INSERT INTO refresh_tokens (token, user_id, expires_at)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetRefreshTokenByToken :one
SELECT *
FROM refresh_tokens
WHERE token = $1
AND revoked_at is null;

-- name: RevokeToken :exec
UPDATE refresh_tokens
SET revoked_at = current_timestamp
WHERE token = $1;