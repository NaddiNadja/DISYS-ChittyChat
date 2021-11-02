package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	chat "github.com/NaddiNadja/DISYS-ChittyChat/Chat"
	"google.golang.org/grpc"

	goodbye "github.com/thecodeteam/goodbye"
)

var roomName = flag.String("room", "default", "Chat room for chatting")
var senderName = flag.String("sender", "default", "Senders name")
var tcpServer = flag.String("server", ":9080", "Tcp server")
var lamportTime = flag.Int64("time", 0, "lamportTimeStamp")

func main() {

	flag.Parse()

	fmt.Println("--- CHITTY CHAT ---\n - CTRL + C to leave \n - All bad words allowed")
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithBlock(), grpc.WithInsecure())

	conn, err := grpc.Dial(*tcpServer, opts...)
	if err != nil {
		log.Fatalf("Fail to dail: %v", err)
	}

	defer conn.Close()

	ctx := context.Background()
	client := chat.NewChittyChatServiceClient(conn)
	sendMessage(ctx, client, "This user just joined.")

	defer goodbye.Exit(ctx, -1)

	goodbye.Notify(context.Background())

	goodbye.RegisterWithPriority(func(ctx context.Context, sig os.Signal) {

		if !goodbye.IsNormalExit(sig) {
			sendMessage(ctx, client, "This user just left in the hardcore way")
		}

	}, -1)

	go joinChannel(ctx, client)
	go leaveChannel(ctx, client)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		go sendMessage(ctx, client, scanner.Text())
	}

}

func joinChannel(ctx context.Context, client chat.ChittyChatServiceClient) {

	channel := chat.Room{Name: *roomName, SendersName: *senderName}
	stream, err := client.JoinRoom(ctx, &channel)
	if err != nil {
		log.Fatalf("client.JoinChannel(ctx, &channel) throws: %v", err)
	}
	sendMessage(ctx, client, "This user just joined.")

	waitc := make(chan struct{}) //go never stops with this

	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive message from channel joining. \nErr: %v", err)
			}
			if *senderName != in.Sender {
				if in.LamportTime > *lamportTime {
					*lamportTime = in.LamportTime + 1
				} else {
					*lamportTime++
				}
				fmt.Printf("(%v) %v: %v \n", *lamportTime, in.Sender, in.Message)
			}
		}
	}()
	<-waitc
}

func leaveChannel(ctx context.Context, client chat.ChittyChatServiceClient) {

	channel := chat.Room{Name: *roomName, SendersName: *senderName}
	stream, err := client.LeaveRoom(ctx, &channel)
	if err != nil {
		log.Fatalf("client.LeaveRoom(ctx, &channel) throws: %v", err)
	}

	waitc := make(chan struct{})

	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive message from leaving channel. \nErr: %v", err)
			}
			if *senderName != in.Sender {
				if in.LamportTime > *lamportTime {
					*lamportTime = in.LamportTime + 1
				} else {
					*lamportTime++
				}
				fmt.Printf("(%v) %v: %v \n", *lamportTime, in.Sender, in.Message)
			}
		}
	}()
	<-waitc
}

func sendMessage(ctx context.Context, client chat.ChittyChatServiceClient, message string) {
	*lamportTime++
	stream, err := client.SendMessage(ctx)
	if err != nil {
		log.Printf("Cannot send message: error: %v", err)
	}
	msg := chat.Message{
		Room: &chat.Room{
			Name:        *roomName,
			SendersName: *senderName},
		Message:     message,
		Sender:      *senderName,
		LamportTime: *lamportTime,
	}
	stream.Send(&msg)
}
