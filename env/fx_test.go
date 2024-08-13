package env

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type FxSuite struct {
	suite.Suite
}

func TestFxSuite(t *testing.T) {
	suite.Run(t, new(FxSuite))
}

func (suite *FxSuite) TestNewEnv() {
	logger, err := NewEnv()
	expected := EnvResult{
		EnvService: NewEnvService(),
	}

	suite.Nil(err)
	suite.Equal(expected, logger)
}
