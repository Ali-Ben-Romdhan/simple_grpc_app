package main

import (
	"log"
	proto "github.com/alibnr/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)
const (port =  ":8080")

func main() {

conn,err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))

if err != nil {
	log.Fatalf("did not connect: %v", err)

}
defer conn.Close()
client := proto.NewGreetServiceClient(conn)

names := &proto.NamesList{
	Names: []string{"Ali BNR","Adem BC","Akrem CH","Mouez MT","Nicolas MG1","Nicolas MG2","Nicolas MG3"},
}
//callSayHello(client)
//sayHelloFromServerStream(client,names)
//callSayHelloClientStream(client,names)
callSayHelloBidirectionalStream(client,names)
}