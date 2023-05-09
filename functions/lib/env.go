package lib

import (
	"os"
)

type env struct {
	BotPk      string
	ConsumerFn string
}

var Env *env

func init() {
	Env = &env{
		ConsumerFn: os.Getenv("CONSUMER_FN"),
		BotPk:      os.Getenv("BOT_PUBLIC_KEY"),
	}
}

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
var GetBotPk = getEnvFac("BOT_PUBLIC_KEY")
var GetConsumerFn = getEnvFac("CONSUMER_FN")
var GetTableName = getEnvFac("TABLE_NAME")
