package recaptcha

import "github.com/arturoguerra/arturonet-api/internal/utils"

// Config recaptcha config
type Config struct {
	Secret string `env:"RECAPTCHA_SECRET,required"`
}

func loadConfig() (*Config, error) {
	cfg := Config{}

	if err := utils.EnvLoader(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
