package main

import (
	"context"
	"encoding/json"
	"glsst/functions/lib"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	lambdaSvc "github.com/aws/aws-sdk-go-v2/service/lambda"
)

func Handler(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	var interactionBody lib.InteractionBody
	if err := json.Unmarshal([]byte(request.Body), &interactionBody); err != nil {
		panic("unmarshal")
	}

	switch interactionBody.Type {
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
			// https://github.com/aws/aws-sdk-go/issues/3385
			bodyBytes, _ := json.Marshal(interactionBody) // todo: pass request.Body directly
			payload := struct{ Body string }{Body: string(bodyBytes)}
			payloadBytes, _ := json.Marshal(payload)

			lib.LambdaCl.Invoke(context.TODO(), &lambdaSvc.InvokeInput{
				FunctionName:   aws.String(lib.GetConsumerFn()),
				InvocationType: "Event",
				Payload:        payloadBytes,
			})

			r := lib.ResponseBody{
				Type: 5,
				Data: lib.ResponseData{
					Flags: 64,
				},
			}
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
