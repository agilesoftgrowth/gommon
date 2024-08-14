package gommonfx

import (
	"github.com/agilesoftgrowth/gommon/crypto"
	"github.com/agilesoftgrowth/gommon/logger"
	"go.uber.org/fx"
)

type CryptoServiceParams struct {
	Logger logger.LoggerService
	Key    string
}

type CryptoServiceResult struct {
	fx.Out
	CryptoService crypto.CryptoService
}

func NewCryptoService(params CryptoServiceParams) (CryptoServiceResult, error) {
	cryptoService := crypto.NewCryptoService(params.Logger, params.Key)
	return CryptoServiceResult{CryptoService: cryptoService}, nil
}
