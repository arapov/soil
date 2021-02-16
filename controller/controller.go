package controller

import (
	"github.com/arapov/soil/controller/asset"
	"github.com/arapov/soil/controller/debug"
	"github.com/arapov/soil/controller/home"
)

// LoadRoutes loads controllers' routes
func LoadRoutes() {
	asset.Load()
	debug.Load()
	home.Load()
}
