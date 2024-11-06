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

	mux.Handle("/app", http.StripPrefix("/app", apiCfg.WithMetricsInc(http.FileServer(http.Dir("./public")))))
	mux.Handle("/app/", http.StripPrefix("/app", apiCfg.WithMetricsInc(http.FileServer(http.Dir("./public")))))

	mux.Handle("/api/ping", ping.HandlePing())
	mux.Handle("/api/healthz", health.HandleHealthz())

	mux.Handle("POST /api/reset", apiCfg.WithResetMetrics())
	mux.Handle("GET /api/metrics", apiCfg)
}

func NewServer(cfg *config.Config) http.Handler {
	mux := http.NewServeMux()
	addRoutes(mux, cfg)
	var handler http.Handler = mux
	//handler = someMiddleware(handler)
	//handler = someMiddleware2(handler)
	//handler = someMiddleware3(handler)
	return handler
}
