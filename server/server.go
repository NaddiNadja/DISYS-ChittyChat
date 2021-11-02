package main

import (
	"io"
	"log"
	"net"

	chat "github.com/NaddiNadja/DISYS-ChittyChat/Chat"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	chat.RegisterChittyChatServiceServer(grpcServer, &chittyChatServiceServer{
		rooms: make(map[string][]chan *chat.Message),
	})
	//grpcServer.Serve(lis)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to server %v", err)
	}
}

type chittyChatServiceServer struct {
	chat.UnimplementedChittyChatServiceServer
	rooms map[string][]chan *chat.Message
}

func (s *chittyChatServiceServer) JoinRoom(ch *chat.Room, msgStream chat.ChittyChatService_JoinRoomServer) error {

	msgChannel := make(chan *chat.Message)
	s.rooms[ch.Name] = append(s.rooms[ch.Name], msgChannel)

	log.Printf("Client \"%v\" joined", ch.SendersName)

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

func (s *chittyChatServiceServer) LeaveRoom(ch *chat.Room, msgStream chat.ChittyChatService_LeaveRoomServer) error {
	log.Printf("Client \"%v\" left", ch.SendersName)
	return nil
}

func (s *chittyChatServiceServer) SendMessage(msgStream chat.ChittyChatService_SendMessageServer) error {
	msg, err := msgStream.Recv()

	if err == io.EOF {
		return nil
	}

	if err != nil {
		return err
	}

	log.Printf("Message received from sender \"%v\"", msg.Sender)
	ack := chat.MessageAck{MessageAck: "SENT"}
	msgStream.SendAndClose(&ack)

	go func() {
		streams := s.rooms[msg.Room.Name]
		for i := 0; i < len(streams); i++ {
			select {
			case streams[i] <- msg:
			default:
				streams[i] = streams[len(streams)-1]
				streams[len(streams)-1] = nil
				streams = streams[:len(streams)-1]
			}
		}
	}()

	return nil
}
