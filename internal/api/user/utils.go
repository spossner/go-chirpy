package user

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/spossner/go-chirpy/internal/database"
)

type UserPresentation struct {
	ID          pgtype.UUID      `json:"id"`
	Email       string           `json:"email"`
	IsChirpyRed pgtype.Bool      `json:"is_chirpy_red"`
	CreatedAt   pgtype.Timestamp `json:"created_at"`
	UpdatedAt   pgtype.Timestamp `json:"updated_at"`
}

func NewUserPresentation(user database.User) UserPresentation {
	return UserPresentation{
		ID:          user.ID,
		Email:       user.Email,
		IsChirpyRed: user.IsChirpyRed,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}
}
