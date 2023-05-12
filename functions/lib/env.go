package lib

import (
	"os"
)

func getEnvFac(s string) func() string {
	return func() string {
		a, b := os.LookupEnv(s)
		switch {
		case b == false:
			fallthrough
		case a == "":
			panic(s + " undefined")
		default:
			return a
		}
	}
}

// todo: use these functions
var GetBotAppId = getEnvFac("BOT_APP_ID")
var GetBotPk = getEnvFac("BOT_PUBLIC_KEY")
var GetConsumerFn = getEnvFac("CONSUMER_FN")
var GetMnemonicFn = getEnvFac("MNEMONIC_FN")
var GetTableName = getEnvFac("TABLE_NAME")
