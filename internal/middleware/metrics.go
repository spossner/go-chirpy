package middleware

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"sync/atomic"
)

type ApiConfig struct {
	mu             *sync.Mutex
	hits           map[string]int
	fileserverHits atomic.Int32
}

func NewApiConfig() *ApiConfig {
	return &ApiConfig{
		mu:   &sync.Mutex{},
		hits: make(map[string]int),
	}
}

func (cfg *ApiConfig) track(url *url.URL) {
	parts := strings.Split(url.Path, "/")
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	// tracking root
	cfg.hits["/"] += 1

	key := ""
	for _, p := range parts {
		if p == "" {
			continue
		}
		// tracking additional parts of URL path
		key = key + "/" + p
		cfg.hits[key] += 1
	}
}

func (cfg *ApiConfig) get(part string) int {
	return cfg.hits[part]
}

func (cfg *ApiConfig) WithTracking(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.track(r.URL)
		next.ServeHTTP(w, r)
	})
}
func (cfg *ApiConfig) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost || r.Method == http.MethodDelete {
		cfg.fileserverHits.Store(0)
		w.WriteHeader(200)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(200)
	w.Write([]byte(fmt.Sprintf(`<html>
  <body>
    <h1>Welcome, Chirpy Admin</h1>
    <p>Chirpy has been visited %d times!</p>
	<pre>%s</pre>
  </body>
</html>`, cfg.get("/assets"), cfg.hits)))
}
