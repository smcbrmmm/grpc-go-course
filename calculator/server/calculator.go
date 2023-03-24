package main

import (
	"context"
	"io"
	"log"

	pb "github.com/smcbrmmm/grpc-go-course/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum function was invoked with %v\n", in)
	return &pb.SumResponse{
		Result: in.FirstNumber + in.SecondNumber,
	}, nil
}

func (s *Server) Prime(in *pb.PrimeRequest, stream pb.CalculatorService_PrimeServer) error {
	log.Printf("Prime function was invoke with %v\n", in)

	var k int32 = 2
	N := in.Request
	for N > 1 {
		if N%k == 0 {
			stream.Send(&pb.PrimeResponse{
				Result: k,
			})
			N = N / k
		} else {
			k = k + 1
		}
	}

	return nil
}

func (s *Server) Average(stream pb.CalculatorService_AverageServer) error {
	log.Println("Average function was invoked")

	var res float32
	var reqNo float32
	var result float32

	for {
		req, err := stream.Recv()

		if err == io.EOF {

			result = res / reqNo

			return stream.SendAndClose(&pb.AverageResponse{
				Result: result,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		log.Printf("Receving: %v\n", req)
		res += req.Number
		reqNo++
	}
}
