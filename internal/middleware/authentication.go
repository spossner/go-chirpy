package middleware

import (
	"fmt"
	"github.com/spossner/go-chirpy/internal/auth"
	"github.com/spossner/go-chirpy/internal/config"
	"github.com/spossner/go-chirpy/internal/database"
	"github.com/spossner/go-chirpy/internal/utils"
	"net/http"
)

type AuthenticatedHandler func(*config.ApiConfig, database.User, http.ResponseWriter, *http.Request)

func WithAuthentication(cfg *config.ApiConfig, handlerFunc func(apiConfig *config.ApiConfig, usr database.User) http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, ok := utils.GetBearerToken(r)
		if !ok {
			fmt.Printf("no token found in %v\n", r.Header)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		userId, err := auth.CheckJWT(token, cfg.JWTSecret)
		if err != nil {
			fmt.Println("invalid token:", err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		usr, err := cfg.Queries.GetUserById(r.Context(), userId)
		if err != nil {
			fmt.Println("unknown user:", err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		handlerFunc(cfg, usr).ServeHTTP(w, r)
	})
}
