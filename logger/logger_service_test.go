package logger

import (
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

type LoggerServiceSuite struct {
	suite.Suite
}

func TestLoggerServiceSuite(t *testing.T) {
	suite.Run(t, new(LoggerServiceSuite))
}

func (suite *LoggerServiceSuite) TestNewLoggerServiceFormat() {
	tests := []struct {
		format   Format
		expected string
	}{
		{
			format:   Format(Text),
			expected: "time=.* level=INFO msg=foobar",
		},
		{
			format:   Format(Json),
			expected: `{"time":".*","level":"INFO","msg":"foobar"}`,
		},
	}

	for _, test := range tests {
		buf := new(strings.Builder)
		logger := NewLoggerService(buf, test.format, Level(LevelInfo))
		logger.Info("foobar")

		suite.Regexp(test.expected, buf.String())
	}
}

func (suite *LoggerServiceSuite) TestNewLoggerServiceLevel() {
	tests := []struct {
		level Level
	}{
		{level: LevelDebug},
		{level: LevelInfo},
		{level: LevelWarn},
		{level: LevelError},
	}

	for _, test := range tests {
		logger := NewLoggerService(io.Discard, Text, test.level)
		suite.True(logger.IsActive(test.level))
	}
}

func (suite *LoggerServiceSuite) TestLoggerMethod() {
	tests := []struct {
		methodName string
		expected   string
	}{
		{
			methodName: "Debug",
			expected:   "time=.* level=DEBUG msg=foobar",
		},
		{
			methodName: "Info",
			expected:   "time=.* level=INFO msg=foobar",
		},
		{
			methodName: "Warn",
			expected:   "time=.* level=WARN msg=foobar",
		},
		{
			methodName: "Error",
			expected:   "time=.* level=ERROR msg=foobar",
		},
	}

	for _, test := range tests {
		buf := new(strings.Builder)
		logger := NewLoggerService(buf, Text, LevelDebug)
		method := reflect.ValueOf(logger).MethodByName(test.methodName)
		method.Call([]reflect.Value{reflect.ValueOf("foobar")})

		suite.Regexp(test.expected, buf.String())
	}
}
