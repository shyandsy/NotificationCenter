// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.2
// source: proto/proto.proto

package __

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

// NotificationClient is the client API for Notification service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotificationClient interface {
	Subscribe(ctx context.Context, in *Topic, opts ...grpc.CallOption) (Notification_SubscribeClient, error)
}

type notificationClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationClient(cc grpc.ClientConnInterface) NotificationClient {
	return &notificationClient{cc}
}

func (c *notificationClient) Subscribe(ctx context.Context, in *Topic, opts ...grpc.CallOption) (Notification_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &Notification_ServiceDesc.Streams[0], "/pb.Notification/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &notificationSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Notification_SubscribeClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type notificationSubscribeClient struct {
	grpc.ClientStream
}

func (x *notificationSubscribeClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NotificationServer is the server API for Notification service.
// All implementations should embed UnimplementedNotificationServer
// for forward compatibility
type NotificationServer interface {
	Subscribe(*Topic, Notification_SubscribeServer) error
}

// UnimplementedNotificationServer should be embedded to have forward compatible implementations.
type UnimplementedNotificationServer struct {
}

func (UnimplementedNotificationServer) Subscribe(*Topic, Notification_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}

// UnsafeNotificationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotificationServer will
// result in compilation errors.
type UnsafeNotificationServer interface {
	mustEmbedUnimplementedNotificationServer()
}

func RegisterNotificationServer(s grpc.ServiceRegistrar, srv NotificationServer) {
	s.RegisterService(&Notification_ServiceDesc, srv)
}

func _Notification_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Topic)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NotificationServer).Subscribe(m, &notificationSubscribeServer{stream})
}

type Notification_SubscribeServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type notificationSubscribeServer struct {
	grpc.ServerStream
}

func (x *notificationSubscribeServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

// Notification_ServiceDesc is the grpc.ServiceDesc for Notification service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Notification_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Notification",
	HandlerType: (*NotificationServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _Notification_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/proto.proto",
}
