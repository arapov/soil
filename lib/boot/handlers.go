package boot

import (
	"net/http"

	"github.com/arapov/soil/handlers/log"
	"github.com/arapov/soil/lib/core/router"
	"github.com/gorilla/context"
)

// RegisterHandlers sets up soil handlers
func RegisterHandlers() http.Handler {
	handler := router.GetInstance()
	return router.ChainHandler(
		handler,              // handler to wrap
		log.Handler,          // log request
		context.ClearHandler, // don't leak as gorilla.sessions documentation
	)
}
