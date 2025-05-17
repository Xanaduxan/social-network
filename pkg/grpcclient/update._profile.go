package grpcclient

import (
	"context"
	"fmt"

	pb "gitlab.golang-school.ru/potok-1/okarpova/my-app/gen/grpc/profile_v1"
)

func (c *Client) Update(id string, name *string, age *int, email, phone *string) error {
	input := &pb.UpdateProfileInput{
		Id:    id,
		Name:  name,
		Age:   parseAge(age),
		Email: email,
		Phone: phone,
	}

	_, err := c.client.UpdateProfile(context.Background(), input)
	if err != nil {
		return fmt.Errorf("c.client.UpdateProfile: %w", err)
	}

	return nil
}

func parseAge(age *int) *int32 {
	if age == nil {
		return nil
	}

	a := int32(*age)

	return &a
}
