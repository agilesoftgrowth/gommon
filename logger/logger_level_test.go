package logger

import (
	"log/slog"
	"testing"

	"github.com/stretchr/testify/suite"
)

type LoggerLevelSuite struct {
	suite.Suite
}

func TestLoggerLevelSuite(t *testing.T) {
	suite.Run(t, new(LoggerLevelSuite))
}

func (suite *LoggerLevelSuite) TestLoggerLevelConst() {
	suite.Equal(LoggerLevel(0), LevelDebug)
	suite.Equal(LoggerLevel(1), LevelInfo)
	suite.Equal(LoggerLevel(2), LevelWarn)
	suite.Equal(LoggerLevel(3), LevelError)
}

func (suite *LoggerLevelSuite) TestLoggerLevelString() {
	tests := []struct {
		level    LoggerLevel
		expected string
	}{
		{
			level:    LoggerLevel(LevelDebug),
			expected: "debug",
		},
		{
			level:    LoggerLevel(LevelInfo),
			expected: "info",
		},
		{
			level:    LoggerLevel(LevelWarn),
			expected: "warn",
		},
		{
			level:    LoggerLevel(LevelError),
			expected: "error",
		},
	}

	for _, test := range tests {
		suite.Equal(test.expected, test.level.String())
	}
}

func (suite *LoggerLevelSuite) TestLoggerLevelSlog() {
	tests := []struct {
		level    LoggerLevel
		expected slog.Level
	}{
		{
			level:    LoggerLevel(LevelDebug),
			expected: slog.LevelDebug,
		},
		{
			level:    LoggerLevel(LevelInfo),
			expected: slog.LevelInfo,
		},
		{
			level:    LoggerLevel(LevelWarn),
			expected: slog.LevelWarn,
		},
		{
			level:    LoggerLevel(LevelError),
			expected: slog.LevelError,
		},
	}

	for _, test := range tests {
		suite.Equal(test.expected, test.level.Slog())
	}
}

func (suite *LoggerLevelSuite) TestLevel() {
	tests := []struct {
		level    string
		expected LoggerLevel
	}{
		{
			level:    "debug",
			expected: LevelDebug,
		},
		{
			level:    "info",
			expected: LevelInfo,
		},
		{
			level:    "warn",
			expected: LevelWarn,
		},
		{
			level:    "error",
			expected: LevelError,
		},
		{
			level:    "InFo",
			expected: LevelInfo,
		},
		{
			level:    "not valid level",
			expected: LevelInfo,
		},
	}

	for _, test := range tests {
		suite.Equal(test.expected, Level(test.level))
	}
}
