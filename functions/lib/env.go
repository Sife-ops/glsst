package lib

import (
	"os"
)

type env struct {
	BotPk      string
	ConsumerFn string
}

var Env *env

func getEnvOrPanic(s string) string {
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

func init() {
	Env = &env{
		ConsumerFn: getEnvOrPanic("CONSUMER_FN"),
		BotPk:      getEnvOrPanic("BOT_PUBLIC_KEY"),
	}
}
