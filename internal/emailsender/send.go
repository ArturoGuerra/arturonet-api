package emailsender

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ses"
)

// Email is the request payload
type Email struct {
	Name    string
	Email   string
	Message string
}

func (es *emailSender) Send(payload *Email) error {
	sesEmailInput := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{aws.String(es.Config.ReplyEmail)},
		},
		Source:           aws.String(es.Config.SenderEmail),
		ReplyToAddresses: []*string{aws.String(es.Config.ReplyEmail)},
		Message: &ses.Message{
			Subject: &ses.Content{
				Data: aws.String(fmt.Sprintf("Email from ArturoNet Name: %s Email: %s", payload.Name, payload.Email)),
			},
			Body: &ses.Body{
				Text: &ses.Content{
					Data: aws.String(payload.Message),
				},
			},
		},
	}

	if _, err := es.SES.SendEmail(sesEmailInput); err != nil {
		es.Logger.Error(err)
		return err
	}

	return nil
}
