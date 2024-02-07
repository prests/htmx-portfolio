package main

import (
	"os"

	config "github.com/prests/htmx-portfolio/internal/configs"
	"github.com/prests/htmx-portfolio/internal/pkg/logging"
	"github.com/prests/htmx-portfolio/internal/server/http"
)

func main() {
	cfg := config.New()
	logger := logging.NewLogger()

	server := http.NewService(cfg.HTTP(), logger)

	if err := server.Start(); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
