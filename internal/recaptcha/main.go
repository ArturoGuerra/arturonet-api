package recaptcha

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

type (
	// Recaptcha exports recaptcha methods
	Recaptcha interface {
		Validate(token string, ip string) (bool, error)
	}

	recaptcha struct {
		Logger *logrus.Logger
		Config *Config
		Client *http.Client
	}
)

// NewDefault creates a new instance of recaptcha with env config
func NewDefault(logger *logrus.Logger) (Recaptcha, error) {
	cfg, err := loadConfig()
	if err != nil {
		return nil, err
	}

	return New(logger, cfg)
}

// New creates a new instance of recaptcha
func New(logger *logrus.Logger, config *Config) (Recaptcha, error) {
	r := &recaptcha{
		Logger: logger,
		Config: config,
		Client: &http.Client{},
	}

	return r, nil
}
