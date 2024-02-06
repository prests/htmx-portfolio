package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/julienschmidt/httprouter"
	"github.com/prests/htmx-portfolio/components"
	"github.com/prests/htmx-portfolio/static"
	"github.com/prests/htmx-portfolio/views"
)

func main() {
	component := views.BaseHTML("Test", components.Hello("Shayne"))

	router := httprouter.New()

	httpFS := http.FileServer(http.FS(static.Files))
	router.Handler(http.MethodGet, "/static/*filepath", http.StripPrefix("/static/", httpFS))

	router.Handler(http.MethodGet, "/", templ.Handler(component))

	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	fmt.Println("Listening on :3000")
	server.ListenAndServe()
}
