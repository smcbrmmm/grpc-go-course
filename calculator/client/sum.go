package main

import (
	"context"
	"log"

	pb "github.com/smcbrmmm/grpc-go-course/calculator/proto"
)

func doSum(cal pb.CalculatorServiceClient) {
	log.Println("Sum was invoked")

	res, err := cal.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  3,
		SecondNumber: 10,
	})

	if err != nil {
		log.Fatalf("Could not sum: %v\n", err)
	}

	log.Printf("Sum value: %d\n", res.Result)
}
