package main

import (
	"encoding/json"
	"glsst/functions/bot/command"
	"glsst/functions/lib"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) error {
	var ib lib.InteractionBody
	if err := json.Unmarshal([]byte(request.Body), &ib); err != nil {
		return err
	}

	switch ib.Data.Name {
	case "foo":
		return command.Foo(ib)
	case "create":
		return command.Create(ib)
	case "vote":
		return command.Vote(ib)
	case "info":
		return command.Info(ib)
	}

	return nil
}

func main() {
	lambda.Start(Handler)
}
