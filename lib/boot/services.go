package boot

import (
	"github.com/arapov/soil/controller"
	"github.com/arapov/soil/lib/core/env"
	"github.com/arapov/soil/lib/flight"
)

// RegisterServices sets up soil components.
func RegisterServices(config *env.Info) {

	// Load controller routes.
	controller.LoadRoutes()

	// Store resource, to make it available in-flight.
	flight.Store(config)
}
