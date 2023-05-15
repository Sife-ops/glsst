package command

import (
	"fmt"
	"glsst/functions/lib"
)

func User(b lib.InteractionBody) error {
	u := b.Data.Options[0].Value.(string)
	return b.FollowUp(lib.ResponseData{
		// todo: user details
		Embeds: []lib.Embed{
			{
				Title: "URL",
				Url:   fmt.Sprintf("%s/user/%s", lib.GetSiteUrl(), u),
			},
		},
	})
}
