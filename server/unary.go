package main

import (
	"context"

	proto "github.com/alibnr/grpc/proto"
)




func (s *helloServer) SayHello(ctx context.Context,req *proto.NoParam)(*proto.HelloResponse,error){
	return &proto.HelloResponse{
		Message: "Hello",
	},nil
}