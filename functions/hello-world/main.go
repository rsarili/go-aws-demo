package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

func handleRequest(ctx context.Context, event json.RawMessage) (string, error) {
	// Parse the input event
	var message map[string]interface{}
	if err := json.Unmarshal(event, &message); err != nil {
		log.Printf("Failed to unmarshal event: %v", err)
		return "", err
	}

	fmt.Println("Received event:", message)

	return message["name"].(string), nil
}

func main() {
	lambda.Start(handleRequest)
}
