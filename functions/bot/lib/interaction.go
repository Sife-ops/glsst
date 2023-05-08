package lib

// https://mholt.github.io/json-to-go/
type InteractionBody struct {
	ApplicationID string `json:"application_id"`
	ID            string `json:"id"`
	Token         string `json:"token"`
	Type          int    `json:"type"`
	User          struct {
		Avatar           string `json:"avatar"`
		AvatarDecoration any    `json:"avatar_decoration"`
		Discriminator    string `json:"discriminator"`
		DisplayName      any    `json:"display_name"`
		GlobalName       any    `json:"global_name"`
		ID               string `json:"id"`
		PublicFlags      int    `json:"public_flags"`
		Username         string `json:"username"`
	} `json:"user"`
	Version int `json:"version"`
}
