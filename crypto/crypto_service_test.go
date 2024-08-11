package crypto

import (
	"io"
	"regexp"
	"testing"

	"github.com/agilesoftgrowth/gommon/logger"
	"github.com/stretchr/testify/suite"
)

type CryptoServiceSuite struct {
	suite.Suite
}

func TestCryptoServiceSuite(t *testing.T) {
	suite.Run(t, new(CryptoServiceSuite))
}

func (suite *CryptoServiceSuite) TestEncrypt() {
	text := "text to be encrypted"
	cryptor := NewCryptoService(suite.getLogger(), "1234567890abcdef")
	encText, _ := cryptor.Encrypt(text)
	match, _ := regexp.MatchString("^([A-Za-z0-9+/]{4})*([A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{2}==)?$", encText)
	suite.True(match)
}

func (suite *CryptoServiceSuite) TestEncryptWithBadKey() {
	text := "text to be encrypted"
	cryptor := NewCryptoService(suite.getLogger(), "key too short")
	_, err := cryptor.Encrypt(text)
	suite.ErrorContains(err, "cannot encrypt data")
}

func (suite *CryptoServiceSuite) TestDecrypt() {
	text := "text to be encrypted"
	cryptor := NewCryptoService(suite.getLogger(), "1234567890abcdef")
	encText, _ := cryptor.Encrypt(text)
	decText, _ := cryptor.Decrypt(encText)
	suite.Equal(text, decText)
}

func (suite *CryptoServiceSuite) TestDecryptWithBadKey() {
	text := "text to be encrypted"

	cryptor1 := NewCryptoService(suite.getLogger(), "1234567890abcdef")
	encText, _ := cryptor1.Encrypt(text)

	cryptor2 := NewCryptoService(suite.getLogger(), "key to short")
	_, err := cryptor2.Decrypt(encText)

	suite.ErrorContains(err, "cannot decrypt data")
}

func (suite *CryptoServiceSuite) TestDecryptTextNotBase64Encoded() {
	cryptor := NewCryptoService(suite.getLogger(), "1234567890abcdef")
	_, err := cryptor.Decrypt("not valid base64 encoded text")
	suite.ErrorContains(err, "text is not base64 encoded")
}

func (suite *CryptoServiceSuite) getLogger() logger.LoggerService {
	return logger.NewLoggerService(io.Discard, logger.FormatText, logger.LevelInfo)
}
