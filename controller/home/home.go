package home

import (
	"net/http"

	"github.com/arapov/soil/lib/core/router"
	"github.com/arapov/soil/lib/flight"
)

// Load routes.
func Load() {
	router.Get("/", Index)
}

// Index renders home page.
func Index(w http.ResponseWriter, r *http.Request) {
	f := flight.Get()
	v := f.View.New("home/index")

	v.Render(w)
}
