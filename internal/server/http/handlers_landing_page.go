package http

import (
	"net/http"

	"github.com/prests/htmx-portfolio/internal/server/http/web/templates"
)

func (h *Handlers) landingPageHandler(w http.ResponseWriter, r *http.Request) {
	props := &templates.LandingPageProps{
		Title: "Test",
		Name:  "Shayne",
	}

	tmpl := templates.LandingPage(props)
	tmpl.Render(r.Context(), w)
}
