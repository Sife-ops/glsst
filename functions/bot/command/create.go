package command

import (
	"context"
	"encoding/json"
	"glsst/functions/lib"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	lambdaSvc "github.com/aws/aws-sdk-go-v2/service/lambda"
)

func Create(ib lib.InteractionBody) error {
	condition := ib.Data.Options[0]

	o, err := lib.LambdaCl.Invoke(context.TODO(), &lambdaSvc.InvokeInput{
		FunctionName: aws.String(lib.GetMnemonicFn()),
	})
	if err != nil {
		return err
	}
	m := strings.Trim(string(o.Payload), "\"")

	// todo: use goroutine for puts?
	if _, err := lib.Put(lib.Prediction{
		PredictionId: m,
		UserId:       ib.Member.User.UserId,
		Condition:    condition.Value,
		CreatedAt:    time.Now().Format(time.RFC3339),
	}); err != nil {
		return err
	}

	// todo: embeds, info...
	r := lib.ResponseData{
		Content: "create lmao",
	}
	rj, err := json.Marshal(r)
	if err != nil {
		return err
	}

	return FollowUp(ib, rj)
}
