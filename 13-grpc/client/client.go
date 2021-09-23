package main

import (
	"context"
	"io"
	"log"
	"time"

	"grpc-app/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	//doServerStreaming(clientConn)
	//doBiDiStreaming(clientConn)
	doRequestResponseWithTimeout(clientConn)
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

func doBiDiStreaming(clientConn proto.AppServiceClient) {
	data := []*proto.Greeting{
		&proto.Greeting{
			FirstName: "Magesh",
			LastName:  "Kuppan",
		},
		&proto.Greeting{
			FirstName: "Ramesh",
			LastName:  "Jayaraman",
		},
		&proto.Greeting{
			FirstName: "Suresh",
			LastName:  "Kannan",
		},
		&proto.Greeting{
			FirstName: "Rajesh",
			LastName:  "Pandit",
		},
		&proto.Greeting{
			FirstName: "Dinesh",
			LastName:  "Kumar",
		},
	}
	stream, err := clientConn.Greet(context.Background())
	if err != nil {
		log.Fatalln(err)
	}
	go func() {
		for _, greeting := range data {
			time.Sleep(400 * time.Millisecond)
			log.Println("Sending request : ", greeting.FirstName)
			req := &proto.GreetRequest{
				Greeting: greeting,
			}
			if err := stream.Send(req); err != nil {
				log.Fatalln(err)
			}
		}
		stream.CloseSend()
	}()
	done := make(chan bool)
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				log.Println("All greetings received")
				done <- true
				return
			}
			if err != nil {
				log.Fatalln(err)
			}
			log.Println("Greeting received : ", res.GetMessage())
		}
	}()
	<-done
}

func doRequestResponseWithTimeout(client proto.AppServiceClient) {
	req := &proto.AddRequest{X: 100, Y: 200}
	timeoutContext, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	resp, err := client.Add(timeoutContext, req)
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				log.Fatalln("Timeout was hit! Deadline Exceeded")
			} else {
				log.Fatalln(err)
			}
		}
	} else {
		log.Println("Add Result:", resp.GetResult())
	}
}
