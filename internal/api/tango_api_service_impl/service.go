package tango_api_service_impl

import (
	"context"

	gw_runtime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "gitlab.com/zigal0-group/nica/tango-api/internal/generated/api/tango_api_service"
	"google.golang.org/grpc"
)

type Service struct {
	pb.UnimplementedTangoApiServiceServer
}

func New() *Service {
	return &Service{}
}

func (s *Service) RegisterGRPC(server *grpc.Server) {
	pb.RegisterTangoApiServiceServer(server, s)
}

func (s *Service) RegisterGatewayEndpoint(
	ctx context.Context,
	mux *gw_runtime.ServeMux,
	endpoint string,
	dialOptions []grpc.DialOption,
) error {
	return pb.RegisterTangoApiServiceHandlerFromEndpoint(ctx, mux, endpoint, dialOptions)
}

func (s *Service) RegisterGateway(
	ctx context.Context,
	mux *gw_runtime.ServeMux,
) error {
	return pb.RegisterTangoApiServiceHandlerServer(ctx, mux, s)
}
