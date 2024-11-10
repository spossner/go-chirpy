package server

import (
	"github.com/spossner/go-chirpy/internal/admin/metrics"
	"github.com/spossner/go-chirpy/internal/admin/reset"
	"github.com/spossner/go-chirpy/internal/api/chirps"
	"github.com/spossner/go-chirpy/internal/api/health"
	"github.com/spossner/go-chirpy/internal/api/login"
	"github.com/spossner/go-chirpy/internal/api/ping"
	"github.com/spossner/go-chirpy/internal/api/user"
	"github.com/spossner/go-chirpy/internal/config"
	"github.com/spossner/go-chirpy/internal/middleware"
	"net/http"
)

func addRoutes(mux *http.ServeMux, cfg *config.ApiConfig) {
	mux.Handle("/app", http.StripPrefix("/app", metrics.WithTracking(cfg, http.FileServer(http.Dir("./public")))))
	mux.Handle("/app/", http.StripPrefix("/app", metrics.WithTracking(cfg, http.FileServer(http.Dir("./public")))))

	mux.Handle("/api/ping", ping.HandlePing())
	mux.Handle("/api/healthz", health.HandleHealthz())

	// auth
	mux.Handle("POST /api/login", login.HandleLogin(cfg))
	mux.Handle("POST /api/refresh", login.HandleRefresh(cfg))
	mux.Handle("POST /api/revoke", login.HandleRevoke(cfg))

	// users
	mux.Handle("POST /api/users", user.HandleCreateUser(cfg))
	mux.Handle("PUT /api/users", middleware.WithAuthentication(cfg, user.HandleUpdateUser))

	// chirps
	mux.Handle("/api/chirps", chirps.HandleGetChirps(cfg))
	mux.Handle("/api/chirps/{id}", chirps.HandleGetChirpById(cfg))
	mux.Handle("POST /api/chirps", middleware.WithAuthentication(cfg, chirps.HandleCreateChirp))
	mux.Handle("DELETE /api/chirps/{id}", middleware.WithAuthentication(cfg, chirps.HandleDeleteChirp))

	mux.Handle("GET /admin/metrics", metrics.HandleMetrics(cfg))
	mux.Handle("DELETE /admin/metrics", reset.HandleReset(cfg))
	mux.Handle("POST /admin/reset", reset.HandleReset(cfg))
}
