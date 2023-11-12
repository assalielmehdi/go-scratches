package main

import (
	pb "assalielmehdi/scratches/grpc/calculator/proto"
	"context"
	"log"
)

func Calculate(c pb.CalculatorServiceClient) {
	res, err := c.Calculate(context.Background(), &pb.CalculatorRequest{
		Operator1: 10,
		Operator2: 5,
		Type:      pb.CalculatorRequest_ADD,
	})

	if err != nil {
		log.Fatalf("Could not calculate: %v\n", err)
	}

	log.Printf("10 + 5 = %v", res.Result)
}
