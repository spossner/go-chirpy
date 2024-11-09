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

	// chirps
	mux.Handle("/api/chirps", chirps.HandleGetChirps(cfg))
	mux.Handle("/api/chirps/{id}", chirps.HandleGetChirpById(cfg))
	mux.Handle("POST /api/chirps", chirps.HandleCreateChirp(cfg))

	mux.Handle("GET /admin/metrics", metrics.HandleMetrics(cfg))
	mux.Handle("DELETE /admin/metrics", reset.HandleReset(cfg))
	mux.Handle("POST /admin/reset", reset.HandleReset(cfg))
}
