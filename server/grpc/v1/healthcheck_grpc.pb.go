// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: v1/healthcheck.proto

package protobuf

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

// HealthCheckServiceClient is the client API for HealthCheckService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HealthCheckServiceClient interface {
	GetStatus(ctx context.Context, in *GetStatusRequest, opts ...grpc.CallOption) (*GetStatusResponse, error)
}

type healthCheckServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHealthCheckServiceClient(cc grpc.ClientConnInterface) HealthCheckServiceClient {
	return &healthCheckServiceClient{cc}
}

func (c *healthCheckServiceClient) GetStatus(ctx context.Context, in *GetStatusRequest, opts ...grpc.CallOption) (*GetStatusResponse, error) {
	out := new(GetStatusResponse)
	err := c.cc.Invoke(ctx, "/healthcheck.HealthCheckService/GetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthCheckServiceServer is the server API for HealthCheckService service.
// All implementations must embed UnimplementedHealthCheckServiceServer
// for forward compatibility
type HealthCheckServiceServer interface {
	GetStatus(context.Context, *GetStatusRequest) (*GetStatusResponse, error)
	mustEmbedUnimplementedHealthCheckServiceServer()
}

// UnimplementedHealthCheckServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHealthCheckServiceServer struct {
}

func (UnimplementedHealthCheckServiceServer) GetStatus(context.Context, *GetStatusRequest) (*GetStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (UnimplementedHealthCheckServiceServer) mustEmbedUnimplementedHealthCheckServiceServer() {}

// UnsafeHealthCheckServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HealthCheckServiceServer will
// result in compilation errors.
type UnsafeHealthCheckServiceServer interface {
	mustEmbedUnimplementedHealthCheckServiceServer()
}

func RegisterHealthCheckServiceServer(s grpc.ServiceRegistrar, srv HealthCheckServiceServer) {
	s.RegisterService(&HealthCheckService_ServiceDesc, srv)
}

func _HealthCheckService_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthCheckServiceServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/healthcheck.HealthCheckService/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthCheckServiceServer).GetStatus(ctx, req.(*GetStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HealthCheckService_ServiceDesc is the grpc.ServiceDesc for HealthCheckService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HealthCheckService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "healthcheck.HealthCheckService",
	HandlerType: (*HealthCheckServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStatus",
			Handler:    _HealthCheckService_GetStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/healthcheck.proto",
}
