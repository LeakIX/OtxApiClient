package OtxApiClient

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type pulsesService struct {
	client *Client
}

type PulseIndicatorType string

const (
	PulseIPv4Indicator     PulseIndicatorType = "IPv4"
	PulseHostnameIndicator PulseIndicatorType = "hostname"
)

var ErrInvalidResponse = errors.New("invalid response")

func (p *pulsesService) AddIndicator(pulseId string, indicator PulseIndicator) error {
	editRequest := &PulseIndicatorEdit{Id: pulseId}
	editRequest.Indicators.Add = []PulseIndicator{indicator}
	postData, err := json.Marshal(editRequest)
	log.Println(string(postData))

	if err != nil {
		return err
	}
	resp, err := p.client.doHttpRequest(http.MethodPatch, "/pulses/"+pulseId+"/", bytes.NewReader(postData))
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return ErrInvalidResponse
	}
	return nil
}

type PulseIndicatorEdit struct {
	Id          string `json:"id"`
	Description string `json:"description,omitempty"`
	Indicators  struct {
		Add []PulseIndicator `json:"add"`
	} `json:"indicators"`
}

type PulseIndicator struct {
	Type      PulseIndicatorType `json:"type"`
	Indicator string             `json:"indicator"`
}
