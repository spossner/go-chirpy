package chirps

import (
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/spossner/go-chirpy/internal/auth"
	"github.com/spossner/go-chirpy/internal/config"
	"github.com/spossner/go-chirpy/internal/database"
	"github.com/spossner/go-chirpy/internal/utils"
	"net/http"
)

func HandleCreateChirp(cfg *config.ApiConfig) http.Handler {
	type request struct {
		Body string `json:"body"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, ok := utils.GetBearerToken(r)
		if !ok {
			fmt.Printf("no token found in %v\n", r.Header)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		userId, err := auth.CheckJWT(token, cfg.JWTSecret)
		if err != nil {
			fmt.Printf("invalid token: %w\n", err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		req, err := utils.Decode[request](r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := utils.ValidateChirp(req.Body); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		chirp, err := cfg.Queries.CreateChirp(r.Context(), database.CreateChirpParams{
			Body:   req.Body,
			UserID: userId,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		utils.EncodeWithStatus(w, http.StatusCreated, chirp)
	})
}

func HandleGetChirps(cfg *config.ApiConfig) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		chirps, err := cfg.Queries.GetChirps(r.Context())
		if err != nil {
			utils.EncodeWithError(w, http.StatusInternalServerError, "error fetching chirps", err)
			return
		}
		utils.Encode(w, chirps)
	})
}

func HandleGetChirpById(cfg *config.ApiConfig) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		var chirpID pgtype.UUID
		if err := chirpID.Scan(id); err != nil {
			utils.EncodeWithError(w, http.StatusBadRequest, fmt.Sprintf("invalid id %s", id), err)
			return
		}

		chirp, err := cfg.Queries.GetChirpById(r.Context(), chirpID)
		if err != nil {
			switch {
			case errors.Is(err, pgx.ErrNoRows):
				utils.EncodeWithError(w, http.StatusNotFound, fmt.Sprintf("chirp %s not found", id), err)
			default:
				utils.EncodeWithError(w, http.StatusInternalServerError, "error fetching chirp", err)
			}
			return
		}
		utils.Encode(w, chirp)
	})
}
