package main

import (
	"encoding/json"
	"glsst/functions/lib"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
)

type response struct {
	User        lib.User         `json:"user"`
	Predictions []lib.Prediction `json:"predictions"`
}

type predictionVotersChan struct {
	Voters []lib.Voter
	Index  int
}

func Handler(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	var b struct {
		UserId string `json:"userId"`
	}
	if err := json.Unmarshal([]byte(request.Body), &b); err != nil {
		panic("unmarshal")
	}

	// fetch user
	ulc := make(chan []lib.User)
	go func() {
		q, err := lib.Query(lib.User{
			UserId: b.UserId,
		})
		if err != nil {
			panic(err)
		}

		var ul []lib.User
		attributevalue.UnmarshalListOfMaps(q.Items, &ul)

		ulc <- ul
	}()

	// fetch predictions
	plc := make(chan []lib.Prediction)
	go func() {
		q, err := lib.Query(lib.Prediction{
			UserId: b.UserId,
		}, lib.Gsi1)
		if err != nil {
			panic(err)
		}

		var pl []lib.Prediction
		attributevalue.UnmarshalListOfMaps(q.Items, &pl)

		plc <- pl
	}()

	ul := <-ulc
	close(ulc)
	var u lib.User
	if len(ul) < 1 {
		u = lib.User{}
		u.UserId = b.UserId
	} else {
		u = ul[0]
	}

	pl := <-plc
	close(plc)

	// fetch prediction voters
	// todo: better concurrency
	vlc := make(chan predictionVotersChan) // don't need buffer
	for i, p := range pl {
		go func(i int, p lib.Prediction) {
			q, err := lib.Query(lib.Voter{
				PredictionId: p.PredictionId,
			}, lib.Gsi2)
			if err != nil {
				panic(err)
			}
			var vl []lib.Voter
			attributevalue.UnmarshalListOfMaps(q.Items, &vl)
			vlc <- predictionVotersChan{
				Voters: vl,
				Index:  i,
			}
		}(i, p)
	}
	for i := range pl {
		v := <-vlc
		pl[v.Index].Voters = v.Voters
		if i == len(pl)-1 {
			close(vlc)
		}
	}

	// response
	r := response{
		User:        u,
		Predictions: pl,
	}

	rj, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}

	return events.APIGatewayProxyResponse{
		Body: string(rj),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
