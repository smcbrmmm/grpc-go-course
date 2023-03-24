package main

import (
	"context"
	"io"
	"log"

	pb "github.com/smcbrmmm/grpc-go-course/calculator/proto"
)

func doPrime(cal pb.CalculatorServiceClient) {
	log.Println("Prime was invoked")

	req := pb.PrimeRequest{
		Request: 120,
	}

	stream, err := cal.Prime(context.Background(), &req)

	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("Prime number: %s\n", msg)
	}
}
