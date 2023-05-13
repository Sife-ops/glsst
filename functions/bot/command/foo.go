package command

import (
	"glsst/functions/lib"
)

func Foo(ib lib.InteractionBody) error {
	return ib.FollowUp(lib.ResponseData{
		Content: "bar",
	})
}
