package main

import (
	"fmt"

	"github.com/arturoguerra/arturonet-api/internal/emailmanager"
	"github.com/arturoguerra/arturonet-api/internal/emailsender"
	"github.com/arturoguerra/arturonet-api/internal/projects"
	"github.com/arturoguerra/arturonet-api/internal/recaptcha"
	"github.com/arturoguerra/arturonet-api/internal/router"
	logging "github.com/arturoguerra/go-logging"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// .env for development purposes
	godotenv.Load()

	logger := logging.New()

	r, rconfig, err := router.NewDefault(logger)
	if err != nil {
		logger.Fatal(err)
	}

	rcptcha, err := recaptcha.NewDefault(logger)
	if err != nil {
		logger.Fatal(err)
	}

	es, err := emailsender.NewDefault(logger)
	if err != nil {
		logger.Fatal(err)
	}

	apiGroup := r.Group("/api", middleware.Logger())

	if err := emailmanager.NewDefault(logger, rcptcha, es, apiGroup.Group("/email")); err != nil {
		logger.Fatal(err)
	}

	if err := projects.New(logger, apiGroup.Group("/projects")); err != nil {
		logger.Fatal(err)
	}

	logger.Infof("Running on %s:%s", rconfig.Host, rconfig.Port)
	r.Logger.Fatal(r.Start(fmt.Sprintf("%s:%s", rconfig.Host, rconfig.Port)))

}
