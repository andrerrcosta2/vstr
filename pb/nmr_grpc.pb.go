// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.2
// source: crdnmr.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Grpcnms_Snm_FullMethodName = "/msg.Grpcnms/Snm"
)

// GrpcnmsClient is the client API for Grpcnms service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GrpcnmsClient interface {
	Snm(ctx context.Context, in *Nms, opts ...grpc.CallOption) (*Nms, error)
}

type grpcnmsClient struct {
	cc grpc.ClientConnInterface
}

func NewGrpcnmsClient(cc grpc.ClientConnInterface) GrpcnmsClient {
	return &grpcnmsClient{cc}
}

func (c *grpcnmsClient) Snm(ctx context.Context, in *Nms, opts ...grpc.CallOption) (*Nms, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Nms)
	err := c.cc.Invoke(ctx, Grpcnms_Snm_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GrpcnmsServer is the server API for Grpcnms service.
// All implementations must embed UnimplementedGrpcnmsServer
// for forward compatibility
type GrpcnmsServer interface {
	Snm(context.Context, *Nms) (*Nms, error)
	mustEmbedUnimplementedGrpcnmsServer()
}

// UnimplementedGrpcnmsServer must be embedded to have forward compatible implementations.
type UnimplementedGrpcnmsServer struct {
}

func (UnimplementedGrpcnmsServer) Snm(context.Context, *Nms) (*Nms, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Snm not implemented")
}
func (UnimplementedGrpcnmsServer) mustEmbedUnimplementedGrpcnmsServer() {}

// UnsafeGrpcnmsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GrpcnmsServer will
// result in compilation errors.
type UnsafeGrpcnmsServer interface {
	mustEmbedUnimplementedGrpcnmsServer()
}

func RegisterGrpcnmsServer(s grpc.ServiceRegistrar, srv GrpcnmsServer) {
	s.RegisterService(&Grpcnms_ServiceDesc, srv)
}

func _Grpcnms_Snm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Nms)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcnmsServer).Snm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Grpcnms_Snm_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcnmsServer).Snm(ctx, req.(*Nms))
	}
	return interceptor(ctx, in, info, handler)
}

// Grpcnms_ServiceDesc is the grpc.ServiceDesc for Grpcnms service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Grpcnms_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "msg.Grpcnms",
	HandlerType: (*GrpcnmsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Snm",
			Handler:    _Grpcnms_Snm_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "crdnmr.proto",
}
