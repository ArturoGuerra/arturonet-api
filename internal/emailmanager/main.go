package emailmanager

import (
	"github.com/arturoguerra/arturonet-api/internal/emailsender"
	"github.com/arturoguerra/arturonet-api/internal/recaptcha"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type (
	emailManager struct {
		Config    *Config
		Recaptcha recaptcha.Recaptcha
		Logger    *logrus.Logger
		Sender    emailsender.EmailSender
	}
)

// NewDefault returns an instance of emailmanager with an env config
func NewDefault(logger *logrus.Logger, rcptch recaptcha.Recaptcha, es emailsender.EmailSender, g *echo.Group) error {
	cfg, err := loadConfig()
	if err != nil {
		return err
	}

	return New(logger, rcptch, es, cfg, g)
}

// New returns a new instance of emailmanager
func New(logger *logrus.Logger, rcptch recaptcha.Recaptcha, es emailsender.EmailSender, config *Config, g *echo.Group) error {
	em := &emailManager{
		Config:    config,
		Recaptcha: rcptch,
		Logger:    logger,
		Sender:    es,
	}

	g.POST("", em.getEmailHandler())

	return nil
}
