// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: v1/publisher/publisher.proto

package publisher

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

// PublisherClient is the client API for Publisher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PublisherClient interface {
	GetSummary(ctx context.Context, in *SummaryRequest, opts ...grpc.CallOption) (Publisher_GetSummaryClient, error)
}

type publisherClient struct {
	cc grpc.ClientConnInterface
}

func NewPublisherClient(cc grpc.ClientConnInterface) PublisherClient {
	return &publisherClient{cc}
}

func (c *publisherClient) GetSummary(ctx context.Context, in *SummaryRequest, opts ...grpc.CallOption) (Publisher_GetSummaryClient, error) {
	stream, err := c.cc.NewStream(ctx, &Publisher_ServiceDesc.Streams[0], "/v1.publisher.Publisher/GetSummary", opts...)
	if err != nil {
		return nil, err
	}
	x := &publisherGetSummaryClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Publisher_GetSummaryClient interface {
	Recv() (*SummaryResponse, error)
	grpc.ClientStream
}

type publisherGetSummaryClient struct {
	grpc.ClientStream
}

func (x *publisherGetSummaryClient) Recv() (*SummaryResponse, error) {
	m := new(SummaryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PublisherServer is the server API for Publisher service.
// All implementations must embed UnimplementedPublisherServer
// for forward compatibility
type PublisherServer interface {
	GetSummary(*SummaryRequest, Publisher_GetSummaryServer) error
	mustEmbedUnimplementedPublisherServer()
}

// UnimplementedPublisherServer must be embedded to have forward compatible implementations.
type UnimplementedPublisherServer struct {
}

func (UnimplementedPublisherServer) GetSummary(*SummaryRequest, Publisher_GetSummaryServer) error {
	return status.Errorf(codes.Unimplemented, "method GetSummary not implemented")
}
func (UnimplementedPublisherServer) mustEmbedUnimplementedPublisherServer() {}

// UnsafePublisherServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PublisherServer will
// result in compilation errors.
type UnsafePublisherServer interface {
	mustEmbedUnimplementedPublisherServer()
}

func RegisterPublisherServer(s grpc.ServiceRegistrar, srv PublisherServer) {
	s.RegisterService(&Publisher_ServiceDesc, srv)
}

func _Publisher_GetSummary_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SummaryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PublisherServer).GetSummary(m, &publisherGetSummaryServer{stream})
}

type Publisher_GetSummaryServer interface {
	Send(*SummaryResponse) error
	grpc.ServerStream
}

type publisherGetSummaryServer struct {
	grpc.ServerStream
}

func (x *publisherGetSummaryServer) Send(m *SummaryResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Publisher_ServiceDesc is the grpc.ServiceDesc for Publisher service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Publisher_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.publisher.Publisher",
	HandlerType: (*PublisherServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetSummary",
			Handler:       _Publisher_GetSummary_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "v1/publisher/publisher.proto",
}
