package controller

import (
	"github.com/arapov/soil/controller/asset"
	"github.com/arapov/soil/controller/debug"
	"github.com/arapov/soil/controller/home"
	"github.com/arapov/soil/controller/status"
)

// LoadRoutes loads controllers' routes
func LoadRoutes() {
	debug.Load() // TODO: acl
	status.Load()
	asset.Load()
	home.Load()
}
