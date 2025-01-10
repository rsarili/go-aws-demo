package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Parse the input event

	fmt.Println("request body:", request.Body)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Request body is: " + request.Body,
	}, nil
}

func main() {
	lambda.Start(handleRequest)
}
