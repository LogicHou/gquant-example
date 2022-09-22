package bootstrap

import (
	"flag"
	"log"

	"go.uber.org/zap"
)

var logFilePath = flag.String("l", "trade.log", "log file path")

func InitializeLog() *zap.Logger {
	loggerCfg := zap.NewDevelopmentConfig()
	loggerCfg.OutputPaths = []string{"stdout", *logFilePath}
	logger, err := loggerCfg.Build()
	if err != nil {
		log.Fatalf("cannot create logger: %v", err)
	}

	return logger
}
