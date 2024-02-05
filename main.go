package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/prests/htmx-portfolio/components"
)

func main() {
	component := components.Hello("Shayne")

	http.Handle("/", templ.Handler(component))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
