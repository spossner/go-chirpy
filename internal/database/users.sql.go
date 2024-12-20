// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (email, hashed_password)
VALUES ($1, $2)
RETURNING id, email, created_at, updated_at, hashed_password, is_chirpy_red
`

type CreateUserParams struct {
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser, arg.Email, arg.HashedPassword)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.HashedPassword,
		&i.IsChirpyRed,
	)
	return i, err
}

const deleteUsers = `-- name: DeleteUsers :exec
DELETE FROM users
`

func (q *Queries) DeleteUsers(ctx context.Context) error {
	_, err := q.db.Exec(ctx, deleteUsers)
	return err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, email, created_at, updated_at, hashed_password, is_chirpy_red
FROM users
WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.HashedPassword,
		&i.IsChirpyRed,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT id, email, created_at, updated_at, hashed_password, is_chirpy_red
FROM users
WHERE id = $1
`

func (q *Queries) GetUserById(ctx context.Context, id pgtype.UUID) (User, error) {
	row := q.db.QueryRow(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.HashedPassword,
		&i.IsChirpyRed,
	)
	return i, err
}

const setRedUserById = `-- name: SetRedUserById :one
UPDATE users
SET is_chirpy_red=true
WHERE id=$1
RETURNING id, email, created_at, updated_at, hashed_password, is_chirpy_red
`

func (q *Queries) SetRedUserById(ctx context.Context, id pgtype.UUID) (User, error) {
	row := q.db.QueryRow(ctx, setRedUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.HashedPassword,
		&i.IsChirpyRed,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET email=$1, hashed_password=$2
WHERE id=$3
RETURNING id, email, created_at, updated_at, hashed_password, is_chirpy_red
`

type UpdateUserParams struct {
	Email          string      `json:"email"`
	HashedPassword string      `json:"hashed_password"`
	ID             pgtype.UUID `json:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUser, arg.Email, arg.HashedPassword, arg.ID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.HashedPassword,
		&i.IsChirpyRed,
	)
	return i, err
}
