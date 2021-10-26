package main

import (
	"github.com/NaddiNadja/DISYS-ChittyChat/chat"
	"io"
	"google.golang.org/grpc"
	"net"
	"log"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:9080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	chat.RegisterChittyChatServiceServer(grpcServer, &chittyChatServiceServer{
		channel: make(map[string][] chan *chat.Message),
	})
	grpcServer.Serve(lis)
}

type chittyChatServiceServer struct {
    chat.UnimplementedChittyChatServiceServer
    channel map[string][]chan *chat.Message
}

func (s *chittyChatServiceServer) JoinChannel(ch *chat.Channel, msgStream chat.ChittyChatService_JoinChannelServer) error {

	msgChannel := make(chan *chat.Message)
	s.channel[ch.Name] = append(s.channel[ch.Name], msgChannel)

	// doing this never closes the stream
	for {
		select {
		case <-msgStream.Context().Done():
			return nil
		case msg := <-msgChannel:
			msgStream.Send(msg)
		}
	}
}

func (s *chittyChatServiceServer) SendMessage(msgStream chat.ChittyChatService_SendMessageServer) error {
	msg, err := msgStream.Recv()

	if err == io.EOF {
		return nil
	}

	if err != nil {
		return err
	}

	ack := chat.MessageAck{MessageAck: "SENT"}
	msgStream.SendAndClose(&ack)

	go func() {
		streams := s.channel[msg.Channel.Name]
		for _, msgChan := range streams {
			msgChan <- msg
		}
	}()

	return nil
}