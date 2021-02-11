package router

import (
	"net/http"

	"github.com/justinas/alice"
)

// ChainHandler returns a handler of chained middleware.
func ChainHandler(h http.Handler, c ...alice.Constructor) http.Handler {
	return alice.New(c...).Then(h)
}
