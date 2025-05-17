package httpclientv2

import (
	"errors"
	"fmt"

	http_client "github.com/okarpova/my-app/gen/http/profile_v2/client"
)

var ErrNotFound = errors.New("not found")

type Client struct {
	client *http_client.ClientWithResponses
}

func New(host string) (*Client, error) {
	client, err := http_client.NewClientWithResponses(host)
	if err != nil {
		return nil, fmt.Errorf("http_client.NewClient: %w", err)
	}

	return &Client{client: client}, nil
}
