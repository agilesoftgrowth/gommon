package gommonfx

import (
	"github.com/agilesoftgrowth/gommon/env"
	"go.uber.org/fx"
)

type EnvServiceResult struct {
	fx.Out
	EnvService env.EnvService
}

func NewEnvService() (EnvServiceResult, error) {
	envSerivce := env.NewEnvService()
	return EnvServiceResult{EnvService: envSerivce}, nil
}
