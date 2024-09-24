package main

import (
	"os"

	controllers "github.com/plutov/gitprint/api/pkg/controllers"
	"github.com/plutov/gitprint/api/pkg/log"
	"github.com/plutov/gitprint/api/pkg/services"
)

func main() {
	log.Named("api")
	logLevel := "info"
	if os.Getenv("LOG_LEVEL") != "" {
		logLevel = os.Getenv("LOG_LEVEL")
	}
	log.SetLogLevel(logLevel)

	svc, err := services.InitServices()
	if err != nil {
		log.WithError(err).Fatal("unable to init dependencies")
	}

	handler := controllers.NewHandler(svc)
	if err != nil {
		log.WithError(err).Fatal("unable to start server")
	}

	r := controllers.NewRouter(handler)

	if err := r.Start(":8080"); err != nil {
		log.WithError(err).Fatal("shutting down the server")
	}
}
