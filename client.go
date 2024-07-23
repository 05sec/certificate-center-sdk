package client

import (
	"certificate-center-sdk/config"
	"net/http"
)

var Client *client

type client struct {
	httpClient *http.Client
	config     *config.Config
}

func NewClient(config *config.Config) *client {
	Client.httpClient = &http.Client{}
	Client.config = config
	return Client
}
