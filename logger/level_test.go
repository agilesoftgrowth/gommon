package logger

import (
	"log/slog"
	"testing"

	"github.com/stretchr/testify/suite"
)

type LevelSuite struct {
	suite.Suite
}

func TestLevelSuite(t *testing.T) {
	suite.Run(t, new(LevelSuite))
}

func (suite *LevelSuite) TestLevelConst() {
	suite.Equal(Level(0), LevelDebug)
	suite.Equal(Level(1), LevelInfo)
	suite.Equal(Level(2), LevelWarn)
	suite.Equal(Level(3), LevelError)
}

func (suite *LevelSuite) TestLevelString() {
	tests := []struct {
		level    Level
		expected string
	}{
		{
			level:    Level(LevelDebug),
			expected: "debug",
		},
		{
			level:    Level(LevelInfo),
			expected: "info",
		},
		{
			level:    Level(LevelWarn),
			expected: "warn",
		},
		{
			level:    Level(LevelError),
			expected: "error",
		},
	}

	for _, test := range tests {
		suite.Equal(test.expected, test.level.String())
	}
}

func (suite *LevelSuite) TestLevelSlog() {
	tests := []struct {
		level    Level
		expected slog.Level
	}{
		{
			level:    Level(LevelDebug),
			expected: slog.LevelDebug,
		},
		{
			level:    Level(LevelInfo),
			expected: slog.LevelInfo,
		},
		{
			level:    Level(LevelWarn),
			expected: slog.LevelWarn,
		},
		{
			level:    Level(LevelError),
			expected: slog.LevelError,
		},
	}

	for _, test := range tests {
		suite.Equal(test.expected, test.level.Slog())
	}
}
