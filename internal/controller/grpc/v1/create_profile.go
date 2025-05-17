package v1

import (
	"context"

	pb "gitlab.golang-school.ru/potok-1/okarpova/my-app/gen/grpc/profile_v1"
	"gitlab.golang-school.ru/potok-1/okarpova/my-app/internal/dto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h Handlers) CreateProfile(ctx context.Context, i *pb.CreateProfileInput) (*pb.CreateProfileOutput, error) {
	input := dto.CreateProfileInput{
		Name:  i.Name,
		Age:   int(i.Age),
		Email: i.Email,
		Phone: i.Phone,
	}

	output, err := h.usecase.CreateProfile(ctx, input)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &pb.CreateProfileOutput{
		Id: output.ID.String(),
	}, nil
}
