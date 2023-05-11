package lib

const Gsi1 IndexName = "gsi1"

// todo: can't use pointer???
func (u User) GetEntitySchema() EntitySchema {
	return EntitySchema{
		Service: "Glsst",
		Entity:  "User",
		Indexes: map[IndexName]Index{
			Primary: {
				PartitionKey: Key{
					Field: Field{
						Struct:   "Pk",
						DynamoDB: "pk",
					},
					Composite: []string{"UserId"},
				},
				SortKey: Key{
					Field: Field{
						Struct:   "Sk",
						DynamoDB: "sk",
					},
					Composite: []string{},
				},
			},
			Gsi1: {
				PartitionKey: Key{
					Field: Field{
						Struct:   "Gsi1pk",
						DynamoDB: "gsi1pk",
					},
					Composite: []string{"UserId"},
				},
				SortKey: Key{
					Field: Field{
						Struct:   "Gsi1sk",
						DynamoDB: "gsi1sk",
					},
					Composite: []string{},
				},
			},
		},
	}
}

func (p Prediction) GetEntitySchema() EntitySchema {
	return EntitySchema{
		Service: "Glsst",
		Entity:  "Prediction",
		Indexes: map[IndexName]Index{
			Primary: {
				PartitionKey: Key{
					Field: Field{
						Struct:   "Pk",
						DynamoDB: "pk",
					},
					Composite: []string{"PredictionId"},
				},
				SortKey: Key{
					Field: Field{
						Struct:   "Sk",
						DynamoDB: "sk",
					},
					Composite: []string{},
				},
			},
			Gsi1: {
				PartitionKey: Key{
					Field: Field{
						Struct:   "Gsi1pk",
						DynamoDB: "gsi1pk",
					},
					Composite: []string{"UserId"},
				},
				SortKey: Key{
					Field: Field{
						Struct:   "Gsi1sk",
						DynamoDB: "gsi1sk",
					},
					Composite: []string{},
				},
			},
		},
	}
}
