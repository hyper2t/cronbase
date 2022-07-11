package main

import (
	"cronbase/commands"
	_ "cronbase/config/configor"
	_ "cronbase/config/dotenv"
	_ "cronbase/di"
	"github.com/mix-go/dotenv"
	"github.com/mix-go/xcli"
)

func main() {
	xcli.SetName("app").
		SetVersion("0.0.0-alpha").
		SetDebug(dotenv.Getenv("APP_DEBUG").Bool(false))
	xcli.AddCommand(commands.Commands...).Run()
}
