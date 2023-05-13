package main

import (
	"fmt"
	"glsst/functions/lib"
	"sync"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/oklog/ulid/v2"
)

type Event struct {
	Predictions []lib.Prediction `json:"predictions"`
	Voters      []lib.Voter      `json:"voters"`
}

// todo: // import users
func Handler(event Event) error {
	fmt.Println("import")

	var wg sync.WaitGroup

	for _, p := range event.Predictions {
		wg.Add(1)

		go func(p lib.Prediction) {
			wg.Done()
			p.CreatedAt = time.UnixMilli(int64(p.ImportCreatedAt)).Format(time.RFC3339)
			fmt.Println("insert prediction", p.PredictionId)
			if _, err := lib.Put(p); err != nil {
				panic(err)
			}
		}(p)
	}

	for _, v := range event.Voters {
		wg.Add(1)

		go func(v lib.Voter) {
			wg.Done()

			switch v.Verdict {
			case "correct":
				v.Vote = true
			case "incorrect":
				v.Vote = false
			default:
				return
			}

			fmt.Println("insert voter for", v.PredictionId)
			v.VoterId = ulid.Make().String()
			if _, err := lib.Put(v); err != nil {
				// return err
				panic(err)
			}
		}(v)
	}

	wg.Wait()
	fmt.Println("done")

	return nil
}

func main() {
	lambda.Start(Handler)
}
