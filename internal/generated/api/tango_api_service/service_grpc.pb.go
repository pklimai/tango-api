// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.0
// source: tango_api_service/service.proto

package tango_api_service

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	TangoApiService_GetTangoParamsV1_FullMethodName = "/gitlab.com.zigal0_group.nica.tango_api.api.tango_api_service.TangoApiService/GetTangoParamsV1"
)

// TangoApiServiceClient is the client API for TangoApiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TangoApiServiceClient interface {
	// Get tango params by filter.
	GetTangoParamsV1(ctx context.Context, in *GetTangoParamsV1Request, opts ...grpc.CallOption) (*GetTangoParamsV1Response, error)
}

type tangoApiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTangoApiServiceClient(cc grpc.ClientConnInterface) TangoApiServiceClient {
	return &tangoApiServiceClient{cc}
}

func (c *tangoApiServiceClient) GetTangoParamsV1(ctx context.Context, in *GetTangoParamsV1Request, opts ...grpc.CallOption) (*GetTangoParamsV1Response, error) {
	out := new(GetTangoParamsV1Response)
	err := c.cc.Invoke(ctx, TangoApiService_GetTangoParamsV1_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TangoApiServiceServer is the server API for TangoApiService service.
// All implementations must embed UnimplementedTangoApiServiceServer
// for forward compatibility
type TangoApiServiceServer interface {
	// Get tango params by filter.
	GetTangoParamsV1(context.Context, *GetTangoParamsV1Request) (*GetTangoParamsV1Response, error)
	mustEmbedUnimplementedTangoApiServiceServer()
}

// UnimplementedTangoApiServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTangoApiServiceServer struct {
}

func (UnimplementedTangoApiServiceServer) GetTangoParamsV1(context.Context, *GetTangoParamsV1Request) (*GetTangoParamsV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTangoParamsV1 not implemented")
}
func (UnimplementedTangoApiServiceServer) mustEmbedUnimplementedTangoApiServiceServer() {}

// UnsafeTangoApiServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TangoApiServiceServer will
// result in compilation errors.
type UnsafeTangoApiServiceServer interface {
	mustEmbedUnimplementedTangoApiServiceServer()
}

func RegisterTangoApiServiceServer(s grpc.ServiceRegistrar, srv TangoApiServiceServer) {
	s.RegisterService(&TangoApiService_ServiceDesc, srv)
}

func _TangoApiService_GetTangoParamsV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTangoParamsV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TangoApiServiceServer).GetTangoParamsV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TangoApiService_GetTangoParamsV1_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TangoApiServiceServer).GetTangoParamsV1(ctx, req.(*GetTangoParamsV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

// TangoApiService_ServiceDesc is the grpc.ServiceDesc for TangoApiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TangoApiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gitlab.com.zigal0_group.nica.tango_api.api.tango_api_service.TangoApiService",
	HandlerType: (*TangoApiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTangoParamsV1",
			Handler:    _TangoApiService_GetTangoParamsV1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tango_api_service/service.proto",
}
