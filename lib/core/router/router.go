package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

var (
	r *mux.Router
	// TODO: thread safe
)

// init subtly creates an instance.
func init() {
	r = mux.NewRouter()

}

// GetInstance returns router.
func GetInstance() *mux.Router {
	return r
}

// NotFound sets 404 error handler.
func NotFound(fn http.HandlerFunc) {
	r.NotFoundHandler = fn
}
