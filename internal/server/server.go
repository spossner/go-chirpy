package server

import (
	"github.com/spossner/go-chirpy/internal/config"
	"net/http"
)

func NewServer(cfg *config.ApiConfig) http.Handler {
	mux := http.NewServeMux()
	addRoutes(mux, cfg)
	var handler http.Handler = mux
	//handler = someMiddleware(handler)
	//handler = someMiddleware2(handler)
	//handler = someMiddleware3(handler)
	return handler
}
