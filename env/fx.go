package env

import "go.uber.org/fx"

type EnvResult struct {
	fx.Out
	EnvService EnvService
}

func NewEnv() (EnvResult, error) {
	envSerivce := NewEnvService()
	return EnvResult{EnvService: envSerivce}, nil
}
