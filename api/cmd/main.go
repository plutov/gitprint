package main

import (
	"os"
	"path/filepath"

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

	reposDir := os.Getenv("GITHUB_REPOS_DIR")
	if entries, err := os.ReadDir(reposDir); err == nil {
		for _, entry := range entries {
			os.RemoveAll(filepath.Join(reposDir, entry.Name()))
		}
	}

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
