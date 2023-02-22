// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: goakt/v1/remoting.proto

package goaktv1

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

// RemotingServiceClient is the client API for RemotingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RemotingServiceClient interface {
	// RemoteAsk is used to send a message to an actor remotely and expect a response
	// immediately. With this type of message the receiver cannot communicate back to Sender
	// except reply the message with a response. This one-way communication
	RemoteAsk(ctx context.Context, in *RemoteAskRequest, opts ...grpc.CallOption) (*RemoteAskResponse, error)
	// RemoteTell is used to send a message to an actor remotely by another actor
	// This is the only way remote actors can interact with each other. The actor on the
	// other line can reply to the sender by using the Sender in the message
	RemoteTell(ctx context.Context, in *RemoteTellRequest, opts ...grpc.CallOption) (*RemoteTellResponse, error)
	// Lookup for an actor on a remote host.
	RemoteLookup(ctx context.Context, in *RemoteLookupRequest, opts ...grpc.CallOption) (*RemoteLookupResponse, error)
}

type remotingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRemotingServiceClient(cc grpc.ClientConnInterface) RemotingServiceClient {
	return &remotingServiceClient{cc}
}

func (c *remotingServiceClient) RemoteAsk(ctx context.Context, in *RemoteAskRequest, opts ...grpc.CallOption) (*RemoteAskResponse, error) {
	out := new(RemoteAskResponse)
	err := c.cc.Invoke(ctx, "/goakt.v1.RemotingService/RemoteAsk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remotingServiceClient) RemoteTell(ctx context.Context, in *RemoteTellRequest, opts ...grpc.CallOption) (*RemoteTellResponse, error) {
	out := new(RemoteTellResponse)
	err := c.cc.Invoke(ctx, "/goakt.v1.RemotingService/RemoteTell", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remotingServiceClient) RemoteLookup(ctx context.Context, in *RemoteLookupRequest, opts ...grpc.CallOption) (*RemoteLookupResponse, error) {
	out := new(RemoteLookupResponse)
	err := c.cc.Invoke(ctx, "/goakt.v1.RemotingService/RemoteLookup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RemotingServiceServer is the server API for RemotingService service.
// All implementations should embed UnimplementedRemotingServiceServer
// for forward compatibility
type RemotingServiceServer interface {
	// RemoteAsk is used to send a message to an actor remotely and expect a response
	// immediately. With this type of message the receiver cannot communicate back to Sender
	// except reply the message with a response. This one-way communication
	RemoteAsk(context.Context, *RemoteAskRequest) (*RemoteAskResponse, error)
	// RemoteTell is used to send a message to an actor remotely by another actor
	// This is the only way remote actors can interact with each other. The actor on the
	// other line can reply to the sender by using the Sender in the message
	RemoteTell(context.Context, *RemoteTellRequest) (*RemoteTellResponse, error)
	// Lookup for an actor on a remote host.
	RemoteLookup(context.Context, *RemoteLookupRequest) (*RemoteLookupResponse, error)
}

// UnimplementedRemotingServiceServer should be embedded to have forward compatible implementations.
type UnimplementedRemotingServiceServer struct {
}

func (UnimplementedRemotingServiceServer) RemoteAsk(context.Context, *RemoteAskRequest) (*RemoteAskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoteAsk not implemented")
}
func (UnimplementedRemotingServiceServer) RemoteTell(context.Context, *RemoteTellRequest) (*RemoteTellResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoteTell not implemented")
}
func (UnimplementedRemotingServiceServer) RemoteLookup(context.Context, *RemoteLookupRequest) (*RemoteLookupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoteLookup not implemented")
}

// UnsafeRemotingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RemotingServiceServer will
// result in compilation errors.
type UnsafeRemotingServiceServer interface {
	mustEmbedUnimplementedRemotingServiceServer()
}

func RegisterRemotingServiceServer(s grpc.ServiceRegistrar, srv RemotingServiceServer) {
	s.RegisterService(&RemotingService_ServiceDesc, srv)
}

func _RemotingService_RemoteAsk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoteAskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemotingServiceServer).RemoteAsk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goakt.v1.RemotingService/RemoteAsk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemotingServiceServer).RemoteAsk(ctx, req.(*RemoteAskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemotingService_RemoteTell_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoteTellRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemotingServiceServer).RemoteTell(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goakt.v1.RemotingService/RemoteTell",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemotingServiceServer).RemoteTell(ctx, req.(*RemoteTellRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemotingService_RemoteLookup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoteLookupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemotingServiceServer).RemoteLookup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goakt.v1.RemotingService/RemoteLookup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemotingServiceServer).RemoteLookup(ctx, req.(*RemoteLookupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RemotingService_ServiceDesc is the grpc.ServiceDesc for RemotingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RemotingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "goakt.v1.RemotingService",
	HandlerType: (*RemotingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RemoteAsk",
			Handler:    _RemotingService_RemoteAsk_Handler,
		},
		{
			MethodName: "RemoteTell",
			Handler:    _RemotingService_RemoteTell_Handler,
		},
		{
			MethodName: "RemoteLookup",
			Handler:    _RemotingService_RemoteLookup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "goakt/v1/remoting.proto",
}
