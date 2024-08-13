package env

import "go.uber.org/fx"

type Result struct {
	fx.Out
	EnvService EnvService
}

func New() (Result, error) {
	envSerivce := NewEnvService()
	return Result{EnvService: envSerivce}, nil
}
