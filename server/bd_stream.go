package main

import (
	"io"
	"log"

	"github.com/alibnr/grpc/proto"
)

func (s *helloServer) SayHelloBidirectionalStreaming(stream proto.GreetService_SayHelloBidirectionalStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with name : %v", req.Name)
		res := &proto.HelloResponse{
			Message: "Good Morning " + req.Name + "from server streaming",
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
}