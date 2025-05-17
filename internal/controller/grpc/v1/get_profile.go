package v1

import (
	"context"
	"errors"

	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/okarpova/my-app/gen/grpc/profile_v1"
	"github.com/okarpova/my-app/internal/domain"
	"github.com/okarpova/my-app/internal/dto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h Handlers) GetProfile(ctx context.Context, i *pb.GetProfileInput) (*pb.GetProfileOutput, error) {
	input := dto.GetProfileInput{
		ID: i.Id,
	}

	o, err := h.usecase.GetProfile(ctx, input)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrNotFound):
			return nil, status.Error(codes.NotFound, "not found")

		default:
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
	}

	return &pb.GetProfileOutput{
		Id:        o.ID.String(),
		CreatedAt: timestamppb.New(o.CreatedAt),
		UpdatedAt: timestamppb.New(o.UpdatedAt),
		Name:      string(o.Name),
		Age:       int32(o.Age),
		Verified:  o.Verified,
		Status:    int32(o.Status),
		Contacts: &pb.GetProfileOutput_Contacts{
			Email: o.Contacts.Email,
			Phone: o.Contacts.Phone,
		},
	}, nil
}
