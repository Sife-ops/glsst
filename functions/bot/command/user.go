package command

import (
	"fmt"
	"glsst/functions/lib"
)

func User(ib lib.InteractionBody) error {
	infoUser := ib.Data.Options[0].Value.(string)
	fmt.Println(infoUser)

	return ib.FollowUp(lib.ResponseData{
		// todo: web url
		Content: "User",
	})
}
