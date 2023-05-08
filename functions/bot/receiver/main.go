package main

import (
	"encoding/json"
	botLib "glsst/functions/bot/lib"
	"glsst/functions/lib"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	lambdaService "github.com/aws/aws-sdk-go/service/lambda"
)

func Handler(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	// todo: global
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	var interactionBody botLib.InteractionBody // todo: constructor method?
	if err := json.Unmarshal([]byte(request.Body), &interactionBody); err != nil {
		panic("unmarshal")
	}

	switch interactionBody.Type {
	case 1:
		{
			switch botLib.VerifyInteraction(request) {
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
			lambdaPayloadBytes, _ := json.Marshal(payload)

			client := lambdaService.New(sess, &aws.Config{})
			client.Invoke(&lambdaService.InvokeInput{
				FunctionName:   aws.String(lib.GetBotConsumerFn()),
				InvocationType: aws.String("Event"),
				Payload:        lambdaPayloadBytes,
			})

			return events.APIGatewayProxyResponse{
				Body:       "Hello, World! Your yddidtiyti request was received at OK? " + request.RequestContext.Time + ".",
				StatusCode: 200,
			}, nil
		}
	}
}

func main() {
	lambda.Start(Handler)
}
