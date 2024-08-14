package gommonfx

import (
	"os"
	"testing"

	"github.com/agilesoftgrowth/gommon/crypto"
	"github.com/agilesoftgrowth/gommon/logger"
	"github.com/stretchr/testify/suite"
)

type CryptoServiceSuite struct {
	suite.Suite
}

func TestCryptoServiceSuite(t *testing.T) {
	suite.Run(t, new(CryptoServiceSuite))
}

func (suite *CryptoServiceSuite) TestNew() {
	logger := logger.NewLoggerService(os.Stdout, logger.FormatText, logger.LevelInfo)
	cryptoService, err := NewCryptoService(CryptoServiceParams{
		Logger: logger,
		Key:    "123456789",
	})
	expected := CryptoServiceResult{
		CryptoService: crypto.NewCryptoService(logger, "123456789"),
	}

	suite.Nil(err)
	suite.Equal(expected, cryptoService)
}
