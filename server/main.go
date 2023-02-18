package main

import (
	"log"
	"net"

	proto "github.com/alibnr/grpc/proto"
	"google.golang.org/grpc"
)
const (port =  ":8080")

type helloServer struct{
	proto.GreetServiceServer
}


func main() {

listen,err := net.Listen("tcp",port)
if err != nil {
	log.Fatalf("Failed to start the server %v",err)
}

grpcServer := grpc.NewServer()
proto.RegisterGreetServiceServer(grpcServer,&helloServer{})
log.Printf("Server listening on port %v",listen.Addr())
if err := grpcServer.Serve(listen); err!= nil {
	log.Fatalf("Failed to start %v",err)
}


}