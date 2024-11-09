package reset

import (
	"github.com/spossner/go-chirpy/internal/config"
	"net/http"
)

func HandleReset(cfg *config.ApiConfig) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !cfg.Debug {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		cfg.Reset()
		cfg.Queries.DeleteUsers(r.Context())
		w.WriteHeader(http.StatusOK)
	})
}
