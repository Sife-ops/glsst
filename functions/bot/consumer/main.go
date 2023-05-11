package main

import (
	"fmt"
	"glsst/functions/lib"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	// "github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	// "github.com/davecgh/go-spew/spew"
)

func Handler(request events.APIGatewayProxyRequest) error {
	fmt.Println("consumer")

	u := lib.User{
		UserId:        "toad",
		Username:      "bbb",
		Discriminator: "ccc",
		DisplayName:   "ddd",
		GlobalName:    "eee",
		Avatar:        "fff",
	}
	fmt.Println(u)

	o := lib.Put(u)
	fmt.Println(o)

	return nil
}

func main() {
	lambda.Start(Handler)
}
