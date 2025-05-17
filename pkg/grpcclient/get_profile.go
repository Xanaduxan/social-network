package grpcclient

import (
	"context"
	"fmt"

	"github.com/okarpova/my-app/pkg/httpclient"

	pb "github.com/okarpova/my-app/gen/grpc/profile_v1"
)

type Profile httpclient.Profile

func (c *Client) Get(id string) (Profile, error) {
	input := &pb.GetProfileInput{
		Id: id,
	}

	o, err := c.client.GetProfile(context.Background(), input)
	if err != nil {
		return Profile{}, fmt.Errorf("client.Get: %w", err)
	}

	return Profile{
		ID:        o.Id,
		CreatedAt: o.CreatedAt.String(),
		UpdatedAt: o.UpdatedAt.String(),
		Name:      o.Name,
		Age:       int(o.Age),
		Status:    int(o.Status),
		Verified:  o.Verified,
		Contacts: struct {
			Email string `json:"email"`
			Phone string `json:"phone"`
		}{
			Email: o.Contacts.Email,
			Phone: o.Contacts.Phone,
		},
	}, nil
}
