package debug

import (
	"net/http"
	"net/http/pprof"

	"github.com/gorilla/mux"

	"github.com/arapov/soil/lib/core/router"
)

// Load routes.
func Load() {
	router.Get("/debug/pprof/", Index)
	router.Get("/debug/pprof/{pprof}", Profile)
}

// Index displays pprof page.
func Index(w http.ResponseWriter, r *http.Request) {
	pprof.Index(w, r)
}

// Profile shows specific profile.
func Profile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)["pprof"]
	switch vars {
	case "cmdline":
		pprof.Cmdline(w, r)
	case "profile":
		pprof.Profile(w, r)
	case "symbol":
		pprof.Symbol(w, r)
	case "trace":
		pprof.Trace(w, r)
	default:
		Index(w, r)
	}
}
