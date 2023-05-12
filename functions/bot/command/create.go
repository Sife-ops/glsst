package command

import (
	"encoding/json"
	"glsst/functions/lib"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/oklog/ulid/v2"
)

func Create(ib lib.InteractionBody) error {
	condition := ib.Data.Options[0]
	judge := ib.Data.Options[1]

	// todo: generate by javascript lambda lmao
	// todo: use goroutine for puts?
	pid := ulid.Make().String()
	lib.Put(lib.Prediction{
		PredictionId: pid,
		UserId:       ib.Member.User.UserId,
		Condition:    condition.Value,
		CreatedAt:    time.Now().Format(time.RFC3339),
		Verdict:      lib.Undecided,
	})

	o := lib.Put(lib.Judge{
		JudgeId:      ulid.Make().String(),
		UserId:       judge.Value,
		PredictionId: pid,
		Verdict:      lib.Undecided,
	})

	spew.Dump(o)

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
