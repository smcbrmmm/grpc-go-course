package main

import (
	"context"
	"log"
	"time"

	pb "github.com/smcbrmmm/grpc-go-course/calculator/proto"
)

func doAverage(c pb.CalculatorServiceClient) {
	log.Println("doLongGreet was invoked")

	reqs := []*pb.AverageRequest{
		{Number: 1},
		{Number: 2},
		{Number: 3},
		{Number: 4},
	}

	stream, err := c.Average(context.Background())

	if err != nil {
		log.Fatalf("Error while calling Average %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)

		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receving response from Average: %v\n", err)
	}

	log.Printf("Average: %.1f\n", res.Result)
}
