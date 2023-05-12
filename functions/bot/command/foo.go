package command

import (
	"bytes"
	"encoding/json"
	"fmt"
	"glsst/functions/lib"
	"net/http"
)

func FollowUp(ib lib.InteractionBody, rb []byte) error {
	url := fmt.Sprintf(
		"https://discord.com/api/v10/webhooks/%s/%s",
		lib.GetBotAppId(),
		ib.Token,
	)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(rb))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", lib.GetBotPk())

	cl := &http.Client{}   // todo: declare global client?
	res, err := cl.Do(req) // todo: can use _ ??????
	if err != nil {
		return err
	}
	if res != nil {
	}

	return nil
}

func Foo(ib lib.InteractionBody) error {
	r := lib.ResponseData{
		Content: "ree",
	}
	rj, err := json.Marshal(r)
	if err != nil {
		return err
	}

	return FollowUp(ib, rj)
}
