package main

import (
	"encoding/json"
	"glsst/functions/bot/command"
	"glsst/functions/lib"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) error {
	var b lib.InteractionBody
	if err := json.Unmarshal([]byte(request.Body), &b); err != nil {
		return err
	}

	// onboard user
	// todo: compare and update
	ul, err := lib.Query(lib.User{
		UserId: b.Member.User.UserId,
	})
	if err != nil {
		panic(err)
	}
	if len(ul.Items) < 1 {
		if _, err := lib.Put(b.Member.User); err != nil {
			panic(err)
		}
	}

	switch b.Data.Name {
	case "foo":
		return command.Foo(b)
	case "create":
		return command.Create(b)
	case "vote":
		return command.Vote(b)
	case "user":
		return command.User(b)
	}

	return nil
}

func main() {
	lambda.Start(Handler)
}
