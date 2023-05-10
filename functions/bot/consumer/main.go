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
		UserId:        "squeex",
		Username:      "bbb",
		Discriminator: "ccc",
		// DisplayName:   "ddd",
		GlobalName: "eee",
		Avatar:     "fff",
	}

	o := lib.Put(u)
	// fmt.Println(u)

	// out := lib.I2m(u)
	// out := lib.Pe(u)

	fmt.Println(o)

	return nil
}

func main() {
	lambda.Start(Handler)
}
