package config

import (
	"flag"
	"strconv"
)

type Config struct {
	Host string
	Port string
}

func NewConfig(args []string, getenv func(string) string) *Config {
	var port = flag.Int("port", 8080, "the port the server is listening at (defaults to 8080)")
	flag.Parse()
	return &Config{
		Host: "",
		Port: strconv.Itoa(*port),
	}
}
