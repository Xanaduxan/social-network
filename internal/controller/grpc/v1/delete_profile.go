package v1

import (
	"context"

	pb "gitlab.golang-school.ru/potok-1/okarpova/my-app/gen/grpc/profile_v1"
	"gitlab.golang-school.ru/potok-1/okarpova/my-app/internal/dto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h Handlers) DeleteProfile(ctx context.Context, i *pb.DeleteProfileInput) (*emptypb.Empty, error) {
	input := dto.DeleteProfileInput{
		ID: i.Id,
	}

	err := h.usecase.DeleteProfile(ctx, input)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &emptypb.Empty{}, nil
}
