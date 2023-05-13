package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (ib InteractionBody) FollowUp(rd ResponseData) error {
	rb, err := json.Marshal(rd)
	if err != nil {
		return err
	}

	url := fmt.Sprintf(
		"https://discord.com/api/v10/webhooks/%s/%s",
		GetBotAppId(),
		ib.Token,
	)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(rb))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", GetBotPk())

	res, err := HttpCl.Do(req) // todo: cant use _ ??????
	if err != nil {
		return err
	}
	if res != nil {
	}

	return nil
}
