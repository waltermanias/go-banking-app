package main

import (
	"banking-app/app"
	"banking-app/logger"
)

func main() {

	logger.Info("Starting the application")

	app.Start()

}
