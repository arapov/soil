package router

import (
	"net/http"

	"github.com/justinas/alice"
)

// ChainHandler returns a handler of chained middleware.
func ChainHandler(h http.Handler, c ...alice.Constructor) http.Handler {
	return alice.New(c...).Then(h)
}

// Get wraps router.HandleFunc(path, handle).Methods("GET")
func Get(path string, fn http.HandlerFunc, c ...alice.Constructor) {
	r.HandleFunc(path,
		alice.New(c...).ThenFunc(fn).(http.HandlerFunc),
	).Methods("GET")
}

// GetSub wraps router.PathPrefix(path).Handler(handle).Methods("GET")
func GetSub(path string, fn http.HandlerFunc, c ...alice.Constructor) {
	r.PathPrefix(path).Handler(
		alice.New(c...).ThenFunc(fn).(http.HandlerFunc),
	).Methods("GET")
}
