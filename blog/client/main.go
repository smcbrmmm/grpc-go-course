package main

import (
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/smcbrmmm/grpc-go-course/blog/proto"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close() // execute at the end of the function

	c := pb.NewBlogServiceClient(conn)

	id := CreateBlog(c)

	ReadBlog(c, id) // Valid
	time.Sleep(2 * time.Second)
	ReadBlog(c, "invalid id")
}
