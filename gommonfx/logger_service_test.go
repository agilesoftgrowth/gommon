package gommonfx

import (
	"os"
	"testing"

	"github.com/agilesoftgrowth/gommon/logger"
	"github.com/stretchr/testify/suite"
)

type LoggerServiceSuite struct {
	suite.Suite
}

func TestLoggerServiceSuite(t *testing.T) {
	suite.Run(t, new(LoggerServiceSuite))
}

func (suite *LoggerServiceSuite) TestNew() {
	log, err := NewLoggerService(LoggerServiceParams{
		Output: os.Stdout,
		Format: logger.FormatText,
		Level:  logger.LevelInfo,
	})
	expected := LoggerServiceResult{
		Logger: logger.NewLoggerService(os.Stdout, logger.FormatText, logger.LevelInfo),
	}

	suite.Nil(err)
	suite.Equal(expected, log)
}
