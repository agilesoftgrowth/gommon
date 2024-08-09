package logger

import (
	"context"
	"io"
	"log/slog"
)

type LoggerService interface {
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

	return loggerService{
		logger: logger.With(args...),
	}
}

type loggerService struct {
	logger *slog.Logger
}

func (s loggerService) Debug(text string, args ...any) {
	s.logger.Debug(text, args...)
}

func (s loggerService) Info(text string, args ...any) {
	s.logger.Info(text, args...)
}

func (s loggerService) Warn(text string, args ...any) {
	s.logger.Warn(text, args...)
}

func (s loggerService) Error(text string, args ...any) {
	s.logger.Error(text, args...)
}

func (s loggerService) IsActive(level Level) bool {
	return s.logger.Enabled(context.Background(), level.Slog())
}
