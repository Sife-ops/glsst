package lib

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

var Sess *session.Session
var LambdaCl *lambda.Lambda

func init() {
	Sess = session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	})) // todo: ???

	LambdaCl = lambda.New(Sess, &aws.Config{})
}
