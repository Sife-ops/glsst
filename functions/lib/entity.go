package lib

// https://mholt.github.io/json-to-go/

type User struct {
	// Pk     string `dynamodbav:"pk"`
	// Sk     string `dynamodbav:"sk"`
	// Gsi1pk string `dynamodbav:"gsi1pk"`
	// Gsi1sk string `dynamodbav:"gsi1sk"`

	UserId           string `json:"id" dynamodbav:"userid"`
	Username         string `json:"username" dynamodbav:"username"`
	Discriminator    string `json:"discriminator" dynamodbav:"discriminator"`
	DisplayName      string `json:"display_name" dynamodbav:"displayname"`
	GlobalName       string `json:"global_name" dynamodbav:"globalname"`
	Avatar           string `json:"avatar" dynamodbav:"avatar"`
	AvatarDecoration any    `json:"avatar_decoration" dynamodbav:"omitempty"`
	PublicFlags      int    `json:"public_flags" dynamodbav:"omitempty"`
}

type Member struct {
	User User `json:"user"`
}

type Option struct {
	Name  string `json:"name"`
	Type  int    `json:"type"`
	Value string `json:"value"`
}

type InteractionData struct {
	GuildId string   `json:"guild_id"`
	Id      string   `json:"id"`
	Name    string   `json:"name"`
	Type    int      `json:"type"`
	Options []Option `json:"options"`
}

type InteractionBody struct {
	ApplicationID string          `json:"application_id"`
	ID            string          `json:"id"`
	Token         string          `json:"token"`
	Type          int             `json:"type"`
	Version       int             `json:"version"`
	Member        Member          `json:"member"`
	Data          InteractionData `json:"data"`
}

type Prediction struct {
	// Pk     string `dynamodbav:"pk"` // todo: delete?
	// Sk     string `dynamodbav:"sk"`
	// Gsi1pk string `dynamodbav:"gsi1pk"`
	// Gsi1sk string `dynamodbav:"gsi1sk"`

	PredictionId string `json:"predictionid" dynamodbav:"predictionid"`
	UserId       string `json:"userid" dynamodbav:"userid"`
	Condition    string `json:"condition" dynamodbav:"condition"`
	Verdict      string `json:"verdict" dynamodbav:"verdict"`
	CreatedAt    string `json:"createdat" dynamodbav:"createdat"`
}

type ResponseData struct {
	Flags   int    `json:"flags"`
	Content string `json:"content"`
}

type ResponseBody struct {
	Type int          `json:"type"`
	Data ResponseData `json:"data"` // todo: embed data
}
