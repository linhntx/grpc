package grpc_chatserver

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

type messageUnit struct {
	ClientName        string
	MessageBody       string
	MessageUniqueCode int
	ClientUniqueCode  int
}

type messageQueue struct {
	MQue []messageUnit
	mu   sync.Mutex
}

var messageQueueObject = messageQueue{}

type ChatServer struct {
	UnimplementedServicesServer
}

func(is *ChatServer) ChatService(csi Services_ChatServiceServer) error {

	clientUniqueCode := rand.Intn(1e6)
	errch := make(chan error)

	// recevice messages
	go receviceFromStream(csi, clientUniqueCode, errch)

	// send messages
	go sendToStream(csi, clientUniqueCode, errch)
	return <- errch
}

//recevice messages
func receviceFromStream(csi Services_ChatServiceServer, clientUniqueCode int, errch chan error) {
	
	//implement a loop 
	for {
		mssg, err := csi.Recv()
		if err != nil {
			log.Printf("Error when receive message from client :: %v", err)
			errch <- err
		} else {
			messageQueueObject.mu.Lock()

			messageQueueObject.MQue = append(messageQueueObject.MQue, messageUnit{
				ClientName: mssg.Name,
				MessageBody: mssg.Body,
				MessageUniqueCode: rand.Intn(1e8),
				ClientUniqueCode: clientUniqueCode,
			})

			log.Printf("%v", messageQueueObject.MQue[len(messageQueueObject.MQue)-1])

			messageQueueObject.mu.Unlock()
		}
	}
}

func sendToStream(csi Services_ChatServiceServer, clientUniqueCode int, errch chan error) {
	
	//implement a loop
	for {

		for {

			time.Sleep(500 * time.Millisecond)

			messageQueueObject.mu.Lock()

			if len(messageQueueObject.MQue) == 0 {
				messageQueueObject.mu.Unlock()
				break
			}

			senderUniqueCode := messageQueueObject.MQue[0].ClientUniqueCode
			senderNameToClient := messageQueueObject.MQue[0].ClientName
			messsageToClient := messageQueueObject.MQue[0].MessageBody

			messageQueueObject.mu.Unlock()

			if senderUniqueCode != clientUniqueCode {
				err := csi.Send(&FromServer{Name: senderNameToClient, Body: messsageToClient})

				if err != nil {
					errch <- err
				}

				messageQueueObject.mu.Lock()

				if len(messageQueueObject.MQue) > 1 {
					messageQueueObject.MQue = messageQueueObject.MQue[1:]
				} else {
					messageQueueObject.MQue = []messageUnit{}
				}

				messageQueueObject.mu.Unlock()
			}
		}

		time.Sleep(100 * time.Millisecond)
	}
}