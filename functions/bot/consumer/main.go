package main

import (
	"encoding/json"
	"glsst/functions/bot/command"
	"glsst/functions/lib"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) error {
	var interactionBody lib.InteractionBody
	if err := json.Unmarshal([]byte(request.Body), &interactionBody); err != nil {
		return err
	}

	switch interactionBody.Data.Name {
	case "foo":
		return command.Foo(interactionBody)
	}

	// u := lib.User{
	// 	UserId:        "8732894082384932948239",
	// 	Username:      "wyatt",
	// 	Discriminator: "1234",
	// 	// DisplayName:   "ddd",
	// 	// GlobalName:    "eee",
	// 	Avatar: "89379472974328904823",
	// }
	// fmt.Println(u)

	// o := lib.Put(u)
	// fmt.Println(o)

	// qq := lib.Query(lib.User{
	// 	UserId: "8732894082384932948239",
	// })
	// var uu []lib.User
	// err := attributevalue.UnmarshalListOfMaps(qq.Items, &uu)
	// if err != nil {
	// 	panic(err)
	// }
	// spew.Dump(uu)

	return nil
}

func main() {
	lambda.Start(Handler)
}
