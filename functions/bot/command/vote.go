package command

import (
	"fmt"
	"glsst/functions/lib"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/oklog/ulid/v2"
)

func Vote(b lib.InteractionBody) error {
	memberId := b.Member.User.UserId
	pid := b.Data.Options[0].Value.(string)
	vote := b.Data.Options[1].Value.(bool)

	// 1) prediction must exist
	q, err := lib.Query(lib.Prediction{
		PredictionId: pid,
	})
	if err != nil {
		return err
	}

	if len(q.Items) < 1 {
		return b.FollowUp(lib.ResponseData{
			Content: fmt.Sprintf("<@%s> voted on a nonexistent prediction.", memberId),
		})
	}

	// 2) no self vote
	var pl []lib.Prediction
	attributevalue.UnmarshalListOfMaps(q.Items, &pl)
	p := pl[0]
	if p.UserId == memberId {
		return b.FollowUp(lib.ResponseData{
			Content: fmt.Sprintf("<@%s> tried to vote on their own prediction.", memberId),
		})
	}

	// 3) insert vote
	qq, err := lib.Query(lib.Voter{
		UserId:       memberId,
		PredictionId: pid,
	}, lib.Gsi1)
	if err != nil {
		return err
	}

	if len(qq.Items) > 0 {
		return b.FollowUp(lib.ResponseData{
			Content: fmt.Sprintf("<@%s> tried to vote twice.", memberId),
		})
	}

	if _, err := lib.Put(lib.Voter{
		VoterId:      ulid.Make().String(),
		PredictionId: pid,
		UserId:       memberId,
		Vote:         vote, // todo: tristate
	}); err != nil {
		return err
	}

	var fa string
	var color int
	if vote {
		fa = "in favor of"
		color = 65280
	} else {
		fa = "against"
		color = 16711680
	}

	return b.FollowUp(lib.ResponseData{
		Content: "voted",
		Embeds: []lib.Embed{
			{
				Title:       "Vote",
				Color:       color,
				Description: fmt.Sprintf("<@%s> voted %s <@%s>'s prediction:", memberId, fa, p.UserId),
				Fields: []lib.Field{
					{
						Name:   "Condition(s)",
						Value:  p.Condition,
						Inline: false,
					},
					{
						Name:   "ID",
						Value:  p.PredictionId,
						Inline: false,
					},
				},
			},
		},
	})
}
