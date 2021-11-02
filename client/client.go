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

	"github.com/thecodeteam/goodbye"
)

var roomName = flag.String("room", "default", "Chat room for chatting")
var senderName = flag.String("sender", "default", "Senders name")
var tcpServer = ":9080"
var lamportTime = int64(0)

func main() {

	flag.Parse()

	fmt.Print("╒═════ CHITTY CHAT ═════╕\n│ CTRL + C to leave     │\n│ All bad words allowed │\n└───────────────────────┘\n\n")
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithBlock(), grpc.WithInsecure())

	conn, err := grpc.Dial(tcpServer, opts...)
	if err != nil {
		log.Fatalf("Fail to dail: %v", err)
	}

	ctx := context.Background()
	client := chat.NewChittyChatServiceClient(conn)

	ctx2 := context.Background()
	defer goodbye.Exit(ctx2, -1)
	goodbye.Notify(ctx2)
	goodbye.RegisterWithPriority(func(ctx context.Context, sig os.Signal) {
		leaveRoom(ctx, client)
		conn.Close()
	}, 0)

	go joinRoom(ctx, client)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		go sendMessage(ctx, client, scanner.Text())
	}

}

func joinRoom(ctx context.Context, client chat.ChittyChatServiceClient) {

	room := chat.Room{Name: *roomName, SendersName: *senderName}
	stream, err := client.JoinRoom(ctx, &room)
	if err != nil {
		log.Fatalf("client.JoinRoom(ctx, &room) throws: %v", err)
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
				if in.LamportTime > lamportTime {
					lamportTime = in.LamportTime + 1
				} else {
					lamportTime++
				}
				fmt.Printf("(%v) %v: %v \n", lamportTime, in.Sender, in.Message)
			}
		}
	}()
	<-waitc

}

func leaveRoom(ctx context.Context, client chat.ChittyChatServiceClient) {
	sendMessage(ctx, client, "just left in the hardcore way")
	log.Println("Ses b")
	room := chat.Room{Name: *roomName, SendersName: *senderName}
	client.LeaveRoom(ctx, &room)
}

func sendMessage(ctx context.Context, client chat.ChittyChatServiceClient, message string) {
	lamportTime++
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
		LamportTime: lamportTime,
	}
	stream.Send(&msg)
}
