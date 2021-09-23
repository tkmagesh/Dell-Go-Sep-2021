package main

import (
	"context"
	"grpc-app/proto"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedAppServiceServer
}

//request & respone
func (s *server) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	x := req.GetX()
	y := req.GetY()
	result := x + y
	log.Printf("Add: %d + %d = %d\n", x, y, result)
	res := &proto.AddResponse{
		Result: result,
	}
	return res, nil
}

//client streaming
func (s *server) Average(stream proto.AppService_AverageServer) error {
	var sum, count int64
	log.Println("Average: start")
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			avg := sum / count
			avgResponse := &proto.AverageResponse{
				Result: avg,
			}
			log.Println("Sending response : avg = ", avg)
			return stream.SendAndClose(avgResponse)
		}
		if err != nil {
			return err
		}
		log.Println("Received : ", req.GetNum())
		sum += req.GetNum()
		count++
	}
	return nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer listener.Close()
	grpcServer := grpc.NewServer()
	proto.RegisterAppServiceServer(grpcServer, &server{})
	e := grpcServer.Serve(listener)
	if e != nil {
		log.Fatalf("failed to serve: %v", e)
	}
}
