package main

import (
	"bcbtest/config"
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"sync"

	//"github.com/bcbchain/bclib/tendermint/tmlibs/log"
	"os"
)

var (
	once   sync.Once
	logger *logrus.Logger
	entry  *logrus.Entry
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

	initLogger()

	logger = GetLogger()

	fmt.Println(logger)
}

func initLogger() {
	once.Do(
		func() {
			logger = &logrus.Logger{}
		})

	cfg := config.GetConfig()
	logger = logrus.New()
	file, _ := rotatelogs.New(cfg.LogFile + ".%Y%m%d")
	logger.SetOutput(file)
	logger.SetLevel(logrus.DebugLevel)

	formatter := &easy.Formatter{
		// 不需要彩色日志
		//DisableColors: false,
		// 定义时间戳格式
		TimestampFormat: "2006-01-02 15:04:05.000",
		LogFormat:       "[%time%] [%lvl%] [%module%] - %msg%\n",
	}
	logger.Formatter = formatter
}

func GetLogger() *logrus.Logger {
	return logger
}
