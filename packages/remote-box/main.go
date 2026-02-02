package main

import (
	"os"

	"remote-box/config"
	"remote-box/controller"
)

func main() {
	config.LoadConfig("")

	ginPort := os.Getenv("GIN_PORT")
	if ginPort == "" {
		ginPort = ":9074"
	}

	controller.RouteController(ginPort)
}
