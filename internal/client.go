package internal

import (
	"net/http"
	"time"
)

type Client struct {
	Client http.Client
}

func NewClient() *Client {
	return &Client {
		Client : http.Client{
			Timeout: 5 * time.Second,
		},
	}
}
