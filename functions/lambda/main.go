package main

import (
	"bytes"
	"crypto/ed25519"
	"encoding/hex"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	key, err := hex.DecodeString("")
	if err != nil {
		panic("k")
	}

	var msg bytes.Buffer

	signature := request.Headers["x-signature-ed25519"]
	if signature == "" {
		panic("l")
	}

	sig, err := hex.DecodeString(signature)
	if err != nil {
		panic(err)
	}

	if len(sig) != 64 {
		panic("s")
	}

	timestamp := request.Headers["x-signature-timestamp"]
	if timestamp == "" {
		// return false
		panic("t")
	}

	msg.WriteString(timestamp)
	msg.WriteString(request.Body)

	result := ed25519.Verify(key, msg.Bytes(), sig)
	fmt.Println(result)

	return events.APIGatewayProxyResponse{
		Body:       "Hello, World! Your yddidtiyti request was received at OK? " + request.RequestContext.Time + ".",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
