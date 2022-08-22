package lib

import "net/http"

// HttpClient -
type HttpClient struct {
	client *http.Client
}

// NewHttpClient -
func NewHttpClient() (*HttpClient, error) {
	c := HttpClient{
		client: &http.Client{},
	}

	return &c, nil
}
