package logger

import (
	"context"
	"fmt"
	"io"
	"log/slog"

	"go.uber.org/fx/fxevent"
)

type LoggerService interface {
	fxevent.Logger
	Debug(text string, args ...any)
	Info(text string, args ...any)
	Warn(text string, args ...any)
	Error(text string, args ...any)

	IsActive(level Level) bool
}

func NewLoggerService(output io.Writer, format Format, level Level, args ...any) LoggerService {
	opts := &slog.HandlerOptions{Level: level.Slog()}
	logger := slog.New(slog.NewTextHandler(output, opts))

	if format.String() == "json" {
		logger = slog.New(slog.NewJSONHandler(output, opts))
	}

	return &loggerService{
		logger: logger.With(args...),
	}
}

type loggerService struct {
	logger *slog.Logger
}

func (s *loggerService) Debug(text string, args ...any) {
	s.logger.Debug(text, args...)
}

func (s *loggerService) Info(text string, args ...any) {
	s.logger.Info(text, args...)
}

func (s *loggerService) Warn(text string, args ...any) {
	s.logger.Warn(text, args...)
}

func (s *loggerService) Error(text string, args ...any) {
	s.logger.Error(text, args...)
}

func (s *loggerService) IsActive(level Level) bool {
	return s.logger.Enabled(context.Background(), level.Slog())
}

func (s *loggerService) LogEvent(e fxevent.Event) {
	var eventType string
	var message string
	var err error

	setLog := func(event string, msg string, e error) {
		eventType = event
		message = msg
		err = e
	}

	switch e := e.(type) {
	case *fxevent.Decorated:
		setLog(FXEVENT_DECORATED, fmt.Sprintf("%s decorated", e.DecoratorName), e.Err)
	case *fxevent.Invoked:
		setLog(FXEVENT_INVOKED, fmt.Sprintf("invoked function '%s'", e.FunctionName), e.Err)
	case *fxevent.Invoking:
		setLog(FXEVENT_INVOKING, fmt.Sprintf("invoking function '%s'", e.FunctionName), nil)
	case *fxevent.LoggerInitialized:
		setLog(FXEVENT_LOGGER_INITIALIZED, fmt.Sprintf("logger initialized with '%s'", e.ConstructorName), e.Err)
	case *fxevent.OnStartExecuted:
		setLog(FXEVENT_ON_START_EXECUTED, fmt.Sprintf("executed function '%s' (OnStart hook)", e.FunctionName), e.Err)
	case *fxevent.OnStartExecuting:
		setLog(FXEVENT_ON_START_EXECUTING, fmt.Sprintf("executing function '%s' (OnStart hook)", e.FunctionName), nil)
	case *fxevent.OnStopExecuted:
		setLog(FXEVENT_ON_STOP_EXECUTED, fmt.Sprintf("executed function '%s' (OnStop hook)", e.FunctionName), e.Err)
	case *fxevent.OnStopExecuting:
		setLog(FXEVENT_ON_STOP_EXECUTING, fmt.Sprintf("executing function '%s' (OnStop hook)", e.FunctionName), nil)
	case *fxevent.Provided:
		setLog(FXEVENT_PROVIDED, fmt.Sprintf("%s provided", e.ConstructorName), e.Err)
	case *fxevent.Replaced:
		setLog(FXEVENT_REPLACED, fmt.Sprintf("%s replaced", e.ModuleName), e.Err)
	case *fxevent.RolledBack:
		setLog(FXEVENT_ROLLEDBACK, "rolled back", e.Err)
	case *fxevent.RollingBack:
		setLog(FXEVENT_ROLLINGBACK, "rolling back", e.StartErr)
	case *fxevent.Run:
		setLog(FXEVENT_RUN, fmt.Sprintf("%s run", e.Name), e.Err)
	case *fxevent.Started:
		setLog(FXEVENT_STARTED, "application started", e.Err)
	case *fxevent.Stopped:
		setLog(FXEVENT_STOPPED, "application stopped", e.Err)
	case *fxevent.Stopping:
		setLog(FXEVENT_STOPPING, fmt.Sprintf("'%s' signal received", e.Signal.String()), nil)
	case *fxevent.Supplied:
		setLog(FXEVENT_SUPPLIED, fmt.Sprintf("%s supplied", e.ModuleName), e.Err)
	}

	if err != nil {
		s.logger.Error("something went wrong", "error", err, "fxevent", eventType)
	} else {
		s.logger.Debug(message, "fxevent", eventType)
	}
}
