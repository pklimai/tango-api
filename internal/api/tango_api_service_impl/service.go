package tango_api_service_impl

import (
	"context"

	gw_runtime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"gitlab.com/zigal0-group/nica/tango-api/internal/domain"
	pb "gitlab.com/zigal0-group/nica/tango-api/internal/generated/api/tango_api_service"
	"google.golang.org/grpc"
)

type paramManager interface {
	GetTangoParamByFilter(
		ctx context.Context,
		filter domain.ParamFilter,
	) (domain.Params, error)
}

type Service struct {
	paramManager

	pb.UnimplementedTangoApiServiceServer
}

func New(
	paramManager paramManager,
) *Service {
	return &Service{
		paramManager: paramManager,
	}
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
