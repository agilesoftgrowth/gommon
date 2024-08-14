package gommonfx

import (
	"io"

	"github.com/agilesoftgrowth/gommon/logger"
	"go.uber.org/fx"
)

type LoggerServiceParams struct {
	Output io.Writer
	Format logger.LoggerFormat
	Level  logger.LoggerLevel
}

type LoggerServiceResult struct {
	fx.Out
	Logger logger.LoggerService
}

func NewLoggerService(params LoggerServiceParams) (LoggerServiceResult, error) {
	logger := logger.NewLoggerService(params.Output, params.Format, params.Level)
	return LoggerServiceResult{Logger: logger}, nil
}
