package lib

import (
	"context"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

var LambdaCl *lambda.Client
var DdbCl *dynamodb.Client
var HttpCl *http.Client

func init() {
	// todo: hardcoded region
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	LambdaCl = lambda.NewFromConfig(cfg)
	DdbCl = dynamodb.NewFromConfig(cfg)
	HttpCl = &http.Client{}
}
