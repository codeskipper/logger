package logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

type logrusLogEntry struct {
	entry *logrus.Entry
}

type logrusLogger struct {
	logger *logrus.Logger
}

type logrusLogLevel struct {
	level logrus.Level
}

func newLogrusLogger(config Configuration) (Logger, error) {
	l, err := getLogLevel(config.ConsoleLevel, config.FileLevel)
	if err != nil {
		return nil, err
	}

	lLogger := &logrus.Logger{
		Out:       os.Stdout,
		Formatter: getFormatter(config.ConsoleJSONFormat),
		Hooks:     make(logrus.LevelHooks),
		Level:     l.level,
	}

	log := &logrusLogger{
		logger: lLogger,
	}

	log.setOutput(config.EnableConsole, config.EnableFile, config.FileJSONFormat, config.FileLocation)
	return log, nil
}

func getLogLevel(consoleLevel, filelevel string) (*logrusLogLevel, error) {
	logLevel := consoleLevel
	if logLevel == "" {
		logLevel = filelevel
	}

	l, err := logrus.ParseLevel(logLevel)
	if err != nil {
		return nil, err
	}

	return &logrusLogLevel{
		level: l,
	}, nil
}

func getFormatter(isJSON bool) logrus.Formatter {
	if isJSON {
		return &logrus.JSONFormatter{}
	}
	return &logrus.TextFormatter{
		FullTimestamp:          true,
		DisableLevelTruncation: true,
	}
}

func (l *logrusLogger) setOutput(enableConsole, enableFile, isJSON bool, fileLocation string) {
	fileHandler := &lumberjack.Logger{
		Filename: fileLocation,
		MaxSize:  100,
		Compress: true,
		MaxAge:   28,
	}

	if enableConsole && enableFile {
		l.logger.SetOutput(io.MultiWriter(l.logger.Out, fileHandler))
	} else {
		if enableFile {
			l.logger.SetOutput(fileHandler)
			l.logger.SetFormatter(getFormatter(isJSON))
		}
	}
}

func (l *logrusLogger) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

func (l *logrusLogger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

func (l *logrusLogger) Warnf(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

func (l *logrusLogger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

func (l *logrusLogger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}

func (l *logrusLogger) Panicf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args...)
}

func (l *logrusLogger) WithFields(fields Fields) Logger {
	return &logrusLogEntry{
		entry: l.logger.WithFields(convertToLogrusFields(fields)),
	}
}

func (l *logrusLogger) GetLogger() interface{} {
	return l.logger
}

func (l *logrusLogEntry) Debugf(format string, args ...interface{}) {
	l.entry.Debugf(format, args...)
}

func (l *logrusLogEntry) Infof(format string, args ...interface{}) {
	l.entry.Infof(format, args...)
}

func (l *logrusLogEntry) Warnf(format string, args ...interface{}) {
	l.entry.Warnf(format, args...)
}

func (l *logrusLogEntry) Errorf(format string, args ...interface{}) {
	l.entry.Errorf(format, args...)
}

func (l *logrusLogEntry) Fatalf(format string, args ...interface{}) {
	l.entry.Fatalf(format, args...)
}

func (l *logrusLogEntry) Panicf(format string, args ...interface{}) {
	l.entry.Fatalf(format, args...)
}

func (l *logrusLogEntry) WithFields(fields Fields) Logger {
	return l.WithFields(fields)
}

func (l *logrusLogEntry) GetLogger() interface{} {
	return l.entry
}

func convertToLogrusFields(fields Fields) logrus.Fields {
	logrusFields := logrus.Fields{}
	for index, val := range fields {
		logrusFields[index] = val
	}
	return logrusFields
}
