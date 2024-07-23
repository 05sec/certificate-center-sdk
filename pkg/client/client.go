package client

import (
	"net/http"
)

type Client struct {
	httpClient *http.Client
	config     *Config
}

func NewClient(config *Config) *Client {
	if config.BaseURL == "" {
		config.BaseURL = "http://license-center.lwsec.cn"
	}

	return &Client{
		httpClient: &http.Client{},
		config:     config,
	}
}
