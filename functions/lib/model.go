package lib

const (
	Gsi1 IndexName = "gsi1" // User
	Gsi2 IndexName = "gsi2" // Prediction
)

// todo: can't use pointer???
func (u User) GetEntitySchema() EntitySchema {
	return EntitySchema{
		Service: "Glsst",
		Entity:  "User",
		Indexes: map[IndexName]Index{
			Primary: {
				PartitionKey: Key{
					Field:     "pk",
					Composite: []string{"UserId"},
				},
				SortKey: Key{
					Field:     "sk",
					Composite: []string{},
				},
			},
			Gsi1: {
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
		Indexes: map[IndexName]Index{
			Primary: {
				PartitionKey: Key{
					Field:     "pk",
					Composite: []string{"PredictionId"},
				},
				SortKey: Key{
					Field:     "sk",
					Composite: []string{},
				},
			},
			Gsi1: {
				PartitionKey: Key{
					Field:     "gsi1pk",
					Composite: []string{"UserId"},
				},
				SortKey: Key{
					Field:     "gsi1sk",
					Composite: []string{"PredictionId"},
				},
			},
			Gsi2: {
				PartitionKey: Key{
					Field:     "gsi2pk",
					Composite: []string{"PredictionId"},
				},
				SortKey: Key{
					Field:     "gsi2sk",
					Composite: []string{},
				},
			},
		},
	}
}

func (v Voter) GetEntitySchema() EntitySchema {
	return EntitySchema{
		Service: "Glsst",
		Entity:  "Voter",
		Indexes: map[IndexName]Index{
			Primary: {
				PartitionKey: Key{
					Field:     "pk",
					Composite: []string{"VoterId"},
				},
				SortKey: Key{
					Field:     "sk",
					Composite: []string{},
				},
			},
			Gsi1: {
				PartitionKey: Key{
					Field:     "gsi1pk",
					Composite: []string{"UserId"},
				},
				SortKey: Key{
					Field:     "gsi1sk",
					Composite: []string{"PredictionId"},
				},
			},
			Gsi2: {
				PartitionKey: Key{
					Field:     "gsi2pk",
					Composite: []string{"PredictionId"},
				},
				SortKey: Key{
					Field:     "gsi2sk",
					Composite: []string{"UserId"},
				},
			},
		},
	}
}
