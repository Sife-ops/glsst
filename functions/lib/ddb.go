package lib

import (
	"bytes"
	"context"
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Key struct {
	Field     string
	Composite []string
}

type Index struct {
	PartitionKey Key
	SortKey      Key
}

type Attribute struct {
	Field    string
	Type     string
	Required bool
	Default  interface{}
}

type EntitySchema struct {
	Service    string
	Entity     string
	Indexes    map[string]Index
	Attributes map[string]Attribute
}

type Entity interface {
	GetEntitySchema() EntitySchema
}

// todo: can't use pointer???
func (u User) GetEntitySchema() EntitySchema {
	return EntitySchema{
		Service: "Glsst",
		Entity:  "User",
		Indexes: map[string]Index{
			"primary": {
				PartitionKey: Key{
					Field:     "pk",
					Composite: []string{"UserId", "Username"}, // todo: composite array
				},
				SortKey: Key{
					Field:     "sk",
					Composite: []string{},
				},
			},
			"gsi1": {
				PartitionKey: Key{
					Field:     "gsi1pk",
					Composite: []string{"UserId"},
				},
				SortKey: Key{
					Field:     "gsi1sk",
					Composite: []string{},
				},
			},
		},
		Attributes: map[string]Attribute{
			"UserId": {
				Field: "userid",
				Type:  "string",
			},
			"Username": {
				Field: "username",
				Type:  "string",
			},
			"Discriminator": {
				Field: "discriminator",
				Type:  "string",
			},
			"DisplayName": {
				Field: "displayname",
				Type:  "string",
			},
		},
	}
}

func EntityToMap(e Entity) map[string]types.AttributeValue {
	// 1) map of struct key/val
	// 2) build indexes
	// 3) build attributes

	t := reflect.TypeOf(e)
	v := reflect.ValueOf(e)

	// 1) map of struct key/val
	m := map[string]reflect.Value{}
	for i := 0; i < v.NumField(); i++ {
		tf := t.Field(i)
		vf := v.Field(i)
		m[tf.Name] = vf
	}

	mm := map[string]types.AttributeValue{}
	mm["__entity"] = &types.AttributeValueMemberS{Value: e.GetEntitySchema().Entity}

	// 2) build indexes
	for _, index := range e.GetEntitySchema().Indexes {
		var pkb bytes.Buffer
		pkb.WriteString("$" + e.GetEntitySchema().Service + "#" + e.GetEntitySchema().Entity)

		for _, c := range index.PartitionKey.Composite {
			pkb.WriteString("#" + strings.ToLower(c) + "_") // todo: why tolower?
			pkb.WriteString(reflect.ValueOf(m[c].Interface()).String())
		}
		mm[index.PartitionKey.Field] = &types.AttributeValueMemberS{Value: pkb.String()}

		var skb bytes.Buffer
		skb.WriteString("$" + e.GetEntitySchema().Entity)

		skc := index.SortKey.Composite
		if len(skc) > 0 {
			for _, c := range index.SortKey.Composite {
				skb.WriteString("#" + strings.ToLower(c) + "_")
				skb.WriteString(reflect.ValueOf(m[c].Interface()).String())
			}
		}
		mm[index.SortKey.Field] = &types.AttributeValueMemberS{Value: skb.String()}
	}

	// 3) build attributes
	// todo: more types
	// todo: required attributes
	for attrk, attrv := range e.GetEntitySchema().Attributes {
		v := reflect.ValueOf(m[attrk].Interface())
		af := attrv.Field
		switch attrv.Type {
		case "string":
			mm[af] = &types.AttributeValueMemberS{Value: v.String()}
		case "boolean":
			mm[af] = &types.AttributeValueMemberBOOL{Value: v.Bool()}
		default:
			panic("invalid type: " + attrv.Type)
		}
	}

	return mm
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
