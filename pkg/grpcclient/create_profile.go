package grpcclient

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	pb "gitlab.golang-school.ru/potok-1/okarpova/my-app/gen/grpc/profile_v1"
)

func (c *Client) Create(name string, age int, email, phone string) (uuid.UUID, error) {
	input := &pb.CreateProfileInput{
		Name:  name,
		Age:   int32(age),
		Email: email,
		Phone: phone,
	}

	resp, err := c.client.CreateProfile(context.Background(), input)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("c.client.CreateProfile: %w", err)
	}

	return uuid.Parse(resp.Id)
}
