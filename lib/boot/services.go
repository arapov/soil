package boot

import (
	"github.com/arapov/soil/controller"
	"github.com/arapov/soil/lib/core/env"
)

// RegisterServices sets up soil components
func RegisterServices(config *env.Info) {
	controller.LoadRoutes()
}
