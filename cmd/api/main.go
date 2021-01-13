package main

import (
	"fmt"

	"github.com/arturoguerra/arturonet-api/internal/emailmanager"
	"github.com/arturoguerra/arturonet-api/internal/emailsender"
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

	emailGroup := apiGroup.Group("/email")
	emailManager, err := emailmanager.NewDefault(logger, rcptcha, es)
	if err != nil {
		logger.Fatal(err)
	}

	emailManager.Register(emailGroup)

	logger.Infof("Running on %s:%s", rconfig.Host, rconfig.Port)
	r.Logger.Fatal(r.Start(fmt.Sprintf("%s:%s", rconfig.Host, rconfig.Port)))

}
