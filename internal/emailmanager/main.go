package emailmanager

import (
	"github.com/arturoguerra/arturonet-api/internal/emailsender"
	"github.com/arturoguerra/arturonet-api/internal/recaptcha"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type (
	// EmailManager exports all the email functions
	EmailManager interface {
		Register(group *echo.Group)
	}

	emailManager struct {
		Config    *Config
		Recaptcha recaptcha.Recaptcha
		Logger    *logrus.Logger
		Sender    emailsender.EmailSender
	}
)

// NewDefault returns an instance of emailmanager with an env config
func NewDefault(logger *logrus.Logger, rcptch recaptcha.Recaptcha, es emailsender.EmailSender) (EmailManager, error) {
	cfg, err := loadConfig()
	if err != nil {
		return nil, err
	}

	return New(logger, rcptch, es, cfg)
}

// New returns a new instance of emailmanager
func New(logger *logrus.Logger, rcptch recaptcha.Recaptcha, es emailsender.EmailSender, config *Config) (EmailManager, error) {
	em := &emailManager{
		Config:    config,
		Recaptcha: rcptch,
		Logger:    logger,
		Sender:    es,
	}

	return em, nil
}
