package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "assalielmehdi/scratchs/grpc/calculator/proto"
)

var addr string = "localhost:50051"

func main() {
	con, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer con.Close()

	c := pb.NewCalculatorServiceClient(con)

	Calculate(c)
}
