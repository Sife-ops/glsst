package main

import (
	"fmt"
	"glsst/functions/lib"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/davecgh/go-spew/spew"
)

func Handler(request events.APIGatewayProxyRequest) error {
	fmt.Println("consumer")

	u := lib.User{
		UserId:        "8732894082384932948239",
		Username:      "wyatt",
		Discriminator: "1234",
		// DisplayName:   "ddd",
		// GlobalName:    "eee",
		Avatar: "89379472974328904823",
	}
	fmt.Println(u)

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

	p := lib.Prediction{
		PredictionId: "112",
		UserId:       "8732894082384932948239",
		Condition:    "nome",
		CreatedAt:    "today",
	}
	fmt.Println(p)

	// o := lib.Put(p)
	// fmt.Println(o)

	qq := lib.Query(lib.Prediction{
		PredictionId: "111",
		UserId:       "8732894082384932948239",
	}, lib.Gsi1)
	var uu []lib.Prediction
	err := attributevalue.UnmarshalListOfMaps(qq.Items, &uu)
	if err != nil {
		panic(err)
	}
	spew.Dump(uu)

	return nil
}

func main() {
	lambda.Start(Handler)
}
