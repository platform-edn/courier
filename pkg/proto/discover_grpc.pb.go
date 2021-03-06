// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// DiscoverClient is the client API for Discover service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DiscoverClient interface {
	SubscribeToNodeEvents(ctx context.Context, in *EventStreamRequest, opts ...grpc.CallOption) (Discover_SubscribeToNodeEventsClient, error)
}

type discoverClient struct {
	cc grpc.ClientConnInterface
}

func NewDiscoverClient(cc grpc.ClientConnInterface) DiscoverClient {
	return &discoverClient{cc}
}

func (c *discoverClient) SubscribeToNodeEvents(ctx context.Context, in *EventStreamRequest, opts ...grpc.CallOption) (Discover_SubscribeToNodeEventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Discover_ServiceDesc.Streams[0], "/proto.Discover/SubscribeToNodeEvents", opts...)
	if err != nil {
		return nil, err
	}
	x := &discoverSubscribeToNodeEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Discover_SubscribeToNodeEventsClient interface {
	Recv() (*EventStreamResponse, error)
	grpc.ClientStream
}

type discoverSubscribeToNodeEventsClient struct {
	grpc.ClientStream
}

func (x *discoverSubscribeToNodeEventsClient) Recv() (*EventStreamResponse, error) {
	m := new(EventStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DiscoverServer is the server API for Discover service.
// All implementations must embed UnimplementedDiscoverServer
// for forward compatibility
type DiscoverServer interface {
	SubscribeToNodeEvents(*EventStreamRequest, Discover_SubscribeToNodeEventsServer) error
	mustEmbedUnimplementedDiscoverServer()
}

// UnimplementedDiscoverServer must be embedded to have forward compatible implementations.
type UnimplementedDiscoverServer struct {
}

func (UnimplementedDiscoverServer) SubscribeToNodeEvents(*EventStreamRequest, Discover_SubscribeToNodeEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeToNodeEvents not implemented")
}
func (UnimplementedDiscoverServer) mustEmbedUnimplementedDiscoverServer() {}

// UnsafeDiscoverServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DiscoverServer will
// result in compilation errors.
type UnsafeDiscoverServer interface {
	mustEmbedUnimplementedDiscoverServer()
}

func RegisterDiscoverServer(s grpc.ServiceRegistrar, srv DiscoverServer) {
	s.RegisterService(&Discover_ServiceDesc, srv)
}

func _Discover_SubscribeToNodeEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EventStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DiscoverServer).SubscribeToNodeEvents(m, &discoverSubscribeToNodeEventsServer{stream})
}

type Discover_SubscribeToNodeEventsServer interface {
	Send(*EventStreamResponse) error
	grpc.ServerStream
}

type discoverSubscribeToNodeEventsServer struct {
	grpc.ServerStream
}

func (x *discoverSubscribeToNodeEventsServer) Send(m *EventStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Discover_ServiceDesc is the grpc.ServiceDesc for Discover service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Discover_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Discover",
	HandlerType: (*DiscoverServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeToNodeEvents",
			Handler:       _Discover_SubscribeToNodeEvents_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "discover.proto",
}
