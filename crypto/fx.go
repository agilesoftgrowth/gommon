package crypto

import (
	"github.com/agilesoftgrowth/gommon/logger"
	"go.uber.org/fx"
)

type CryptoServiceParams struct {
	Logger logger.LoggerService
	Key    string
}

type CryptoServiceResult struct {
	fx.Out
	CryptoService CryptoService
}

func NewCrypto(params CryptoServiceParams) (CryptoServiceResult, error) {
	cryptoService := NewCryptoService(params.Logger, params.Key)
	return CryptoServiceResult{CryptoService: cryptoService}, nil
}
