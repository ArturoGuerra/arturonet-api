package main

import (
	"github.com/arturoguerra/arturonet-api/internal/emailmanager"
	"github.com/arturoguerra/arturonet-api/internal/recaptcha"
	"github.com/arturoguerra/arturonet-api/internal/router"
	logging "github.com/arturoguerra/go-logging"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	logger := logging.New()

	r, err := router.NewDefault(logger)
	if err != nil {
		logger.Fatal(err)
	}

	rcptcha, err := recaptcha.NewDefault(logger)
	if err != nil {
		logger.Fatal(err)
	}

	emailGroup := r.Group("/api/email", middleware.Logger())
	emailManager, err := emailmanager.NewDefault(logger, rcptcha)
	if err != nil {
		logger.Fatal(err)
	}

	emailManager.Register(emailGroup)

}
