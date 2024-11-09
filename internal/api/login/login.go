package login

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/spossner/go-chirpy/internal/api/user"
	"github.com/spossner/go-chirpy/internal/auth"
	"github.com/spossner/go-chirpy/internal/config"
	"github.com/spossner/go-chirpy/internal/database"
	"github.com/spossner/go-chirpy/internal/utils"
	"net/http"
	"time"
)

func HandleLogin(cfg *config.ApiConfig) http.Handler {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	type response struct {
		user.UserPresentation
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req, err := utils.Decode[request](r)
		if err != nil {
			utils.EncodeWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)
			return
		}
		usr, err := cfg.Queries.GetUserByEmail(r.Context(), req.Email)
		if err != nil {
			utils.EncodeWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)
			return
		}
		err = auth.CheckPasswordHash(req.Password, usr.HashedPassword)
		if err != nil {
			utils.EncodeWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)
			return
		}

		token, err := auth.CreateJWT(usr.ID, cfg.JWTSecret, time.Hour)
		if err != nil {
			utils.EncodeWithError(w, http.StatusInternalServerError, "could not signin user", err)
			return
		}
		refreshToken, err := auth.CreateRefreshToken()
		if err != nil {
			utils.EncodeWithError(w, http.StatusInternalServerError, "could not create refresh token", err)
			return
		}

		_, err = cfg.Queries.CreateRefreshToken(r.Context(), database.CreateRefreshTokenParams{
			Token:     refreshToken,
			UserID:    usr.ID,
			ExpiresAt: pgtype.Timestamp{Time: time.Now().Add(60 * 24 * time.Hour), Valid: true},
		})
		if err != nil {
			utils.EncodeWithError(w, http.StatusInternalServerError, "could not store refresh token", err)
		}

		utils.Encode(w, response{
			UserPresentation: user.NewUserPresentation(usr),
			Token:            token,
			RefreshToken:     refreshToken,
		})
	})
}

func HandleRefresh(cfg *config.ApiConfig) http.Handler {
	type response struct {
		Token string `json:"token"`
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, ok := utils.GetBearerToken(r)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		refreshToken, err := cfg.Queries.GetRefreshTokenByToken(r.Context(), token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		jwt, err := auth.CreateJWT(refreshToken.UserID, cfg.JWTSecret, time.Hour)
		if err != nil {
			utils.EncodeWithError(w, http.StatusInternalServerError, "error creating JWT token", err)
			return
		}
		utils.Encode(w, response{Token: jwt})
	})
}

func HandleRevoke(cfg *config.ApiConfig) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, ok := utils.GetBearerToken(r)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		err := cfg.Queries.RevokeToken(r.Context(), token)
		if err != nil {
			utils.EncodeWithError(w, http.StatusInternalServerError, "error revoking token", err)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	})
}
