package main

import (
	"log"

	"github.com/amitrai48/logger"
)

func main() {
	config := logger.Configuration{
		EnableConsole:     true,
		ConsoleLevel:      logger.Debug,
		ConsoleJSONFormat: false,
		EnableFile:        true,
		FileLevel:         logger.Debug,
		FileJSONFormat:    false,
		FileLocation:      "log.log",
	}
	err := logger.NewLogger(config, logger.InstanceZapLogger)
	if err != nil {
		log.Fatalf("Could not instantiate log %s", err.Error())
	}

	contextLogger := logger.WithFields(logger.Fields{"animal": "raiju"})
	contextLogger.Debugf("Starting with zap")
	contextLogger.Infof("Zap is awesome")

	err = logger.NewLogger(config, logger.InstanceLogrusLogger)
	if err != nil {
		log.Fatalf("Could not instantiate log %s", err.Error())
	}
	contextLogger = logger.WithFields(logger.Fields{"animal": "walrus"})
	contextLogger.Debugf("Starting with logrus")
	contextLogger.Infof("Logrus is awesome")
}
