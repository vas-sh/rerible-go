package client

import "net/http"

type client struct {
	httpClient *http.Client
	rootURL    string
	apiKey     string
}

func New(apiKey string, httpClient *http.Client, rootURL string) *client {
	return &client{
		httpClient: httpClient,
		rootURL:    rootURL,
		apiKey:     apiKey,
	}
}
