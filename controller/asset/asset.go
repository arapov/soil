package asset

import (
	"net/http"

	"github.com/arapov/soil/lib/core/router"
)

// Load routes.
func Load() {
	// TODO: figure out how to serve files in nested directories
	router.Get("/asset/{dir}/{file}", Index)
	// TODO: rewrite it properly
	router.Get("/favicon.ico", FavIcon)
}

// Index serves asset directory.
func Index(w http.ResponseWriter, r *http.Request) {
	// TODO: check for file and its existance
	http.ServeFile(w, r, r.URL.Path[1:])
}

// FavIcon serves favicon
func FavIcon(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "asset/favicon/favicon.ico")
}
