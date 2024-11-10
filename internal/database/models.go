// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Chirp struct {
	ID        pgtype.UUID      `json:"id"`
	Body      string           `json:"body"`
	UserID    pgtype.UUID      `json:"user_id"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}

type RefreshToken struct {
	Token     string           `json:"token"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
	UserID    pgtype.UUID      `json:"user_id"`
	ExpiresAt pgtype.Timestamp `json:"expires_at"`
	RevokedAt pgtype.Timestamp `json:"revoked_at"`
}

type User struct {
	ID             pgtype.UUID      `json:"id"`
	Email          string           `json:"email"`
	CreatedAt      pgtype.Timestamp `json:"created_at"`
	UpdatedAt      pgtype.Timestamp `json:"updated_at"`
	HashedPassword string           `json:"hashed_password"`
	IsChirpyRed    pgtype.Bool      `json:"is_chirpy_red"`
}
