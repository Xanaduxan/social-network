//go:build grpc

package test

import (
	"github.com/okarpova/my-app/pkg/grpcclient"
)

type ProfileClient = grpcclient.Client

func BuildProfile(s *Suite) {
	var err error
	s.profile, err = grpcclient.New("localhost:50051")
	s.NoError(err)
}
