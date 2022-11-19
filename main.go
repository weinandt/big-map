package main

import (
	"context"
	"log"
	"syscall"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) error {
	bytes, err := syscall.Mmap(-1, 0, 1000000, syscall.PROT_WRITE|syscall.PROT_WRITE, syscall.MAP_ANON|syscall.MAP_SHARED)
	if err != nil {
		log.Fatal("Problem with mmap", err)
	}

	log.Println("length: ", len(bytes))
	log.Println(bytes[0])

	return nil
}

func main() {
	lambda.Start(HandleRequest)
}
