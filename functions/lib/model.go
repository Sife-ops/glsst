package lib

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
	}
}

func (p Prediction) GetEntitySchema() EntitySchema {
	return EntitySchema{
		Service: "Glsst",
		Entity:  "Prediction",
		Indexes: map[string]Index{
			"primary": {
				PartitionKey: Key{
					Field:     "pk",
					Composite: []string{"PredictionId"},
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
	}
}
