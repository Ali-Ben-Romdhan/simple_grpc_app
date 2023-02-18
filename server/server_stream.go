package main

import (
	"log"
	"time"
	"github.com/alibnr/grpc/proto"
)

func (s *helloServer) SayHelloFromServerStreaming(req *proto.NamesList,stream proto.GreetService_SayHelloFromServerStreamingServer) error{
	log.Printf("Send request with names: %v", req.Names)

	for _,name := range req.Names{
		res:= &proto.HelloResponse{
			Message : "Hello "+name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
return nil
}