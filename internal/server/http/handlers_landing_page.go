package http

import (
	"net/http"

	"github.com/prests/htmx-portfolio/internal/server/http/web/components"
	"github.com/prests/htmx-portfolio/internal/server/http/web/templates"
)

type landingPageProps struct {
	Title string
	Name  string
}

func (h *Handlers) landingPageHandler(w http.ResponseWriter, r *http.Request) {
	props := &landingPageProps{
		Title: "Test",
		Name:  "Shayne",
	}

	tmpl := templates.BaseHTML(props.Title, components.Hello(props.Name))
	tmpl.Render(r.Context(), w)
}
