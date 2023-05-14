package command

import (
	"fmt"
	"glsst/functions/lib"
)

func User(b lib.InteractionBody) error {
	return b.FollowUp(lib.ResponseData{
		// todo: user details
		Embeds: []lib.Embed{
			{
				Title: "URL",
				Url:   fmt.Sprintf("%s/user/%s", lib.GetSiteUrl(), b.Member.User.UserId),
			},
		},
	})
}
