package httpclient

import (
	"errors"
	"net"
	"net/http"
	"time"
)

var ErrNotFound = errors.New("not found")

type Config struct {
	Host string `envconfig:"HTTP_CLIENT_HOST" default:"localhost"`
	Port string `envconfig:"HTTP_CLIENT_PORT" default:"8080"`
}

type Client struct {
	client http.Client
	host   string
}

func New(c Config) *Client {
	return &Client{
		client: http.Client{
			Timeout: 5 * time.Second,
		},
		host: net.JoinHostPort(c.Host, c.Port),
	}
}
