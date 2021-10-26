package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"time"
	"math/rand"
	"github.com/NaddiNadja/DISYS-ChittyChat/Chat"
	"google.golang.org/grpc"
	"Chat/chat"
)

type messageUnit struct {
	ClientName        string
	MessageBody       string
	MessageUniqueCode int
	ClientUniqueCode  int
}

type messageHandle struct {
	MQue []messageUnit
	mu   sync.Mutex
}

var messageHandleObject = messageHandle{}

type ChatServer struct {

}

func (is *ChatServer) ChatService(csi Services_ChittyChatServiceClient) {

	clientUniqueCode := rand.Intn(1e6)

	// receive messages - init a go routine
	go receiveFromStream(csi, clientUniqueCode)

	// send messages - init a go routine
	go sendToStream(csi, clientUniqueCode)
}

func receiveFromStream(csi_ Services_ChittyChatServiceClient, clientUniqueCode_ int) {

	//implement a loop
	for {
		mssg, err := csi_.Recv()
		if err != nil {
			log.Printf("Error in receiving message from client :: %v", err)
		} else {

			messageHandleObject.mu.Lock()

			messageHandleObject.MQue = append(messageHandleObject.MQue, messageUnit{
				ClientName:        mssg.Name,
				MessageBody:       mssg.Body,
				MessageUniqueCode: rand.Intn(1e8),
				ClientUniqueCode:  clientUniqueCode_,
			})
			
			log.Printf("%v", messageHandleObject.MQue[len(messageHandleObject.MQue)-1])
			
			messageHandleObject.mu.Unlock()
		}
	}
}

func sendToStream(csi_ Services_ChittyChatServiceClient, clientUniqueCode_ int) {

	//implement a loop
	for {

		//loop through messages in MQue
		for {

			time.Sleep(500 * time.Millisecond)

			messageHandleObject.mu.Lock()

			if len(messageHandleObject.MQue) == 0 {
				messageHandleObject.mu.Unlock()
				break
			}

			senderUniqueCode := messageHandleObject.MQue[0].ClientUniqueCode
			senderName4Client := messageHandleObject.MQue[0].ClientName
			message4Client := messageHandleObject.MQue[0].MessageBody

			messageHandleObject.mu.Unlock()

			//send message to designated client (do not send to the same client)
			if senderUniqueCode != clientUniqueCode_ {

				err := csi_.Send(&FromServer{Name: senderName4Client, Body: message4Client})

				if err != nil {
					log.Printf("Error in sending message from client :: %v", err)
				}

				messageHandleObject.mu.Lock()

				if len(messageHandleObject.MQue) > 1 {
					messageHandleObject.MQue = messageHandleObject.MQue[1:]
				} else {
					messageHandleObject.MQue = []messageUnit{}
				}

				messageHandleObject.mu.Unlock()

			}

		}

		time.Sleep(100 * time.Millisecond)
	}
}
