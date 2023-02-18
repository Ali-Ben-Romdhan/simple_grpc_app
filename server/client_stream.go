package main

import (
	"io"
	"log"

	"github.com/alibnr/grpc/proto"
)

func (s *helloServer) SayHelloFromClientStreaming(stream proto.GreetService_SayHelloFromClientStreamingServer) error {
	var messages []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&proto.MessagesList{Messages: messages})
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with name : %v", req.Name)
		messages = append(messages, "Hello "+req.Name)
	}
}