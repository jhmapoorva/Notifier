// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package Service

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

// EndpointsClient is the client API for Endpoints service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EndpointsClient interface {
	Notify(ctx context.Context, in *NotifyRequest, opts ...grpc.CallOption) (*NotifyResponse, error)
}

type endpointsClient struct {
	cc grpc.ClientConnInterface
}

func NewEndpointsClient(cc grpc.ClientConnInterface) EndpointsClient {
	return &endpointsClient{cc}
}

func (c *endpointsClient) Notify(ctx context.Context, in *NotifyRequest, opts ...grpc.CallOption) (*NotifyResponse, error) {
	out := new(NotifyResponse)
	err := c.cc.Invoke(ctx, "/pb.pb.Endpoints/Notify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EndpointsServer is the server API for Endpoints service.
// All implementations must embed UnimplementedEndpointsServer
// for forward compatibility
type EndpointsServer interface {
	Notify(context.Context, *NotifyRequest) (*NotifyResponse, error)
	mustEmbedUnimplementedEndpointsServer()
}

// UnimplementedEndpointsServer must be embedded to have forward compatible implementations.
type UnimplementedEndpointsServer struct {
}

func (UnimplementedEndpointsServer) Notify(context.Context, *NotifyRequest) (*NotifyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Notify not implemented")
}
func (UnimplementedEndpointsServer) mustEmbedUnimplementedEndpointsServer() {}

// UnsafeEndpointsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EndpointsServer will
// result in compilation errors.
type UnsafeEndpointsServer interface {
	mustEmbedUnimplementedEndpointsServer()
}

func RegisterEndpointsServer(s grpc.ServiceRegistrar, srv EndpointsServer) {
	s.RegisterService(&Endpoints_ServiceDesc, srv)
}

func _Endpoints_Notify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EndpointsServer).Notify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.pb.Endpoints/Notify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EndpointsServer).Notify(ctx, req.(*NotifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Endpoints_ServiceDesc is the grpc.ServiceDesc for Endpoints service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Endpoints_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.pb.Endpoints",
	HandlerType: (*EndpointsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Notify",
			Handler:    _Endpoints_Notify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/notifier.proto",
}
