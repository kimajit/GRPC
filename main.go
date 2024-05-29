package main

import (
	"log"
	"net"

	pb "userdata/data"
	"userdata/server"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserDataServer(s, &server.Server{
		Users: map[int32]*pb.Data{
			1: {Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
			2: {Id: 2, Fname: "John", City: "NY", Phone: 9876543210, Height: 6.0, Married: false},
		},
	})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
