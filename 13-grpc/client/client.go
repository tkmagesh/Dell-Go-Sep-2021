package main

import (
	"context"
	"log"

	"grpc-app/proto"

	"google.golang.org/grpc"
)

func main() {
	client, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()
	clientConn := proto.NewAppServiceClient(client)
	req := &proto.AddRequest{
		X: 100,
		Y: 200,
	}
	res, err := clientConn.Add(context.Background(), req)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res.Result)
}
