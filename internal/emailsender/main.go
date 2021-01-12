package emailsender

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/sirupsen/logrus"
)

type (
	// EmailSender exports an SES wrapper
	EmailSender interface {
		Send(payload *Email) error
	}

	emailSender struct {
		Config *Config
		Logger *logrus.Logger
		SES    *ses.SES
	}
)

// New returns an instance of emailSender
func New(logger *logrus.Logger, config *Config) (EmailSender, error) {
	awsSession := session.New(&aws.Config{
		Region:      aws.String(config.Region),
		Credentials: credentials.NewStaticCredentials(config.AccessKeyID, config.SecretAccessKey, ""),
	})

	sesSession := ses.New(awsSession)

	es := &emailSender{
		Logger: logger,
		Config: config,
		SES:    sesSession,
	}

	return es, nil
}

// NewDefault returns an instance of emailSender
func NewDefault(logger *logrus.Logger) (EmailSender, error) {
	cfg, err := loadConfig()
	if err != nil {
		return nil, err
	}

	return New(logger, cfg)
}
