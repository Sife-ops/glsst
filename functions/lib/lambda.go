package lib

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	lambdaSvc "github.com/aws/aws-sdk-go-v2/service/lambda"
)

func Mnemonic() (string, error) {
	o, err := LambdaCl.Invoke(context.TODO(), &lambdaSvc.InvokeInput{
		FunctionName: aws.String(GetMnemonicFn()),
	})
	if err != nil {
		return "", err
	}
	return strings.Trim(string(o.Payload), "\""), nil
}

// https://github.com/aws/aws-sdk-go/issues/3385
func InvokeConsumer(b InteractionBody) error {
	bodyBytes, _ := json.Marshal(b)
	payload := struct{ Body string }{Body: string(bodyBytes)}
	payloadBytes, _ := json.Marshal(payload)

	if _, err := LambdaCl.Invoke(context.TODO(), &lambdaSvc.InvokeInput{
		FunctionName:   aws.String(GetConsumerFn()),
		InvocationType: "Event",
		Payload:        payloadBytes,
	}); err != nil {
		return err
	}

	return nil
}
