package command

import (
	"glsst/functions/lib"
)

func Foo(b lib.InteractionBody) error {
	return b.FollowUp(lib.ResponseData{
		Content: "bar",
	})
}
