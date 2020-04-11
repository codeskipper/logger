package main

import (
	"log"

	"github.com/amitrai48/logger"
)

func main() {
	useZap()
	useLogrus()
}

func useLogrus() {
	fileConfig := logger.LogrusFileConfiguration{
		Enable:     true,
		Level:      logger.Debug,
		JSONFormat: false,
		Path:       "log.log",
	}

	consoleConfig := logger.LogrusConsoleConfiguration{
		Enable:     true,
		Level:      logger.Debug,
		JSONFormat: false,
	}

	config := logger.Configuration{
		logger.LogrusConsoleConfig: consoleConfig,
		logger.LogrusFileConfig:    fileConfig,
	}

	err := logger.NewLogger(config, logger.InstanceLogrusLogger)
	if err != nil {
		log.Fatalf("Could not instantiate log %s", err.Error())
	}
	contextLogger := logger.WithFields(logger.Fields{"animal": "walrus"})
	contextLogger.Debugf("Starting with logrus")
	contextLogger.Infof("Logrus is awesome")
}

func useZap() {
	fileConfig := logger.ZapFileConfiguration{
		Enable:     true,
		Level:      logger.Debug,
		JSONFormat: false,
		Path:       "log.log",
	}

	consoleConfig := logger.ZapConsoleConfiguration{
		Enable:     true,
		Level:      logger.Debug,
		JSONFormat: false,
	}

	config := logger.Configuration{
		logger.ZapConsoleConfig: consoleConfig,
		logger.ZapFileConfig:    fileConfig,
	}

	err := logger.NewLogger(config, logger.InstanceZapLogger)
	if err != nil {
		log.Fatalf("Could not instantiate log %s", err.Error())
	}
	contextLogger := logger.WithFields(logger.Fields{"zap": "thunder"})
	contextLogger.Debugf("Starting with zap")
	contextLogger.Infof("Zap is awesome")
}
