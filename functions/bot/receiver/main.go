package main

import (
	"encoding/json"
	"glsst/functions/lib"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	var b lib.InteractionBody
	if err := json.Unmarshal([]byte(request.Body), &b); err != nil {
		panic("unmarshal") // todo: send "oops" to discord
	}

	switch b.Type {
	case 1:
		{
			verified, err := lib.VerifyInteraction(lib.VerifyInteractionInput{
				PublicKey: lib.GetBotPk(),
				Timestamp: request.Headers["x-signature-timestamp"],
				Signature: request.Headers["x-signature-ed25519"],
				Body:      request.Body,
			})
			if err != nil {
				panic(err)
			}
			switch verified {
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
			if err := lib.InvokeConsumer(b); err != nil {
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
