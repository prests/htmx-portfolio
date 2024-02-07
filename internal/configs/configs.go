package config

import (
	"flag"

	"github.com/prests/htmx-portfolio/internal/server/http"
)

type Configs struct {
	env  string
	port int
}

func (cfg *Configs) HTTP() *http.Config {
	return &http.Config{
		Env:  cfg.env,
		Port: cfg.port,
	}
}

func New() *Configs {
	var cfg Configs

	flag.IntVar(&cfg.port, "port", 8080, "Web server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")

	return &cfg
}
