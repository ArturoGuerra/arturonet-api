package recaptcha

import (
	"net/url"
	"encoding/json"
	"io/ioutil"
)

type (
	// ReqPayload recaptcha request
	ReqPayload struct {
		Secret   string `json:"secret"`
		Response string `josn:"response"`
		//RemoteIP string `json:"remoteip"`
	}

	// RespPayload recaptcha response
	RespPayload struct {
		Success     bool     `json:"success"`
		ChallengeTS string   `json:"challenge_ts"`
		Hostname    string   `json:"hostname"`
		ErrorCodes  []string `json:"error-codes"`
	}
)

// RecaptchaEndpoint google's v3 recaptcha endpoint
const RecaptchaEndpoint = "https://www.google.com/recaptcha/api/siteverify"

func (r *recaptcha) Validate(token, ip string) (bool, error) {
	form := url.Values{
		"secret": {r.Config.Secret},
		"response": {token},
	}

	resp, err := r.Client.PostForm(RecaptchaEndpoint, form)
	if err != nil {
		r.Logger.Error(err)
		return false, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		r.Logger.Error(err)
		return false, err
	}

	defer resp.Body.Close()

	rpayload := RespPayload{}
	if err = json.Unmarshal(body, &rpayload); err != nil {
		r.Logger.Error(err)
		return false, err
	}

	if rpayload.Success {
		return true, nil
	}

	return false, nil
}
