package main

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/weinandt/big-map/bigmap"
)

type MyEvent struct {
	Name string `json:"name"`
}

func generate500MBRandomString() string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, 500*1024*1024)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func HandleRequest(ctx context.Context, name MyEvent) error {
	bigString := generate500MBRandomString()
	testMap, err := bigmap.NewBigMap(16 * 1024 * 1024 * 1024)
	if err != nil {
		return err
	}

	for i := 0; i < 16*2; i++ {
		err := testMap.Add(fmt.Sprint(i), bigString)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	lambda.Start(HandleRequest)
}
