package main

import (
	 "context"
	 "io"
	"log"

	proto "github.com/alibnr/grpc/proto"
)




func sayHelloFromServerStream(client proto.GreetServiceClient,names *proto.NamesList){
	log.Printf("streaming started")

	stream, err := client.SayHelloFromServerStreaming(context.Background(), names)


	if err != nil {
		log.Fatalf("could not send names: %v",err)
	}

	for{

		message, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("error while streaming::: %v",err)

		}
		log.Println(message)
	}
	log.Println("Streaming finished")
}