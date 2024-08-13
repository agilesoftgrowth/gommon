package logger

import (
	"io"

	"go.uber.org/fx"
)

const (
	FXEVENT_DECORATED          = "Decorated"
	FXEVENT_INVOKED            = "Invoked"
	FXEVENT_INVOKING           = "Invoking"
	FXEVENT_LOGGER_INITIALIZED = "LoggerInitialized"
	FXEVENT_ON_START_EXECUTED  = "OnStartExecuted"
	FXEVENT_ON_START_EXECUTING = "OnStartExecuting"
	FXEVENT_ON_STOP_EXECUTED   = "OnStopExecuted"
	FXEVENT_ON_STOP_EXECUTING  = "OnStopExecuting"
	FXEVENT_PROVIDED           = "Provided"
	FXEVENT_REPLACED           = "Replaced"
	FXEVENT_ROLLEDBACK         = "RolledBack"
	FXEVENT_ROLLINGBACK        = "RollingBack"
	FXEVENT_RUN                = "Run"
	FXEVENT_STARTED            = "Started"
	FXEVENT_STOPPED            = "Stopped"
	FXEVENT_STOPPING           = "Stopping"
	FXEVENT_SUPPLIED           = "Supplied"
)

type Params struct {
	Output io.Writer
	Format LoggerFormat
	Level  LoggerLevel
}

type Result struct {
	fx.Out
	Logger LoggerService
}

func New(params Params) (Result, error) {
	logger := NewLoggerService(params.Output, params.Format, params.Level)
	return Result{Logger: logger}, nil
}
