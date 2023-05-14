package lib

import (
	"bytes"
	"context"
	"fmt"
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

	m["_entity"] = &types.AttributeValueMemberS{Value: e.GetEntitySchema().Entity}

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

func BaseUpdateItemInput() *dynamodb.UpdateItemInput {
	return &dynamodb.UpdateItemInput{
		TableName: aws.String(GetTableName()),
	}
}

func Query(e Entity, i ...IndexName) (*dynamodb.QueryOutput, error) {
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
	ddbq.KeyConditionExpression = aws.String(pk + " = :" + pk + " and begins_with(" + sk + ", :" + sk + ")")

	m := EntityToMap(e)
	ddbq.ExpressionAttributeValues = map[string]types.AttributeValue{
		":" + pk: m[pk],
		":" + sk: m[sk],
	}
	// todo: expression attribute names
	// ExpressionAttributeNames: map[string]string{
	// },

	return DdbCl.Query(context.TODO(), ddbq)
}

// todo: add, subtrace
func Update(e Entity, u map[string]interface{}) (*dynamodb.UpdateItemOutput, error) {
	// key
	em := EntityToMap(e)
	pk := e.GetEntitySchema().Indexes[Primary].PartitionKey.Field
	sk := e.GetEntitySchema().Indexes[Primary].SortKey.Field
	ddbq := BaseUpdateItemInput()
	ddbq.Key = map[string]types.AttributeValue{
		pk: em[pk],
		sk: em[sk],
	}

	um, err := attributevalue.MarshalMap(u)
	if err != nil {
		return nil, err
	}

	// update expression
	var buf bytes.Buffer
	buf.WriteString("SET")
	var keys []string
	for k := range um {
		keys = append(keys, k)
	}
	for i, v := range keys {
		switch {
		case i < 1:
			buf.WriteString(fmt.Sprintf(" #%s = :%s", v, v))
		default:
			buf.WriteString(fmt.Sprintf(", #%s = :%s", v, v))
		}
	}
	ddbq.UpdateExpression = aws.String(buf.String())

	// expression attribute names
	ean := map[string]string{}
	for k := range um {
		ean["#"+k] = k
	}
	ddbq.ExpressionAttributeNames = ean

	// expression attribute values
	eav := map[string]types.AttributeValue{}
	for k, v := range um {
		eav[":"+k] = v
	}
	ddbq.ExpressionAttributeValues = eav

	// update
	return DdbCl.UpdateItem(context.TODO(), ddbq)
}

func Put(e Entity) (*dynamodb.PutItemOutput, error) {
	return DdbCl.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(GetTableName()),
		Item:      EntityToMap(e),
	})
}
