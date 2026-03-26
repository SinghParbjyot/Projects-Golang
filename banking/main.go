package main

import (
	"github.com/singhparbjyot/banking/app"
	"github.com/singhparbjyot/banking/logger"
)

func main() {

	//Logs with the package zapper
	logger.Log.Info("Starting the application")
	app.Start()
}
