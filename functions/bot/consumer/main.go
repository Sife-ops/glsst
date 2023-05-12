package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"glsst/functions/lib"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	// "github.com/davecgh/go-spew/spew"
	"net/http"
)

func Handler(request events.APIGatewayProxyRequest) error {
	fmt.Println("consumer")

	var interactionBody lib.InteractionBody
	if err := json.Unmarshal([]byte(request.Body), &interactionBody); err != nil {
		return err
	}

	switch interactionBody.Data.Name {
	case "foo":
		fmt.Println("foo")

		r := lib.ResponseData{
			Content: "ree",
			// Flags: 64,
		}
		rj, err := json.Marshal(r)
		if err != nil {
			return err
		}

		url := fmt.Sprintf("https://discord.com/api/v10/webhooks/%s/%s", lib.GetBotAppId(), interactionBody.Token)
		req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(rj))
		if err != nil {
			return err
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", lib.GetBotPk())

		cl := &http.Client{}
		res, err := cl.Do(req)
		if err != nil {
			return err
		}
		fmt.Println(res)
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

	// p := lib.Prediction{
	// 	PredictionId: "112",
	// 	UserId:       "8732894082384932948239",
	// 	Condition:    "nome",
	// 	CreatedAt:    "today",
	// }
	// fmt.Println(p)

	// o := lib.Put(p)
	// fmt.Println(o)

	// qq := lib.Query(lib.Prediction{
	// 	PredictionId: "111",
	// 	UserId:       "8732894082384932948239",
	// }, lib.Gsi1)
	// var uu []lib.Prediction
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
