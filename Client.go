package OtxApiClient

import (
	"errors"
	"io"
	"log"
	"net/http"
)

type Client struct {
	httpClient *http.Client
	userAgent  string
	apiKey     string
	baseUrl    string
}

type OtxOpt func(client *Client) error

func NewClient(apiKey string, opts ...OtxOpt) (*Client, error) {
	client := &Client{
		httpClient: http.DefaultClient,
		userAgent:  "Go-OtxApiClient",
		apiKey:     apiKey,
		baseUrl:    "https://otx.alienvault.com/api/v1",
	}
	for _, opt := range opts {
		err := opt(client)
		if err != nil {
			return nil, err
		}
	}
	return client, nil
}

func (client *Client) GetIndicatorsService() *indicatorsService {
	return &indicatorsService{client: client}
}

func (client *Client) GetUserService() *userService {
	return &userService{client: client}
}

func (client *Client) GetPulsesService() *pulsesService {
	return &pulsesService{client: client}
}

func (client *Client) doHttpRequest(method string, path string, body io.Reader) (*http.Response, error) {
	url := client.baseUrl + path
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-OTX-API-KEY", client.apiKey)
	req.Header.Set("User-Agent", client.userAgent)
	req.Header.Set("Content-Type", "application/json")
	log.Println(url, req)
	return client.httpClient.Do(req)
}

var ErrNoHttpClient = errors.New("no http client")
var ErrNoUserAgent = errors.New("no user-agent")
var ErrNoBaseUrl = errors.New("no base url")

func WithHttpClient(httpClient *http.Client) OtxOpt {
	return func(client *Client) error {
		if httpClient == nil {
			return ErrNoHttpClient
		}
		client.httpClient = httpClient
		return nil
	}
}

func WithUserAgent(userAgent string) OtxOpt {
	return func(client *Client) error {
		if len(userAgent) < 1 {
			return ErrNoUserAgent
		}
		client.userAgent = userAgent
		return nil
	}
}

func WithBaseUrl(baseUrl string) OtxOpt {
	return func(client *Client) error {
		if len(baseUrl) < 1 {
			return ErrNoBaseUrl
		}
		client.baseUrl = baseUrl
		return nil
	}
}
