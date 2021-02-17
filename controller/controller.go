package controller

import (
	"github.com/arapov/soil/controller/assets"
	"github.com/arapov/soil/controller/debug"
	"github.com/arapov/soil/controller/home"
	"github.com/arapov/soil/controller/status"
)

// LoadRoutes loads controllers' routes
func LoadRoutes() {
	debug.Load() // TODO: acl
	status.Load()
	assets.Load()
	home.Load()
}
