package main

import (
	"github.com/leslesnoa/go-microservices-demo/app"
	"github.com/leslesnoa/go-microservices-demo/logger"
)

func main() {

	// log.Println("Starting application...")
	logger.Info("Starting the application...")
	app.Start()
}
