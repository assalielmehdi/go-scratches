package main

import (
	"context"
	"log"

	pb "assalielmehdi/scratchs/grpc/calculator/proto"
)

func calculate(operator1, operator2 float32, operation pb.CalculatorRequest_OperationType) float32 {
	switch operation {
	case pb.CalculatorRequest_ADD:
		return operator1 + operator2
	case pb.CalculatorRequest_SUB:
		return operator1 - operator2
	case pb.CalculatorRequest_MUL:
		return operator1 * operator2
	case pb.CalculatorRequest_DIV:
		return operator1 / operator2
	default:
		return -1
	}
}

func (srv *Server) Calculate(ctx context.Context, req *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	log.Printf("Calculate was invoked with %v\n", req)

	return &pb.CalculatorResponse{Result: calculate(req.Operator1, req.Operator2, req.Type)}, nil
}
