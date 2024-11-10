package chirps

import (
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/spossner/go-chirpy/internal/config"
	"github.com/spossner/go-chirpy/internal/database"
	"github.com/spossner/go-chirpy/internal/utils"
	"net/http"
)

func HandleCreateChirp(cfg *config.ApiConfig, usr database.User) http.Handler {
	type request struct {
		Body string `json:"body"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
			Body:   utils.CleanChirp(req.Body),
			UserID: usr.ID,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		utils.EncodeWithStatus(w, http.StatusCreated, chirp)
	})
}

func HandleDeleteChirp(cfg *config.ApiConfig, usr database.User) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, ok := utils.ParseUUID(r.PathValue("id"))
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		chirp, err := cfg.Queries.GetChirpById(r.Context(), id)
		if err != nil {
			fmt.Println("error fetching chirp:", err)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if chirp.UserID != usr.ID {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		err = cfg.Queries.DeleteChirpById(r.Context(), chirp.ID)
		if err != nil {
			fmt.Println("error deleting chirp:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	})
}

func HandleGetChirps(cfg *config.ApiConfig) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var chirps []database.Chirp
		var err error
		authorId := r.URL.Query().Get("author_id")
		if authorId != "" {
			if uid, ok := utils.ParseUUID(authorId); ok {
				chirps, err = cfg.Queries.GetChirpsByUserId(r.Context(), uid)
			}
		} else {
			chirps, err = cfg.Queries.GetChirps(r.Context())
		}
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
