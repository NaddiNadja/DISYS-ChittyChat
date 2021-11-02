package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"

	chat "github.com/NaddiNadja/DISYS-ChittyChat/Chat"
	"google.golang.org/grpc"
)

var channelName = flag.String("channel", "default", "Channel name for chatting")
var senderName = flag.String("sender", "default", "Senders name")
var tcpServer = flag.String("server", ":9080", "Tcp server")
var lamportTime = flag.Int64("time", 0, "lamportTimeStamp")

func main() {

	flag.Parse()

	fmt.Println("--- CHITTY CHAT ---")
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithBlock(), grpc.WithInsecure())

	conn, err := grpc.Dial(*tcpServer, opts...)
	if err != nil {
		log.Fatalf("Fail to dail: %v", err)
	}

	defer conn.Close()

	ctx := context.Background()
	client := chat.NewChittyChatServiceClient(conn)

	setUpClosehandler(ctx, client)

	go joinChannel(ctx, client)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		go sendMessage(ctx, client, scanner.Text())
	}

}

func joinChannel(ctx context.Context, client chat.ChittyChatServiceClient) {

	channel := chat.Channel{Name: *channelName, SendersName: *senderName}
	stream, err := client.JoinChannel(ctx, &channel)
	if err != nil {
		log.Fatalf("client.JoinChannel(ctx, &channel) throws: %v", err)
	}
	sendMessage(ctx, client, "This user just joined.")

	waitc := make(chan struct{})

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

func setUpClosehandler(ctx context.Context, client chat.ChittyChatServiceClient) {
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-ch
		sendMessage(ctx, client, "This user left the chat.")
		os.Exit(1)
	}()
}

func sendMessage(ctx context.Context, client chat.ChittyChatServiceClient, message string) {
	*lamportTime++
	stream, err := client.SendMessage(ctx)
	if err != nil {
		log.Printf("Cannot send message: error: %v", err)
	}
	msg := chat.Message{
		Channel: &chat.Channel{
			Name:        *channelName,
			SendersName: *senderName},
		Message:     message,
		Sender:      *senderName,
		LamportTime: *lamportTime,
	}
	stream.Send(&msg)
}
