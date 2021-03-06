// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package chat

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

// ChittyChatServiceClient is the client API for ChittyChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChittyChatServiceClient interface {
	JoinRoom(ctx context.Context, in *Room, opts ...grpc.CallOption) (ChittyChatService_JoinRoomClient, error)
	LeaveRoom(ctx context.Context, in *Room, opts ...grpc.CallOption) (ChittyChatService_LeaveRoomClient, error)
	SendMessage(ctx context.Context, opts ...grpc.CallOption) (ChittyChatService_SendMessageClient, error)
}

type chittyChatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChittyChatServiceClient(cc grpc.ClientConnInterface) ChittyChatServiceClient {
	return &chittyChatServiceClient{cc}
}

func (c *chittyChatServiceClient) JoinRoom(ctx context.Context, in *Room, opts ...grpc.CallOption) (ChittyChatService_JoinRoomClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChittyChatService_ServiceDesc.Streams[0], "/chat.ChittyChatService/JoinRoom", opts...)
	if err != nil {
		return nil, err
	}
	x := &chittyChatServiceJoinRoomClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChittyChatService_JoinRoomClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type chittyChatServiceJoinRoomClient struct {
	grpc.ClientStream
}

func (x *chittyChatServiceJoinRoomClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chittyChatServiceClient) LeaveRoom(ctx context.Context, in *Room, opts ...grpc.CallOption) (ChittyChatService_LeaveRoomClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChittyChatService_ServiceDesc.Streams[1], "/chat.ChittyChatService/LeaveRoom", opts...)
	if err != nil {
		return nil, err
	}
	x := &chittyChatServiceLeaveRoomClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChittyChatService_LeaveRoomClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type chittyChatServiceLeaveRoomClient struct {
	grpc.ClientStream
}

func (x *chittyChatServiceLeaveRoomClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chittyChatServiceClient) SendMessage(ctx context.Context, opts ...grpc.CallOption) (ChittyChatService_SendMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChittyChatService_ServiceDesc.Streams[2], "/chat.ChittyChatService/SendMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &chittyChatServiceSendMessageClient{stream}
	return x, nil
}

type ChittyChatService_SendMessageClient interface {
	Send(*Message) error
	CloseAndRecv() (*MessageAck, error)
	grpc.ClientStream
}

type chittyChatServiceSendMessageClient struct {
	grpc.ClientStream
}

func (x *chittyChatServiceSendMessageClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chittyChatServiceSendMessageClient) CloseAndRecv() (*MessageAck, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(MessageAck)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChittyChatServiceServer is the server API for ChittyChatService service.
// All implementations must embed UnimplementedChittyChatServiceServer
// for forward compatibility
type ChittyChatServiceServer interface {
	JoinRoom(*Room, ChittyChatService_JoinRoomServer) error
	LeaveRoom(*Room, ChittyChatService_LeaveRoomServer) error
	SendMessage(ChittyChatService_SendMessageServer) error
	mustEmbedUnimplementedChittyChatServiceServer()
}

// UnimplementedChittyChatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChittyChatServiceServer struct {
}

func (UnimplementedChittyChatServiceServer) JoinRoom(*Room, ChittyChatService_JoinRoomServer) error {
	return status.Errorf(codes.Unimplemented, "method JoinRoom not implemented")
}
func (UnimplementedChittyChatServiceServer) LeaveRoom(*Room, ChittyChatService_LeaveRoomServer) error {
	return status.Errorf(codes.Unimplemented, "method LeaveRoom not implemented")
}
func (UnimplementedChittyChatServiceServer) SendMessage(ChittyChatService_SendMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedChittyChatServiceServer) mustEmbedUnimplementedChittyChatServiceServer() {}

// UnsafeChittyChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChittyChatServiceServer will
// result in compilation errors.
type UnsafeChittyChatServiceServer interface {
	mustEmbedUnimplementedChittyChatServiceServer()
}

func RegisterChittyChatServiceServer(s grpc.ServiceRegistrar, srv ChittyChatServiceServer) {
	s.RegisterService(&ChittyChatService_ServiceDesc, srv)
}

func _ChittyChatService_JoinRoom_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Room)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChittyChatServiceServer).JoinRoom(m, &chittyChatServiceJoinRoomServer{stream})
}

type ChittyChatService_JoinRoomServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type chittyChatServiceJoinRoomServer struct {
	grpc.ServerStream
}

func (x *chittyChatServiceJoinRoomServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func _ChittyChatService_LeaveRoom_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Room)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChittyChatServiceServer).LeaveRoom(m, &chittyChatServiceLeaveRoomServer{stream})
}

type ChittyChatService_LeaveRoomServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type chittyChatServiceLeaveRoomServer struct {
	grpc.ServerStream
}

func (x *chittyChatServiceLeaveRoomServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func _ChittyChatService_SendMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChittyChatServiceServer).SendMessage(&chittyChatServiceSendMessageServer{stream})
}

type ChittyChatService_SendMessageServer interface {
	SendAndClose(*MessageAck) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type chittyChatServiceSendMessageServer struct {
	grpc.ServerStream
}

func (x *chittyChatServiceSendMessageServer) SendAndClose(m *MessageAck) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chittyChatServiceSendMessageServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChittyChatService_ServiceDesc is the grpc.ServiceDesc for ChittyChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChittyChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chat.ChittyChatService",
	HandlerType: (*ChittyChatServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "JoinRoom",
			Handler:       _ChittyChatService_JoinRoom_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "LeaveRoom",
			Handler:       _ChittyChatService_LeaveRoom_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SendMessage",
			Handler:       _ChittyChatService_SendMessage_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "Chat/chat.proto",
}
