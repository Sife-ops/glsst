package main

import (
	"encoding/json"
	"glsst/functions/lib"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
)

type Response struct {
	User        lib.User         `json:"user"`
	Predictions []lib.Prediction `json:"predictions"`
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
	go func(ulc chan []lib.User) {
		q, err := lib.Query(lib.User{
			UserId: b.UserId,
		})
		if err != nil {
			panic(err)
		}

		var ul []lib.User
		attributevalue.UnmarshalListOfMaps(q.Items, &ul)

		ulc <- ul
	}(ulc)

	// fetch predictions
	plc := make(chan []lib.Prediction)
	go func(plc chan []lib.Prediction) {
		q, err := lib.Query(lib.Prediction{
			UserId: b.UserId,
		}, lib.Gsi1)
		if err != nil {
			panic(err)
		}

		var pl []lib.Prediction
		attributevalue.UnmarshalListOfMaps(q.Items, &pl)

		plc <- pl
	}(plc)

	ul := <-ulc
	var u lib.User
	if len(ul) < 1 {
		u = lib.User{}
		u.UserId = b.UserId
	} else {
		u = ul[0]
	}

	pl := <-plc

	// fetch prediction voters
	vlc := make(chan []lib.Voter, len(pl))
	for _, p := range pl {
		go func(aa chan []lib.Voter, p lib.Prediction) {
			var vl []lib.Voter

			q, err := lib.Query(lib.Voter{
				PredictionId: p.PredictionId,
			}, lib.Gsi2)
			if err != nil {
				panic(err)
			}
			attributevalue.UnmarshalListOfMaps(q.Items, &vl)

			vlc <- vl
		}(vlc, p)
	}
	for i := range pl {
		pl[i].Voters = <-vlc
	}
	close(vlc)

	// response
	r := Response{
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
