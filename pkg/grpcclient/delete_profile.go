package grpcclient

import (
	"context"
	"fmt"

	pb "github.com/okarpova/my-app/gen/grpc/profile_v1"
)

func (c *Client) Delete(id string) error {
	input := &pb.DeleteProfileInput{
		Id: id,
	}

	_, err := c.client.DeleteProfile(context.Background(), input)
	if err != nil {
		return fmt.Errorf("client.DeleteProfile: %w", err)
	}

	return nil
}
