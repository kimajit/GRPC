package main

import (
	"context"
	"log"
	"time"

	pb "userdata/data"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewUserDataClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Test GetData method
	r, err := c.GetData(ctx, &pb.GetDataRequest{Id: 1})
	if err != nil {
		log.Fatalf("could not get user data: %v", err)
	}
	log.Printf("User Data: %v", r.User)

	// Test SearchAllData method
	searchResp, err := c.SearchAllData(ctx, &pb.SearchAllDataRequest{City: "LA", Phone: 1234567890, Married: true})
	if err != nil {
		log.Fatalf("could not search user data: %v", err)
	}
	log.Printf("Search Results: %v", searchResp.Users)
}
