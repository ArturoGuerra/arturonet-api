package emailsender

import "github.com/arturoguerra/arturonet-api/internal/utils"

// Config contains all the runtime config data for EmailManager
type Config struct {
	SenderEmail     string `env:"AWS_SENDER_MAIL,required"`
	ReplyEmail      string `env:"AWS_REPLY_EMAIL,required"`
	AccessKeyID     string `env:"AWS_ACCESS_KEY,required"`
	SecretAccessKey string `env:"AWS_SECRET_ACCESS_KEY,required"`
	Region          string `env:"AWS_REGION,required"`
}

func loadConfig() (*Config, error) {
	cfg := Config{}

	if err := utils.EnvLoader(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
