package lib

// https://mholt.github.io/json-to-go/

type InteractionBody struct {
	ApplicationID string          `json:"application_id"`
	ID            string          `json:"id"`
	Token         string          `json:"token"`
	Type          int             `json:"type"`
	Version       int             `json:"version"`
	Member        Member          `json:"member"`
	Data          InteractionData `json:"data"`
}

type Member struct {
	User User `json:"user"`
}

type Option struct {
	Name  string      `json:"name"`
	Type  int         `json:"type"`
	Value interface{} `json:"value"`
}

type InteractionData struct {
	GuildId string   `json:"guild_id"`
	Id      string   `json:"id"`
	Name    string   `json:"name"`
	Type    int      `json:"type"`
	Options []Option `json:"options"`
}

////////////////////////////////////////////////////////////////////////////////

type User struct {
	UserId           string `json:"id" dynamodbav:"userid"`
	Username         string `json:"username" dynamodbav:"username"`
	Discriminator    string `json:"discriminator" dynamodbav:"discriminator"`
	DisplayName      string `json:"display_name" dynamodbav:"displayname"`
	GlobalName       string `json:"global_name" dynamodbav:"globalname"`
	Avatar           string `json:"avatar" dynamodbav:"avatar"`
	AvatarDecoration any    `json:"avatar_decoration" dynamodbav:"-"`
	PublicFlags      int    `json:"public_flags" dynamodbav:"-"`
}

type Prediction struct {
	PredictionId    string  `json:"predictionId" dynamodbav:"predictionid"`
	UserId          string  `json:"userId" dynamodbav:"userid"`
	Condition       string  `json:"condition" dynamodbav:"condition"`
	CreatedAt       string  `json:"createdAt" dynamodbav:"createdat"`
	Voters          []Voter `json:"voters" dynamodbav:"-"`
	ImportCreatedAt int     `json:"created_at" dynamodbav:"-"` // todo: DELETE
}

type Voter struct {
	VoterId      string `json:"voterId" dynamodbav:"voterid"`
	PredictionId string `json:"predictionId" dynamodbav:"predictionid"`
	UserId       string `json:"userId" dynamodbav:"userid"`
	Vote         bool   `json:"vote" dynamodbav:"vote"`
	Verdict      string `json:"verdict" dynamodbav:"-"` // todo: DELETE
}

////////////////////////////////////////////////////////////////////////////////

type ResponseBody struct {
	Type int          `json:"type"`
	Data ResponseData `json:"data"`
}

type ResponseData struct {
	Flags   int     `json:"flags"`
	Content string  `json:"content"`
	Embeds  []Embed `json:"embeds"`
}

type Embed struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Color       int     `json:"color"`
	Url         string  `json:"url"`
	Fields      []Field `json:"fields"`
}

type Field struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}
