package main

import (
	"context"
	"fmt"
	"log"
	"github.com/NaddiNadja/DISYS-ChittyChat/Chat"
	"google.golang.org/grpc"
)

func main() {
	// Creat a virtual RPC Client Connection on port  9080 WithInsecure (because  of http)
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %s", err)
	}

	// Defer means: When this function returns, call this method (meaing, one main is done, close connection)
	defer conn.Close()

	//  Create new Client from generated gRPC code from proto
	c := chat.NewChittyChatClient(conn)

	sendJoinRequest(c)

	for {
		getMessages(c)
		if true {
			sendMessage(c)
		}
	}

}

func sendJoinRequest(c chat.ChittyChatClient) {
	//Send Request to join current session, return response with time.
	message := chat.JoinRequest{}

	response, err := c.Join(context.Background(), &message)
	
	if err != nil {
			log.Fatalf("Error when calling Join() : %s", err)
	}

	fmt.Printf("Successfully joined session, at time: %s", response.Reply)

}

func getMessages(c chat.ChittyChatClient) {
	//Send request to server for all messages AND print them.
	message := chat.MessagesRequest{}

	response, err := c.GetMessages(context.Background(), &message)

	if err != nil {
		log.Fatalf("Error when calling GetMessages() : %s", err)
	}
	

	var messageReplies = response.Replies

	for i := 0; i < len(messageReplies); i++ {
		fmt.Println(messageReplies[i])	
	}
	
}

func sendMessage(c chat.ChittyChatClient) {
	//Send post request message to server with text from Std.input adding to list of messages
	//this can later be retrieved via. the normal getMessages() method.

}