package v1

import (
	"context"

	pb "github.com/okarpova/my-app/gen/grpc/profile_v1"
	"github.com/okarpova/my-app/internal/dto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h Handlers) UpdateProfile(ctx context.Context, i *pb.UpdateProfileInput) (*emptypb.Empty, error) {
	input := dto.UpdateProfileInput{
		ID:    i.Id,
		Name:  i.Name,
		Age:   parseAge(i.Age),
		Email: i.Email,
		Phone: i.Phone,
	}

	err := h.usecase.UpdateProfile(ctx, input)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &emptypb.Empty{}, nil
}

func parseAge(age *int32) *int {
	if age == nil {
		return nil
	}

	a := int(*age)

	return &a
}
