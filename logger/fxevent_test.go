package logger

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type FxeventSuite struct {
	suite.Suite
}

func TestFxeventSuite(t *testing.T) {
	suite.Run(t, new(FxeventSuite))
}

func (suite *FxeventSuite) TestConst() {
	suite.Equal("Decorated", FXEVENT_DECORATED)
	suite.Equal("Invoked", FXEVENT_INVOKED)
	suite.Equal("Invoking", FXEVENT_INVOKING)
	suite.Equal("LoggerInitialized", FXEVENT_LOGGER_INITIALIZED)
	suite.Equal("OnStartExecuted", FXEVENT_ON_START_EXECUTED)
	suite.Equal("OnStartExecuting", FXEVENT_ON_START_EXECUTING)
	suite.Equal("OnStopExecuted", FXEVENT_ON_STOP_EXECUTED)
	suite.Equal("OnStopExecuting", FXEVENT_ON_STOP_EXECUTING)
	suite.Equal("Provided", FXEVENT_PROVIDED)
	suite.Equal("Replaced", FXEVENT_REPLACED)
	suite.Equal("RolledBack", FXEVENT_ROLLEDBACK)
	suite.Equal("RollingBack", FXEVENT_ROLLINGBACK)
	suite.Equal("Run", FXEVENT_RUN)
	suite.Equal("Started", FXEVENT_STARTED)
	suite.Equal("Stopped", FXEVENT_STOPPED)
	suite.Equal("Stopping", FXEVENT_STOPPING)
	suite.Equal("Supplied", FXEVENT_SUPPLIED)
}
