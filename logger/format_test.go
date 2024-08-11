package logger

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type FormatSuite struct {
	suite.Suite
}

func TestFormatSuite(t *testing.T) {
	suite.Run(t, new(FormatSuite))
}

func (suite *FormatSuite) TestFormatConst() {
	suite.Equal(Format(0), FormatJson)
	suite.Equal(Format(1), FormatText)
}

func (suite *FormatSuite) TestFormatString() {
	tests := []struct {
		format   Format
		expected string
	}{
		{
			format:   Format(FormatText),
			expected: "text",
		},
		{
			format:   Format(FormatJson),
			expected: `json`,
		},
	}

	for _, test := range tests {
		suite.Equal(test.expected, test.format.String())
	}
}
