package http

import (
	"crypto/tls"
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

type HTTP struct {
	server *http.Server
	cfg    *Config
	logger *slog.Logger
}

func (h *HTTP) Start() error {
	h.logger.Info("Starting http server", slog.String("host", h.cfg.Host), slog.Int("port", h.cfg.Port), slog.String("env", h.cfg.Env))
	return h.server.ListenAndServe()
}

type Config struct {
	Env  string
	Host string
	Port int
}

func NewService(cfg *Config, logger *slog.Logger) *HTTP {
	tlsConfig := &tls.Config{
		MinVersion:       tls.VersionTLS12,
		MaxVersion:       tls.VersionTLS12,
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		},
	}

	h := &Handlers{logger}

	server := &http.Server{
		Addr:           fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Handler:        h.routes(),
		ErrorLog:       slog.NewLogLogger(logger.Handler(), slog.LevelError),
		IdleTimeout:    time.Minute,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 524288, // 0.5 MB
	}
	server.TLSConfig = tlsConfig

	return &HTTP{
		server,
		cfg,
		logger,
	}
}
