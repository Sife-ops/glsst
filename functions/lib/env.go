package lib

import (
	"os"
)

func GetEnvOrPanic(s string) func() string {
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

var GetBotPk = GetEnvOrPanic("BOT_PUBLIC_KEY")
var GetBotConsumerFn = GetEnvOrPanic("CONSUMER_FN")
