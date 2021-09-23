package main

import (
	"context"
	"io"
	"log"
	"time"

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
	//doRequestResponse(clientConn)
	//doClientStreaming(clientConn)
	doServerStreaming(clientConn)
}

func doRequestResponse(clientConn proto.AppServiceClient) {
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

func doClientStreaming(clientConn proto.AppServiceClient) {
	data := []int64{30, 20, 50, 10, 40}

	log.Println("Client Streaming initiated")
	clientStream, err := clientConn.Average(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	for _, no := range data {
		time.Sleep(300 * time.Millisecond)
		log.Println("Sending : ", no)
		avgReq := &proto.AverageRequest{
			Num: no,
		}
		e := clientStream.Send(avgReq)
		if e != nil {
			log.Fatalln(e)
		}
	}
	res, e := clientStream.CloseAndRecv()
	if e != nil {
		log.Fatalln(e)
	}
	log.Println("Average : ", res.Result)
}

func doServerStreaming(clientConn proto.AppServiceClient) {
	req := &proto.PrimeRequest{
		Start: 3,
		End:   100,
	}
	stream, err := clientConn.Prime(context.Background(), req)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			log.Println("All prime numbers received")
			return
		}
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("Prime number received : ", res.GetResult())
	}
}
