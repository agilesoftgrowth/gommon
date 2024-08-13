package logger

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type LoggerFormatSuite struct {
	suite.Suite
}

func TestLoggerFormatSuite(t *testing.T) {
	suite.Run(t, new(LoggerFormatSuite))
}

func (suite *LoggerFormatSuite) TestLoggerFormatConst() {
	suite.Equal(LoggerFormat(0), FormatJson)
	suite.Equal(LoggerFormat(1), FormatText)
}

func (suite *LoggerFormatSuite) TestLoggerFormatString() {
	tests := []struct {
		format   LoggerFormat
		expected string
	}{
		{
			format:   LoggerFormat(FormatText),
			expected: "text",
		},
		{
			format:   LoggerFormat(FormatJson),
			expected: `json`,
		},
	}

	for _, test := range tests {
		suite.Equal(test.expected, test.format.String())
	}
}

func (suite *LoggerFormatSuite) TestFormat() {
	tests := []struct {
		format   string
		expected LoggerFormat
	}{
		{
			format:   "text",
			expected: FormatText,
		},
		{
			format:   "json",
			expected: FormatJson,
		},
		{
			format:   "TeXt",
			expected: FormatText,
		},
		{
			format:   "not valid format",
			expected: FormatText,
		},
	}

	for _, test := range tests {
		suite.Equal(test.expected, Format(test.format))
	}
}
