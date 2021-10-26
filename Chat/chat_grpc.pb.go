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
	JoinChannel(ctx context.Context, in *Channel, opts ...grpc.CallOption) (ChittyChatService_JoinChannelClient, error)
	LeaveChannel(ctx context.Context, in *Channel, opts ...grpc.CallOption) (*LeaveAck, error)
	SendMessage(ctx context.Context, opts ...grpc.CallOption) (ChittyChatService_SendMessageClient, error)
}

type chittyChatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChittyChatServiceClient(cc grpc.ClientConnInterface) ChittyChatServiceClient {
	return &chittyChatServiceClient{cc}
}

func (c *chittyChatServiceClient) JoinChannel(ctx context.Context, in *Channel, opts ...grpc.CallOption) (ChittyChatService_JoinChannelClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChittyChatService_ServiceDesc.Streams[0], "/chat.ChittyChatService/JoinChannel", opts...)
	if err != nil {
		return nil, err
	}
	x := &chittyChatServiceJoinChannelClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChittyChatService_JoinChannelClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type chittyChatServiceJoinChannelClient struct {
	grpc.ClientStream
}

func (x *chittyChatServiceJoinChannelClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chittyChatServiceClient) LeaveChannel(ctx context.Context, in *Channel, opts ...grpc.CallOption) (*LeaveAck, error) {
	out := new(LeaveAck)
	err := c.cc.Invoke(ctx, "/chat.ChittyChatService/LeaveChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chittyChatServiceClient) SendMessage(ctx context.Context, opts ...grpc.CallOption) (ChittyChatService_SendMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChittyChatService_ServiceDesc.Streams[1], "/chat.ChittyChatService/SendMessage", opts...)
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
	JoinChannel(*Channel, ChittyChatService_JoinChannelServer) error
	LeaveChannel(context.Context, *Channel) (*LeaveAck, error)
	SendMessage(ChittyChatService_SendMessageServer) error
	mustEmbedUnimplementedChittyChatServiceServer()
}

// UnimplementedChittyChatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChittyChatServiceServer struct {
}

func (UnimplementedChittyChatServiceServer) JoinChannel(*Channel, ChittyChatService_JoinChannelServer) error {
	return status.Errorf(codes.Unimplemented, "method JoinChannel not implemented")
}
func (UnimplementedChittyChatServiceServer) LeaveChannel(context.Context, *Channel) (*LeaveAck, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveChannel not implemented")
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

func _ChittyChatService_JoinChannel_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Channel)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChittyChatServiceServer).JoinChannel(m, &chittyChatServiceJoinChannelServer{stream})
}

type ChittyChatService_JoinChannelServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type chittyChatServiceJoinChannelServer struct {
	grpc.ServerStream
}

func (x *chittyChatServiceJoinChannelServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func _ChittyChatService_LeaveChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Channel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChittyChatServiceServer).LeaveChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChittyChatService/LeaveChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChittyChatServiceServer).LeaveChannel(ctx, req.(*Channel))
	}
	return interceptor(ctx, in, info, handler)
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
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LeaveChannel",
			Handler:    _ChittyChatService_LeaveChannel_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "JoinChannel",
			Handler:       _ChittyChatService_JoinChannel_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SendMessage",
			Handler:       _ChittyChatService_SendMessage_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "chat/chat.proto",
}
