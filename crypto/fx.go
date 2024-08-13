package crypto

import (
	"github.com/agilesoftgrowth/gommon/logger"
	"go.uber.org/fx"
)

type Params struct {
	Logger logger.LoggerService
	Key    string
}

type Result struct {
	fx.Out
	CryptoService CryptoService
}

func New(params Params) (Result, error) {
	cryptoService := NewCryptoService(params.Logger, params.Key)
	return Result{CryptoService: cryptoService}, nil
}
