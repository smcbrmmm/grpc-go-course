package main

import (
	"context"
	"log"

	pb "github.com/smcbrmmm/grpc-go-course/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doSqrt(c pb.CalculatorServiceClient, n int32) {
	log.Println("doSqrt was invoked")
	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{
		Number: n,
	})

	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			log.Printf("Error message from server: %s\n", e.Message())
			log.Printf("Error code from server: %s\n", e.Code())

			if e.Code() == codes.InvalidArgument {
				log.Printf("We probably send a negative number!")
				return
			}

		} else {
			log.Fatalf("A non gRPC err: %v\n", err)
		}
	}

	log.Printf("Sqrt: %.2f\n", res.Result)
}
