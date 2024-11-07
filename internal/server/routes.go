package server

import (
	"github.com/spossner/go-chirpy/internal/api/health"
	"github.com/spossner/go-chirpy/internal/api/ping"
	"github.com/spossner/go-chirpy/internal/config"
	"github.com/spossner/go-chirpy/internal/middleware"
	"net/http"
)

func addRoutes(mux *http.ServeMux, cfg *config.Config) {
	apiCfg := middleware.NewApiConfig()

	mux.Handle("/app", http.StripPrefix("/app", apiCfg.WithTracking(http.FileServer(http.Dir("./public")))))
	mux.Handle("/app/", http.StripPrefix("/app", apiCfg.WithTracking(http.FileServer(http.Dir("./public")))))

	mux.Handle("/api/ping", ping.HandlePing())
	mux.Handle("/api/healthz", health.HandleHealthz())

	mux.Handle("GET /admin/metrics", apiCfg)
	mux.Handle("DELETE /admin/metrics", apiCfg)
	mux.Handle("POST /admin/reset", apiCfg)
}
