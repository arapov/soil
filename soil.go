package main

import (
	"log"

	"github.com/arapov/soil/lib/boot"
	"github.com/arapov/soil/lib/core/env"
	"github.com/arapov/soil/lib/core/server"
)

func main() {
	config, err := env.LoadConfig("soil.yaml")
	if err != nil {
		log.Fatal(err)
	}

	handlers := boot.RegisterHandlers()

	server.Run(handlers, config.Server)

}
