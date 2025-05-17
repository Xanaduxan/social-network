package v1

import (
	pb "github.com/okarpova/my-app/gen/grpc/profile_v1"
	"github.com/okarpova/my-app/internal/usecase"
)

type Handlers struct {
	pb.UnimplementedProfileV1Server
	usecase *usecase.UseCase
}

func New(uc *usecase.UseCase) *Handlers {
	return &Handlers{
		usecase: uc,
	}
}
