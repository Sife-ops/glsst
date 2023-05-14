package lib

import (
	"bytes"
	"crypto/ed25519"
	"encoding/hex"
)

type VerifyBotError struct {
	Message string
}

func (e *VerifyBotError) Error() string {
	return e.Message
}

type VerifyInteractionInput struct {
	PublicKey string
	Timestamp string
	Signature string
	Body      string
}

// https://github.com/bwmarrin/discordgo/blob/v0.27.1/interactions.go#L572
func VerifyInteraction(request VerifyInteractionInput) (bool, error) {
	key, err := hex.DecodeString(request.PublicKey)
	if err != nil {
		return false, &VerifyBotError{Message: "failed to decode public key hex"} // todo: return error?
	}

	timestamp := request.Timestamp
	if timestamp == "" {
		return false, &VerifyBotError{Message: "missing signature timestamp"}
	}

	signature := request.Signature
	if signature == "" {
		return false, &VerifyBotError{Message: "missing signature"}
	}
	sig, err := hex.DecodeString(signature)
	if err != nil {
		return false, &VerifyBotError{Message: "failed to decude signature hex"}
	}
	if len(sig) != 64 {
		return false, &VerifyBotError{Message: "invalid signature length"}
	}

	var msg bytes.Buffer
	msg.WriteString(timestamp)
	msg.WriteString(request.Body)

	if ed25519.Verify(key, msg.Bytes(), sig) {
		return true, nil
	}

	return false, nil
}
