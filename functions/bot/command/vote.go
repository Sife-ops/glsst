package command

import (
	"fmt"
	"glsst/functions/lib"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/oklog/ulid/v2"
)

func Vote(ib lib.InteractionBody) error {
	pid := ib.Data.Options[0].Value.(string)
	vote := ib.Data.Options[1].Value.(bool)

	// 1) prediction must exist
	q, err := lib.Query(lib.Prediction{
		PredictionId: pid,
	})
	if err != nil {
		return err
	}

	if len(q.Items) < 1 {
		return ib.FollowUp(lib.ResponseData{
			Content: "prediction does not exist",
		})
	}

	// 2) no self vote
	var p []lib.Prediction
	attributevalue.UnmarshalListOfMaps(q.Items, &p)
	if p[0].UserId == ib.Member.User.UserId {
		return ib.FollowUp(lib.ResponseData{
			Content: fmt.Sprintf("<@%s> tried to vote on their own prediction.", ib.Member.User.UserId),
		})
	}

	// 3) insert vote
	qq, err := lib.Query(lib.Voter{
		UserId:       ib.Member.User.UserId,
		PredictionId: pid,
	}, lib.Gsi1)
	if err != nil {
		return err
	}

	if len(qq.Items) < 1 {
		if _, err := lib.Put(lib.Voter{
			VoterId:      ulid.Make().String(),
			PredictionId: pid,
			UserId:       ib.Member.User.UserId,
			Vote:         vote, // todo: tristate
		}); err != nil {
			return err
		}

		return ib.FollowUp(lib.ResponseData{
			Content: "voted",
		})
	}

	// // spew.Dump(qq)
	// // todo: DEBUG
	// fmt.Println(len(qq.Items))
	// return ib.FollowUp(lib.ResponseData{
	// 	Content: "vote",
	// })

	return ib.FollowUp(lib.ResponseData{
		Content: "already voted",
	})
}
