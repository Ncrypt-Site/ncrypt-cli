package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"ncrypt-cli/models"
	"net/http"
)

type HttpService struct {
	Client *http.Client
	Url    string
}

func (h HttpService) SendRequest(payload models.SecureMessageRequest) ([]byte, error) {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", h.Url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := h.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, errors.New("request resulted in a failed state")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
