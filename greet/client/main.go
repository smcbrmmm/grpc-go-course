package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/smcbrmmm/grpc-go-course/greet/proto"
)

var addr string = "localhost:50051"

func main() {
	tls := true // change that to false if needed
	opts := []grpc.DialOption{}

	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")

		if err != nil {
			log.Fatalf("Error while loading CA trust certificates: %v\n", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))

	}

	conn, err := grpc.Dial(addr, opts...)

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close() // execute at the end of the function

	c := pb.NewGreetServiceClient(conn)

	doGreet(c)

	// doGreetManyTimes(c)

	// doLongGreet(c)

	// doGreetEveryone(c)

	// doGreetWithDeadline(c, 5*time.Second)

	// doGreetWithDeadline(c, 1*time.Second)
}
