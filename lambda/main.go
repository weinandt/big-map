package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	bigmap "github.com/weinandt/big-map/bigmap"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) error {
	bigmap.Test()

	return nil
}

func main() {
	lambda.Start(HandleRequest)
}
