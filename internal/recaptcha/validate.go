package recaptcha

import (
	"bytes"
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
	payload := &ReqPayload{
		Secret:   r.Config.Secret,
		Response: token,
		//		RemoteIP: ip,
	}

	form := url.Values{
		"secret": r.Config.Secret,
		"response": token,
	}

	jsonpayload, err := json.Marshal(payload)
	if err != nil {
		r.Logger.Error(err)
		return false, err
	}

	resp, err := r.Client.PostForm(RecaptchaEndpoint, from)
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
