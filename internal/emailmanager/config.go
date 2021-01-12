package emailmanager

import "github.com/arturoguerra/arturonet-api/internal/utils"

// Config contains all the runtime config data for EmailManager
type Config struct{}

func loadConfig() (*Config, error) {
	cfg := Config{}

	if err := utils.EnvLoader(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
