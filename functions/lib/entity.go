package lib

// https://mholt.github.io/json-to-go/

type User struct {
	UserId           string `json:"id" ddb:"pk,sk,userid"`
	Username         string `json:"username" ddb:"username"`
	Discriminator    string `json:"discriminator" ddb:"discriminator"`
	DisplayName      string `json:"display_name" ddb:"displayname"`
	GlobalName       string `json:"global_name" ddb:"globalname"`
	Avatar           string `json:"avatar" ddb:"avatar"`
	AvatarDecoration any    `json:"avatar_decoration"`
	PublicFlags      int    `json:"public_flags"`
}

type InteractionBody struct {
	ApplicationID string `json:"application_id"`
	ID            string `json:"id"`
	Token         string `json:"token"`
	Type          int    `json:"type"`
	User          User   `json:"user"`
	Version       int    `json:"version"`
}
