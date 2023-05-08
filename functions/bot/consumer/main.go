package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// func Handler(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
func Handler(request events.APIGatewayProxyRequest) (bool, error) {
	fmt.Println("ya")
	fmt.Println(request.Body)
	return true, nil
	// return events.APIGatewayProxyResponse{
	// 	Body:       "Hello, World! Your request was received at " + request.RequestContext.Time + ".",
	// 	StatusCode: 200,
	// }, nil
}

func main() {
	lambda.Start(Handler)
}
