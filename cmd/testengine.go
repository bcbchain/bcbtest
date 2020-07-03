package main

import (
	"bcbtest/config"
	"fmt"
	"github.com/bcbchain/bclib/tendermint/tmlibs/log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("need chainVersion value ")
		os.Exit(1)
	}
	config.ChainVersion = os.Args[1]
	if config.ChainVersion != "1.0" && config.ChainVersion != "2.0" {
		fmt.Println("chainVersion must be 1.0 or 2.0")
		os.Exit(1)
	}

	err := config.LoadConfig("./bundle/config.yaml")
	if err != nil {
		panic(err)
	}

	logger := initLogger()

	fmt.Println(logger)
}

func initLogger() log.Loggerf {
	l := log.NewTMLogger("./log", "bcbtest")
	l.SetOutputToFile(true)
	l.SetOutputToScreen(false)
	l.AllowLevel(config.GetConfig().LogLevel)
	return l
}
