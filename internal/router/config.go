package router

import "github.com/arturoguerra/arturonet-api/internal/utils"

type (
	// Config is the configuration for echo router
	Config struct {
		Host string `env:"HOST" envDefault:"0.0.0.0"`
		Port string `env:"PORT" envDefault:"3333"`
	}
)

func loadConfig() (*Config, error) {
	cfg := Config{}

	if err := utils.EnvLoader(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
