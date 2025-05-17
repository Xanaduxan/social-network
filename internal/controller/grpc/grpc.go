package grpc

import (
	"fmt"
	"net"

	"gitlab.golang-school.ru/potok-1/okarpova/my-app/pkg/otel"

	"gitlab.golang-school.ru/potok-1/okarpova/my-app/pkg/logger"

	"google.golang.org/grpc/reflection"

	"github.com/rs/zerolog/log"
	pb "gitlab.golang-school.ru/potok-1/okarpova/my-app/gen/grpc/profile_v1"
	ver1 "gitlab.golang-school.ru/potok-1/okarpova/my-app/internal/controller/grpc/v1"
	"gitlab.golang-school.ru/potok-1/okarpova/my-app/internal/usecase"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Config struct {
	Port string `envconfig:"GRPC_PORT" default:"50051"`
}

type Server struct {
	server *grpc.Server
}

func New(c Config, uc *usecase.UseCase) (*Server, error) {
	s := grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
		grpc.ChainUnaryInterceptor(logger.Interceptor, otel.Interceptor),
		// grpc.ChainUnaryInterceptor(logger.First, logger.Second),
	)

	reflection.Register(s)

	v1 := ver1.New(uc)
	pb.RegisterProfileV1Server(s, v1)

	err := start(s, c.Port)
	if err != nil {
		return nil, fmt.Errorf("start: %w", err)
	}

	return &Server{server: s}, nil
}

func start(server *grpc.Server, port string) error {
	conn, err := net.Listen("tcp", net.JoinHostPort("", port))
	if err != nil {
		return fmt.Errorf("net.Listen: %w", err)
	}

	go func() {
		if err := server.Serve(conn); err != nil {
			log.Error().Err(err).Msg("grpc server: Serve")
		}
	}()

	log.Info().Msg("grpc server: started on port: " + port)

	return nil
}

func (s *Server) Close() {
	s.server.GracefulStop()

	log.Info().Msg("grpc server: closed")
}
