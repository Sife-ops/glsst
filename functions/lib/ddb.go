package lib

import (
	"bytes"
	"context"
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type IndexName string

const Primary IndexName = "primary"

type Key struct {
	Field     string
	Composite []string
}

type Index struct {
	PartitionKey Key
	SortKey      Key
}

// todo: implement default values?
type Attribute struct {
	Field    string
	Type     string
	Required bool
	Default  interface{}
}

type EntitySchema struct {
	Service    string
	Entity     string
	Indexes    map[IndexName]Index
	Attributes map[string]Attribute
}

type Entity interface {
	GetEntitySchema() EntitySchema
}

func EntityToMap(e Entity) map[string]types.AttributeValue {
	m, err := attributevalue.MarshalMap(e)
	if err != nil {
		panic(err)
	}

	// indexes
	r := reflect.ValueOf(e)
	for _, index := range e.GetEntitySchema().Indexes {
		// partition key
		var pkb bytes.Buffer
		pkb.WriteString("$" + e.GetEntitySchema().Service + "#" + e.GetEntitySchema().Entity)
		for _, c := range index.PartitionKey.Composite {
			pkb.WriteString("#" + strings.ToLower(c) + "_") // todo: why tolower?
			f := reflect.Indirect(r).FieldByName(c)
			pkb.WriteString(f.String())
		}
		m[index.PartitionKey.Field] = &types.AttributeValueMemberS{Value: pkb.String()}

		// sort key
		var skb bytes.Buffer
		skb.WriteString("$" + e.GetEntitySchema().Entity)
		skc := index.SortKey.Composite
		if len(skc) > 0 {
			for _, c := range index.SortKey.Composite {
				skb.WriteString("#" + strings.ToLower(c) + "_")
				f := reflect.Indirect(r).FieldByName(c)
				skb.WriteString(f.String())
			}
		}
		m[index.SortKey.Field] = &types.AttributeValueMemberS{Value: skb.String()}
	}

	return m
}

func BaseQueryInput() *dynamodb.QueryInput {
	return &dynamodb.QueryInput{
		TableName: aws.String(GetTableName()),
	}
}

func Query(e Entity, i ...IndexName) *dynamodb.QueryOutput {
	ddbq := BaseQueryInput()

	var ii IndexName
	switch len(i) > 0 {
	case true:
		ii = i[0]
	default:
		ii = Primary
	}
	if ii != Primary {
		ddbq.IndexName = aws.String(string(ii))
	}
	pk := e.GetEntitySchema().Indexes[ii].PartitionKey.Field
	sk := e.GetEntitySchema().Indexes[ii].SortKey.Field
	ddbq.KeyConditionExpression = aws.String(pk + " = :" + pk + " and " + sk + " = :" + sk)

	m := EntityToMap(e)
	ddbq.ExpressionAttributeValues = map[string]types.AttributeValue{
		":" + pk: m[pk],
		":" + sk: m[sk],
	}
	// 	// ExpressionAttributeNames: map[string]string{
	// 	// },

	out, err := DdbCl.Query(context.TODO(), ddbq)
	if err != nil {
		panic(err)
	}

	return out
}

func Put(e Entity) *dynamodb.PutItemOutput {
	out, err := DdbCl.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(GetTableName()),
		Item:      EntityToMap(e),
	})

	if err != nil {
		panic(err)
	}

	return out
}
