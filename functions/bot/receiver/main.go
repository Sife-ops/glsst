package main

import (
	"encoding/json"
	"glsst/functions/lib"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	var ib lib.InteractionBody
	if err := json.Unmarshal([]byte(request.Body), &ib); err != nil {
		panic("unmarshal")
	}

	switch ib.Type {
	case 1:
		{
			switch lib.VerifyInteraction(request) {
			case true:
				{
					return events.APIGatewayProxyResponse{
						Body:       request.Body,
						StatusCode: 200,
					}, nil
				}
			default:
				{
					return events.APIGatewayProxyResponse{
						StatusCode: 401,
					}, nil
				}
			}
		}
	default:
		{
			if err := lib.InvokeConsumer(ib); err != nil {
				panic(err)
			}

			r := lib.ResponseBody{Type: 5}
			rj, err := json.Marshal(r)
			if err != nil {
				panic(err)
			}

			return events.APIGatewayProxyResponse{
				Body: string(rj),
				Headers: map[string]string{
					"Content-Type": "application/json",
				},
				StatusCode: 200,
			}, nil
		}
	}
}

func main() {
	lambda.Start(Handler)
}
