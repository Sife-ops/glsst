package lib

// https://mholt.github.io/json-to-go/

type User struct {
	Pk     string `dynamodbav:"pk"`
	Sk     string `dynamodbav:"sk"`
	Gsi1pk string `dynamodbav:"gsi1pk"`
	Gsi1sk string `dynamodbav:"gsi1sk"`

	UserId           string `json:"id" dynamodbav:"userid"`
	Username         string `json:"username" dynamodbav:"username"`
	Discriminator    string `json:"discriminator" dynamodbav:"discriminator"`
	DisplayName      string `json:"display_name" dynamodbav:"displayname"`
	GlobalName       string `json:"global_name" dynamodbav:"globalname"`
	Avatar           string `json:"avatar" dynamodbav:"avatar"`
	AvatarDecoration any    `json:"avatar_decoration" dynamodbav:"omitempty"`
	PublicFlags      int    `json:"public_flags" dynamodbav:"omitempty"`
}

type InteractionBody struct {
	ApplicationID string `json:"application_id"`
	ID            string `json:"id"`
	Token         string `json:"token"`
	Type          int    `json:"type"`
	User          User   `json:"user"`
	Version       int    `json:"version"`
}

type Prediction struct {
	PredictionId string `json:"predictionid" dynamodbav:"predictionid"`
	UserId       string `json:"userid" dynamodbav:"userid"`
	Condition    string `json:"condition" dynamodbav:"condition"`
	Verdict      string `json:"verdict" dynamodbav:"verdict"`
	CreatedAt    string `json:"createdat" dynamodbav:"createdat"`
}
