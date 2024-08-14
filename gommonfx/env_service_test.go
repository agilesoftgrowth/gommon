package gommonfx

import (
	"testing"

	"github.com/agilesoftgrowth/gommon/env"
	"github.com/stretchr/testify/suite"
)

type EnvServiceSuite struct {
	suite.Suite
}

func TestEnvServiceSuite(t *testing.T) {
	suite.Run(t, new(EnvServiceSuite))
}

func (suite *EnvServiceSuite) TestNew() {
	logger, err := NewEnvService()
	expected := EnvServiceResult{
		EnvService: env.NewEnvService(),
	}

	suite.Nil(err)
	suite.Equal(expected, logger)
}
