package command

import (
	"encoding/json"
	"glsst/functions/lib"
	"time"

	// "github.com/davecgh/go-spew/spew"
	"github.com/oklog/ulid/v2"
)

func Create(ib lib.InteractionBody) error {
	condition := ib.Data.Options[0]

	// todo: generate by javascript lambda lmao
	// todo: use goroutine for puts?
	pid := ulid.Make().String()
	lib.Put(lib.Prediction{
		PredictionId: pid,
		UserId:       ib.Member.User.UserId,
		Condition:    condition.Value,
		CreatedAt:    time.Now().Format(time.RFC3339),
	})

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
