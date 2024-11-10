package user

import (
	"github.com/spossner/go-chirpy/internal/auth"
	"github.com/spossner/go-chirpy/internal/config"
	"github.com/spossner/go-chirpy/internal/database"
	"github.com/spossner/go-chirpy/internal/utils"
	"net/http"
)

type request struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func HandleCreateUser(cfg *config.ApiConfig) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req, err := utils.Decode[request](r)
		if err != nil {
			utils.EncodeWithError(w, http.StatusBadRequest, "invalid request", err)
			return
		}

		hash, err := auth.HashPassword(req.Password)
		if err != nil {
			utils.EncodeWithError(w, http.StatusBadRequest, "invalid password", err)
			return
		}

		usr, err := cfg.Queries.CreateUser(r.Context(), database.CreateUserParams{
			Email:          req.Email,
			HashedPassword: hash,
		})
		if err != nil {
			utils.EncodeWithError(w, http.StatusInternalServerError, "could not create new user", err)
			return
		}
		utils.EncodeWithStatus(w, http.StatusCreated, NewUserPresentation(usr))
	})

}

func HandleUpdateUser(cfg *config.ApiConfig, usr database.User) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req, err := utils.Decode[request](r)
		if err != nil {
			utils.EncodeWithError(w, http.StatusBadRequest, "invalid request", err)
			return
		}

		hash, err := auth.HashPassword(req.Password)
		if err != nil {
			utils.EncodeWithError(w, http.StatusBadRequest, "invalid password", err)
			return
		}

		usr, err := cfg.Queries.UpdateUser(r.Context(), database.UpdateUserParams{
			ID:             usr.ID,
			Email:          req.Email,
			HashedPassword: hash,
		})
		if err != nil {
			utils.EncodeWithError(w, http.StatusInternalServerError, "could not update user", err)
			return
		}
		utils.EncodeWithStatus(w, http.StatusOK, NewUserPresentation(usr))
	})

}
