package main

import (
	"context"
	"log"

	pb "github.com/smcbrmmm/grpc-go-course/blog/proto"
)

func CreateBlog(c pb.BlogServiceClient) string {
	log.Println("---- createBlog was invoked ----")

	blog := &pb.Blog{
		AuthorId: "Samut",
		Title:    "Harry Potter",
		Content:  "Very very small",
	}

	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog has been created: %s\n", res.Id)

	return res.Id
}
