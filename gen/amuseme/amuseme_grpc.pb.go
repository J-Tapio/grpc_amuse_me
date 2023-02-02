// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: amuseme.proto

package amuseme

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

// AmuseMeClient is the client API for AmuseMe service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AmuseMeClient interface {
	GetJoke(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*JokeReply, error)
	GetMeme(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MemeReply, error)
}

type amuseMeClient struct {
	cc grpc.ClientConnInterface
}

func NewAmuseMeClient(cc grpc.ClientConnInterface) AmuseMeClient {
	return &amuseMeClient{cc}
}

func (c *amuseMeClient) GetJoke(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*JokeReply, error) {
	out := new(JokeReply)
	err := c.cc.Invoke(ctx, "/amuseme.AmuseMe/GetJoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *amuseMeClient) GetMeme(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MemeReply, error) {
	out := new(MemeReply)
	err := c.cc.Invoke(ctx, "/amuseme.AmuseMe/GetMeme", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AmuseMeServer is the server API for AmuseMe service.
// All implementations must embed UnimplementedAmuseMeServer
// for forward compatibility
type AmuseMeServer interface {
	GetJoke(context.Context, *Empty) (*JokeReply, error)
	GetMeme(context.Context, *Empty) (*MemeReply, error)
	mustEmbedUnimplementedAmuseMeServer()
}

// UnimplementedAmuseMeServer must be embedded to have forward compatible implementations.
type UnimplementedAmuseMeServer struct {
}

func (UnimplementedAmuseMeServer) GetJoke(context.Context, *Empty) (*JokeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJoke not implemented")
}
func (UnimplementedAmuseMeServer) GetMeme(context.Context, *Empty) (*MemeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMeme not implemented")
}
func (UnimplementedAmuseMeServer) mustEmbedUnimplementedAmuseMeServer() {}

// UnsafeAmuseMeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AmuseMeServer will
// result in compilation errors.
type UnsafeAmuseMeServer interface {
	mustEmbedUnimplementedAmuseMeServer()
}

func RegisterAmuseMeServer(s grpc.ServiceRegistrar, srv AmuseMeServer) {
	s.RegisterService(&AmuseMe_ServiceDesc, srv)
}

func _AmuseMe_GetJoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AmuseMeServer).GetJoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/amuseme.AmuseMe/GetJoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AmuseMeServer).GetJoke(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AmuseMe_GetMeme_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AmuseMeServer).GetMeme(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/amuseme.AmuseMe/GetMeme",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AmuseMeServer).GetMeme(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// AmuseMe_ServiceDesc is the grpc.ServiceDesc for AmuseMe service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AmuseMe_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "amuseme.AmuseMe",
	HandlerType: (*AmuseMeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetJoke",
			Handler:    _AmuseMe_GetJoke_Handler,
		},
		{
			MethodName: "GetMeme",
			Handler:    _AmuseMe_GetMeme_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "amuseme.proto",
}
