package emailmanager

import (
	"net/http"

	"github.com/arturoguerra/arturonet-api/internal/emailsender"
	"github.com/labstack/echo/v4"
)

type (
	// EmailPayload is the request payload for sending an email
	EmailPayload struct {
		Message   string `json:"message"`
		Email     string `json:"email"`
		Name      string `json:"name"`
		Recaptcha string `json:"recaptcha"`
	}
)

func (em *emailManager) send(c echo.Context) error {
	payload := new(EmailPayload)
	if err := c.Bind(payload); err != nil {
		em.Logger.Error(err)
		return err
	}

	_, err := em.Recaptcha.Validate(payload.Recaptcha, c.RealIP())
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	//if !valid {
	//	return c.String(http.StatusForbidden, "Invalid Recaptcha Token")
	//}

	email := &emailsender.Email{
		Name:    payload.Name,
		Email:   payload.Email,
		Message: payload.Message,
	}

	if err = em.Sender.Send(email); err != nil {
		return c.String(http.StatusForbidden, err.Error())
	}

	return c.String(http.StatusOK, "200")
}
