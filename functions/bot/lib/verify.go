package lib

import (
	"bytes"
	"crypto/ed25519"
	"encoding/hex"
	"glsst/functions/lib"

	"github.com/aws/aws-lambda-go/events"
)

// https://github.com/bwmarrin/discordgo/blob/v0.27.1/interactions.go#L572
func VerifyInteraction(request events.APIGatewayV2HTTPRequest) bool {
	key, err := hex.DecodeString(lib.GetBotPk())
	if err != nil {
		return false
	}

	timestamp := request.Headers["x-signature-timestamp"]
	if timestamp == "" {
		return false
	}

	signature := request.Headers["x-signature-ed25519"]
	if signature == "" {
		return false
	}
	sig, err := hex.DecodeString(signature)
	if err != nil {
		return false
	}
	if len(sig) != 64 {
		return false
	}

	var msg bytes.Buffer
	msg.WriteString(timestamp)
	msg.WriteString(request.Body)

	return ed25519.Verify(key, msg.Bytes(), sig)
}
