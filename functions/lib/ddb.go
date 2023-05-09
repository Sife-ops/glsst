package lib

import (
	"context"
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// https://goplay.space/#15XAsJBH5Sb
func ToAVMap(inf interface{}) map[string]types.AttributeValue {
	t := reflect.TypeOf(inf)
	v := reflect.ValueOf(inf)

	m := map[string]types.AttributeValue{}

	for i := 0; i < t.NumField(); i++ {
		vt := t.Field(i)

		tag, hasTag := vt.Tag.Lookup("ddb")
		if hasTag == false {
			continue
		}

		vf := v.Field(i)
		vfi := vf.Interface()
		vfv := reflect.ValueOf(vfi)

		ks := strings.Split(tag, ",")
		for _, k := range ks {
			m[k] = &types.AttributeValueMemberS{Value: vfv.String()} // todo: switch on type
		}
	}

	return m
}

func PutAVMap(m map[string]types.AttributeValue) *dynamodb.PutItemOutput {
	out, err := DdbCl.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(GetTableName()),
		Item:      m,
	})

	if err != nil {
		panic(err)
	}

	return out
}

type Entity interface {
	Put() *dynamodb.PutItemOutput
}

func (u User) Put() *dynamodb.PutItemOutput {
	m := ToAVMap(u)
	return PutAVMap(m)
}
