package main

import (
	"bytes"
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	lambdaService "github.com/aws/aws-sdk-go/service/lambda"
)

// https://mholt.github.io/json-to-go/
type InteractionBody struct {
	ApplicationID string `json:"application_id"`
	ID            string `json:"id"`
	Token         string `json:"token"`
	Type          int    `json:"type"`
	User          struct {
		Avatar           string `json:"avatar"`
		AvatarDecoration any    `json:"avatar_decoration"`
		Discriminator    string `json:"discriminator"`
		DisplayName      any    `json:"display_name"`
		GlobalName       any    `json:"global_name"`
		ID               string `json:"id"`
		PublicFlags      int    `json:"public_flags"`
		Username         string `json:"username"`
	} `json:"user"`
	Version int `json:"version"`
}

// https://github.com/bwmarrin/discordgo/blob/v0.27.1/interactions.go#L572
func VerifyInteraction(request events.APIGatewayV2HTTPRequest) bool {
	publicKey := os.Getenv("BOT_PUBLIC_KEY")
	if publicKey == "" {
		return false
	}
	key, err := hex.DecodeString(publicKey)
	if err != nil {
		return false
	}

	timestamp := request.Headers["x-signature-timestamp"]
	if timestamp == "" {
		return false
	}

	signature := request.Headers["x-signature-ed25519"]
	if signature == "" {
		return false
	}
	sig, err := hex.DecodeString(signature)
	if err != nil {
		return false
	}
	if len(sig) != 64 {
		return false
	}

	var msg bytes.Buffer
	msg.WriteString(timestamp)
	msg.WriteString(request.Body)

	return ed25519.Verify(key, msg.Bytes(), sig)
}

func Handler(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("ok")

	var interactionBody InteractionBody
	if err := json.Unmarshal([]byte(request.Body), &interactionBody); err != nil {
		panic("unmarshal")
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	consumerFn := os.Getenv("CONSUMER_FN")
	if consumerFn == "" {
		panic("fn name")
	}

	// https://github.com/aws/aws-sdk-go/issues/3385
	bodyBytes, _ := json.Marshal(interactionBody) // todo: pass request.Body directly
	payload := struct{ Body string }{Body: string(bodyBytes)}
	lambdaPayloadBytes, _ := json.Marshal(payload)

	client := lambdaService.New(sess, &aws.Config{})
	client.Invoke(&lambdaService.InvokeInput{
		FunctionName:   aws.String(consumerFn),
		InvocationType: aws.String("Event"),
		Payload:        lambdaPayloadBytes,
	})

	return events.APIGatewayProxyResponse{
		Body:       "Hello, World!",
		StatusCode: 200,
	}, nil

	switch interactionBody.Type {
	case 1:
		{
			switch VerifyInteraction(request) {
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
