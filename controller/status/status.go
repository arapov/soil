package status

import (
	"net/http"

	"github.com/arapov/soil/lib/flight"
)

// Load routes.
func Load() {
	// TODO: Routes TBD
}

// Error404 - page not found
func Error404(w http.ResponseWriter, r *http.Request) {
	f := flight.Get()
	w.WriteHeader(http.StatusNotFound)
	v := f.View.New("status/index")
	v.Render(w)
}
