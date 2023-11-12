package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "assalielmehdi/scratches/grpc/calculator/proto"
)

var addr string = "0.0.0.0:50051"

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	defer lis.Close()
	log.Printf("Listening on %s\n", addr)

	srv := grpc.NewServer()

	pb.RegisterCalculatorServiceServer(srv, &Server{})

	defer srv.Stop()
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
