package router

import (
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
