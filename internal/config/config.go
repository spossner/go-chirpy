package config

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/spossner/go-chirpy/internal/database"
	"github.com/spossner/go-chirpy/internal/utils"
	"log"
	"net/url"
	"strings"
	"sync"
)

type ApiConfig struct {
	mu        *sync.RWMutex
	Queries   *database.Queries
	Debug     bool
	Host      string
	Port      string
	JWTSecret string
	PolkaKey  string
	Hits      map[string]int
}

func NewApiConfig() *ApiConfig {
	conn, err := pgx.Connect(context.Background(), utils.MustGetEnvString("DB_URL"))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	return &ApiConfig{
		Queries:   database.New(conn),
		Host:      utils.GetEnvString("HOST", ""),
		Port:      utils.GetEnvString("PORT", "8080"),
		JWTSecret: utils.MustGetEnvString("JWT_SECRET"),
		PolkaKey:  utils.MustGetEnvString("POLKA_KEY"),
		Debug:     utils.GetEnvBool("DEBUG", false),
		mu:        &sync.RWMutex{},
		Hits:      make(map[string]int),
	}
}

func (cfg *ApiConfig) Track(url *url.URL) {
	parts := strings.Split(url.Path, "/")
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	// tracking root
	cfg.Hits["/"] += 1

	key := ""
	for _, p := range parts {
		if p == "" {
			continue
		}
		// tracking additional parts of URL path
		key = key + "/" + p
		cfg.Hits[key] += 1
	}
}

func (cfg *ApiConfig) Get(part string) int {
	cfg.mu.RLock()
	defer cfg.mu.RUnlock()
	return cfg.Hits[part]
}

func (cfg *ApiConfig) Reset() {
	clear(cfg.Hits)
}
