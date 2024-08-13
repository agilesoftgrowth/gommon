package logger

import (
	"errors"
	"io"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.uber.org/fx/fxevent"
)

type LoggerServiceSuite struct {
	suite.Suite
}

func TestLoggerServiceSuite(t *testing.T) {
	suite.Run(t, new(LoggerServiceSuite))
}

func (suite *LoggerServiceSuite) TestNewLoggerServiceFormat() {
	tests := []struct {
		format   Format
		expected string
	}{
		{
			format:   Format(FormatText),
			expected: "time=.* level=INFO msg=foobar",
		},
		{
			format:   Format(FormatJson),
			expected: `{"time":".*","level":"INFO","msg":"foobar"}`,
		},
	}

	for _, test := range tests {
		buf := new(strings.Builder)
		logger := NewLoggerService(buf, test.format, Level(LevelInfo))
		logger.Info("foobar")

		suite.Regexp(test.expected, buf.String())
	}
}

func (suite *LoggerServiceSuite) TestNewLoggerServiceLevel() {
	tests := []struct {
		level Level
	}{
		{level: LevelDebug},
		{level: LevelInfo},
		{level: LevelWarn},
		{level: LevelError},
	}

	for _, test := range tests {
		logger := NewLoggerService(io.Discard, FormatText, test.level)
		suite.True(logger.IsActive(test.level))
	}
}

func (suite *LoggerServiceSuite) TestLoggerMethod() {
	tests := []struct {
		methodName string
		expected   string
	}{
		{
			methodName: "Debug",
			expected:   "time=.* level=DEBUG msg=foobar",
		},
		{
			methodName: "Info",
			expected:   "time=.* level=INFO msg=foobar",
		},
		{
			methodName: "Warn",
			expected:   "time=.* level=WARN msg=foobar",
		},
		{
			methodName: "Error",
			expected:   "time=.* level=ERROR msg=foobar",
		},
	}

	for _, test := range tests {
		buf := new(strings.Builder)
		logger := NewLoggerService(buf, FormatText, LevelDebug)
		method := reflect.ValueOf(logger).MethodByName(test.methodName)
		method.Call([]reflect.Value{reflect.ValueOf("foobar")})

		suite.Regexp(test.expected, buf.String())
	}
}

func (suite *LoggerServiceSuite) TestLogEvent() {
	tests := []struct {
		event    fxevent.Event
		expected string
	}{
		{
			event:    &fxevent.Decorated{DecoratorName: "foobar"},
			expected: `time=.* level=DEBUG msg="foobar decorated" fxevent=Decorated`,
		},
		{
			event:    &fxevent.Invoked{FunctionName: "foobar"},
			expected: `time=.* level=DEBUG msg="invoked function 'foobar'" fxevent=Invoked`,
		},
		{
			event:    &fxevent.Invoking{FunctionName: "foobar"},
			expected: `time=.* level=DEBUG msg="invoking function 'foobar'" fxevent=Invoking`,
		},
		{
			event:    &fxevent.LoggerInitialized{ConstructorName: "foobar"},
			expected: `time=.* level=DEBUG msg="logger initialized with 'foobar'" fxevent=LoggerInitialized`,
		},
		{
			event:    &fxevent.OnStartExecuted{FunctionName: "foobar"},
			expected: `time=.* level=DEBUG msg="executed function 'foobar' \(OnStart hook\)" fxevent=OnStartExecuted`,
		},
		{
			event:    &fxevent.OnStartExecuting{FunctionName: "foobar"},
			expected: `time=.* level=DEBUG msg="executing function 'foobar' \(OnStart hook\)" fxevent=OnStartExecuting`,
		},
		{
			event:    &fxevent.OnStopExecuted{FunctionName: "foobar"},
			expected: `time=.* level=DEBUG msg="executed function 'foobar' \(OnStop hook\)" fxevent=OnStopExecuted`,
		},
		{
			event:    &fxevent.OnStopExecuting{FunctionName: "foobar"},
			expected: `time=.* level=DEBUG msg="executing function 'foobar' \(OnStop hook\)" fxevent=OnStopExecuting`,
		},
		{
			event:    &fxevent.Provided{ConstructorName: "foobar"},
			expected: `time=.* level=DEBUG msg="foobar provided" fxevent=Provided`,
		},
		{
			event:    &fxevent.Replaced{ModuleName: "foobar"},
			expected: `time=.* level=DEBUG msg="foobar replaced" fxevent=Replaced`,
		},
		{
			event:    &fxevent.RolledBack{},
			expected: `time=.* level=DEBUG msg="rolled back" fxevent=RolledBack`,
		},
		{
			event:    &fxevent.RollingBack{},
			expected: `time=.* level=DEBUG msg="rolling back" fxevent=RollingBack`,
		},
		{
			event:    &fxevent.Run{Name: "foobar"},
			expected: `time=.* level=DEBUG msg="foobar run" fxevent=Run`,
		},
		{
			event:    &fxevent.Started{},
			expected: `time=.* level=DEBUG msg="application started" fxevent=Started`,
		},
		{
			event:    &fxevent.Stopped{},
			expected: `time=.* level=DEBUG msg="application stopped" fxevent=Stopped`,
		},
		{
			event:    &fxevent.Stopping{Signal: os.Interrupt},
			expected: `time=.* level=DEBUG msg="'interrupt' signal received" fxevent=Stopping`,
		},
		{
			event:    &fxevent.Supplied{ModuleName: "foobar"},
			expected: `time=.* level=DEBUG msg="foobar supplied" fxevent=Supplied`,
		},
		{
			event:    &fxevent.Started{Err: errors.New("cannot start application")},
			expected: `time=.* level=ERROR msg="something went wrong" error="cannot start application" fxevent=Started`,
		},
	}

	for _, test := range tests {
		buf := new(strings.Builder)
		logger := NewLoggerService(buf, FormatText, LevelDebug)
		logger.LogEvent(test.event)

		suite.Regexp(test.expected, buf.String())
	}
}
