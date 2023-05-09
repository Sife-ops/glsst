package main

import (
	"fmt"
	"glsst/functions/lib"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) error {
	fmt.Println("consumer")

	u := lib.User{
		UserId:        "ree",
		Username:      "bbb",
		Discriminator: "ccc",
		DisplayName:   "ddd",
		GlobalName:    "eee",
		Avatar:        "fff",
	}

	out := u.Put()
	fmt.Println(out)

	return nil
}

func main() {
	lambda.Start(Handler)
}
