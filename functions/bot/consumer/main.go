package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) (bool, error) {
	fmt.Println("ya")
	fmt.Println(request.Body)
	return true, nil
}

func main() {
	lambda.Start(Handler)
}
