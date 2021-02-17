package assets

import (
	"net/http"
	"os"

	"github.com/arapov/soil/controller/status"

	"github.com/arapov/soil/lib/core/router"
)

// Load routes to serve static content.
func Load() {
	router.GetSub("/assets", Index)
}

// Index serves assets directory.
func Index(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]

	if fi, err := os.Stat(path); err == nil && !fi.IsDir() {
		http.ServeFile(w, r, path)
		return
	}

	status.Error404(w, r)
}
