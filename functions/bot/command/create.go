package command

import (
	"encoding/json"
	"glsst/functions/lib"
	"time"
)

func Create(ib lib.InteractionBody) error {
	condition := ib.Data.Options[0]

	m, err := lib.Mnemonic()
	if err != nil {
		return err
	}

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
