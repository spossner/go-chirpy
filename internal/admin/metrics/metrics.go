package metrics

import (
	"fmt"
	"github.com/spossner/go-chirpy/internal/config"
	"net/http"
)

func WithTracking(cfg *config.ApiConfig, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.Track(r.URL)
		next.ServeHTTP(w, r)
	})
}

func HandleMetrics(cfg *config.ApiConfig) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(200)
		w.Write([]byte(fmt.Sprintf(`<html>
  <body>
    <h1>Welcome, Chirpy Admin</h1>
    <p>Chirpy has been visited %d times!</p>
	<pre>%s</pre>
  </body>
</html>`, cfg.Get("/assets"), cfg.Hits)))
	})
}
