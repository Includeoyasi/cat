// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: api/cat.proto

package grpc

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

// CatClient is the client API for Cat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CatClient interface {
	GetCapsule(ctx context.Context, in *GetCapsuleRequest, opts ...grpc.CallOption) (*GetCapsuleResponce, error)
}

type catClient struct {
	cc grpc.ClientConnInterface
}

func NewCatClient(cc grpc.ClientConnInterface) CatClient {
	return &catClient{cc}
}

func (c *catClient) GetCapsule(ctx context.Context, in *GetCapsuleRequest, opts ...grpc.CallOption) (*GetCapsuleResponce, error) {
	out := new(GetCapsuleResponce)
	err := c.cc.Invoke(ctx, "/api.Cat/GetCapsule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CatServer is the server API for Cat service.
// All implementations must embed UnimplementedCatServer
// for forward compatibility
type CatServer interface {
	GetCapsule(context.Context, *GetCapsuleRequest) (*GetCapsuleResponce, error)
	mustEmbedUnimplementedCatServer()
}

// UnimplementedCatServer must be embedded to have forward compatible implementations.
type UnimplementedCatServer struct {
}

func (UnimplementedCatServer) GetCapsule(context.Context, *GetCapsuleRequest) (*GetCapsuleResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCapsule not implemented")
}
func (UnimplementedCatServer) mustEmbedUnimplementedCatServer() {}

// UnsafeCatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CatServer will
// result in compilation errors.
type UnsafeCatServer interface {
	mustEmbedUnimplementedCatServer()
}

func RegisterCatServer(s grpc.ServiceRegistrar, srv CatServer) {
	s.RegisterService(&Cat_ServiceDesc, srv)
}

func _Cat_GetCapsule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCapsuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatServer).GetCapsule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Cat/GetCapsule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatServer).GetCapsule(ctx, req.(*GetCapsuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Cat_ServiceDesc is the grpc.ServiceDesc for Cat service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Cat_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.Cat",
	HandlerType: (*CatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCapsule",
			Handler:    _Cat_GetCapsule_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/cat.proto",
}
