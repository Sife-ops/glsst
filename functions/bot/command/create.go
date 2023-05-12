package command

import (
	"encoding/json"
	"fmt"
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

	r := lib.ResponseData{
		Embeds: []lib.Embed{
			{
				Title:       "New Prediction",
				Description: condition.Value,
				Fields: []lib.Field{
					{
						Name:   "By",
						Value:  fmt.Sprintf("<@%s>", ib.Member.User.UserId),
						Inline: false,
					},
					{
						Name:   "ID",
						Value:  m,
						Inline: false,
					},
				},
			},
		},
	}
	rj, err := json.Marshal(r)
	if err != nil {
		return err
	}

	return FollowUp(ib, rj)
}
