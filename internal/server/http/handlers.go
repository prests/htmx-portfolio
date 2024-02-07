package http

import (
	"log/slog"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"github.com/prests/htmx-portfolio/internal/server/http/web/static"
)

type Handlers struct {
	logger *slog.Logger
}

func (h *Handlers) routes() http.Handler {
	router := httprouter.New()

	httpFS := http.FileServer(http.FS(static.Files))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static/", httpFS))

	dynamic := alice.New(noSurf)
	router.Handler(http.MethodGet, "/", dynamic.ThenFunc(h.landingPageHandler))

	standard := alice.New()
	return standard.Then(router)
}