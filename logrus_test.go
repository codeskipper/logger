package logger

import (
	"testing"

	"github.com/sirupsen/logrus"
	hooks "github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestLogrusEntryF(t *testing.T) {
	config := Configuration{
		EnableConsole:     true,
		ConsoleLevel:      "debug",
		ConsoleJSONFormat: false,
	}

	tests := []struct {
		name    string
		level   logrus.Level
		entries int
		log     func(Logger, string)
		message string
	}{
		{
			name:    "Debugf Test",
			level:   logrus.DebugLevel,
			entries: 1,
			log: func(log Logger, message string) {
				log.Debugf(message)
			},
			message: "Logging Debugf Level Logs",
		},
		{
			name:    "Infof Test",
			level:   logrus.InfoLevel,
			entries: 1,
			log: func(log Logger, message string) {
				log.Infof(message)
			},
			message: "Logging Infof Level Logs",
		},
		{
			name:    "Warnf Test",
			level:   logrus.WarnLevel,
			entries: 1,
			log: func(log Logger, message string) {
				log.Warnf(message)
			},
			message: "Logging Warnf Level Logs",
		},
		{
			name:    "Errorf Test",
			level:   logrus.ErrorLevel,
			entries: 1,
			log: func(log Logger, message string) {
				log.Errorf(message)
			},
			message: "Logging Errorf Level Logs",
		},
	}

	assert := assert.New(t)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l, err := newLogrusLogger(config)
			assert.NoError(err)
			log := l.GetLogger().(*logrus.Logger)
			hook := hooks.NewLocal(log)
			test.log(l, test.message)
			assert.Equal(test.entries, len(hook.Entries))
			assert.Equal(test.level, hook.LastEntry().Level)
			assert.Equal(test.message, hook.LastEntry().Message)
		})
	}
}

func TestGetLogLevel(t *testing.T) {
	tests := []struct {
		name             string
		consoleLevel     string
		fileLevel        string
		isError          bool
		expectedLogLevel logrus.Level
	}{
		{
			name:             "Use Console Level log",
			consoleLevel:     "debug",
			fileLevel:        "warn",
			isError:          false,
			expectedLogLevel: logrus.DebugLevel,
		},
		{
			name:             "Use File Level log",
			consoleLevel:     "",
			fileLevel:        "debug",
			isError:          false,
			expectedLogLevel: logrus.DebugLevel,
		},
		{
			name:         "File and Console Level is invalid",
			consoleLevel: "invalid",
			fileLevel:    "invalid",
			isError:      true,
		},
	}

	assert := assert.New(t)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualLogLevel, err := getLogLevel(test.consoleLevel, test.fileLevel)
			if test.isError {
				assert.Error(err)
			} else {
				assert.NoError(err)
				assert.Equal(test.expectedLogLevel, actualLogLevel)
			}
		})
	}
}

func TestFormat(t *testing.T) {
	tests := []struct {
		name              string
		isJSONFormat      bool
		expectedLogFormat logrus.Formatter
	}{
		{
			name:              "Use JSON Format",
			isJSONFormat:      true,
			expectedLogFormat: &logrus.JSONFormatter{},
		},
		{
			name:         "Use Text Format",
			isJSONFormat: false,
			expectedLogFormat: &logrus.TextFormatter{
				FullTimestamp:          true,
				DisableLevelTruncation: true,
			},
		},
	}

	assert := assert.New(t)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualLogFormat := getFormatter(test.isJSONFormat)
			assert.Equal(test.expectedLogFormat, actualLogFormat)
		})
	}
}
