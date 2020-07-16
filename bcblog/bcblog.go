package bcblog

import (
	"bcbtest/config"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"sync"
)

var (
	once   sync.Once
	logger *logrus.Logger
)

func Init() {
	once.Do(
		func() {
			logger = &logrus.Logger{}
		})

	cfg := config.GetConfig()
	logger = logrus.New()
	file, _ := rotatelogs.New(cfg.ServerLogFile + ".%Y%m%d")
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
