package crypto

import (
	"os"
	"testing"

	"github.com/agilesoftgrowth/gommon/logger"
	"github.com/stretchr/testify/suite"
)

type FxSuite struct {
	suite.Suite
}

func TestFxSuite(t *testing.T) {
	suite.Run(t, new(FxSuite))
}

func (suite *FxSuite) TestNew() {
	logger := logger.NewLoggerService(os.Stdout, logger.FormatText, logger.LevelInfo)
	cryptoService, err := New(Params{
		Logger: logger,
		Key:    "123456789",
	})
	expected := Result{
		CryptoService: NewCryptoService(logger, "123456789"),
	}

	suite.Nil(err)
	suite.Equal(expected, cryptoService)
}
