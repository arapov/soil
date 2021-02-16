package flight

import (
	"github.com/arapov/soil/lib/core/env"
)

var (
	envInfo env.Info
)

// Store stuff.
func Store(config *env.Info) {
	envInfo = *config
}

// Get stuff.
func Get() env.Info {
	return envInfo
}
