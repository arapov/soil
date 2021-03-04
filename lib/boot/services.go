package boot

import (
	"github.com/arapov/soil/controller"
	"github.com/arapov/soil/lib/core/env"
	"github.com/arapov/soil/lib/core/view/extensions/link"
	"github.com/arapov/soil/lib/core/view/modifiers/uri"
	"github.com/arapov/soil/lib/flight"
)

// RegisterServices sets up soil components.
func RegisterServices(config *env.Info) {

	// Load controller routes.
	controller.LoadRoutes()

	config.View.Extensions(
		link.Map(config.View.BaseURI),
	)

	config.View.Modifiers(
		uri.Modify,
	)

	// Store resource, to make it available in-flight.
	flight.Store(config)
}
