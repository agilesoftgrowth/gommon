package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type EnvServiceSuite struct {
	suite.Suite
}

func TestEnvServiceSuite(t *testing.T) {
	suite.Run(t, new(EnvServiceSuite))
}

func (suite *EnvServiceSuite) TestGetString() {
	tests := []struct {
		setEnv   bool
		envName  string
		envValue string
		expected any
	}{
		{
			setEnv:   true,
			envName:  "APP_NAME",
			envValue: "gommon",
			expected: "gommon",
		},
		{
			setEnv:   false,
			envName:  "",
			envValue: "",
			expected: nil,
		},
	}

	for _, test := range tests {
		if test.setEnv {
			suite.T().Setenv(test.envName, test.envValue)
		}
		envService := NewEnvService()
		value := envService.GetString(test.envName)

		if test.expected == nil {
			suite.Nil(value)
		} else {
			suite.Equal(test.expected, *value)
		}
	}
}

func (suite *EnvServiceSuite) TestGetStringDefault() {
	tests := []struct {
		setEnv     bool
		envName    string
		envValue   string
		envDefault string
		expected   any
	}{
		{
			setEnv:     true,
			envName:    "APP_NAME",
			envValue:   "gommon",
			envDefault: "gommonDefault",
			expected:   "gommon",
		},
		{
			setEnv:     false,
			envName:    "APP_NAME",
			envValue:   "",
			envDefault: "gommonDefault",
			expected:   "gommonDefault",
		},
	}

	for _, test := range tests {
		if test.setEnv {
			suite.T().Setenv(test.envName, test.envValue)
		}
		envService := NewEnvService()
		value := envService.GetStringDefault(test.envName, test.envDefault)
		suite.Equal(test.expected, value)

		if test.setEnv {
			os.Unsetenv(test.envName)
		}
	}
}

func (suite *EnvServiceSuite) TestGetInt() {
	tests := []struct {
		setEnv   bool
		envName  string
		envValue string
		expected any
	}{
		{
			setEnv:   true,
			envName:  "MAX_RETRY",
			envValue: "10",
			expected: 10,
		},
		{
			setEnv:   false,
			envName:  "",
			envValue: "",
			expected: nil,
		},
	}

	for _, test := range tests {
		if test.setEnv {
			suite.T().Setenv(test.envName, test.envValue)
		}
		envService := NewEnvService()
		value := envService.GetInt(test.envName)

		if test.expected == nil {
			suite.Nil(value)
		} else {
			suite.Equal(test.expected, *value)
		}
	}
}

func (suite *EnvServiceSuite) TestGetIntDefault() {
	tests := []struct {
		setEnv     bool
		envName    string
		envValue   string
		envDefault int
		expected   any
	}{
		{
			setEnv:     true,
			envName:    "MAX_RETRY",
			envValue:   "10",
			envDefault: 1,
			expected:   10,
		},
		{
			setEnv:     false,
			envName:    "MAX_RETRY",
			envValue:   "",
			envDefault: 1,
			expected:   1,
		},
	}

	for _, test := range tests {
		if test.setEnv {
			suite.T().Setenv(test.envName, test.envValue)
		}
		envService := NewEnvService()
		value := envService.GetIntDefault(test.envName, test.envDefault)
		suite.Equal(test.expected, value)

		if test.setEnv {
			os.Unsetenv(test.envName)
		}
	}
}

func (suite *EnvServiceSuite) TestGetBool() {
	tests := []struct {
		setEnv   bool
		envName  string
		envValue string
		expected any
	}{
		{
			setEnv:   true,
			envName:  "USE_SSL",
			envValue: "true",
			expected: true,
		},
		{
			setEnv:   false,
			envName:  "",
			envValue: "",
			expected: nil,
		},
	}

	for _, test := range tests {
		if test.setEnv {
			suite.T().Setenv(test.envName, test.envValue)
		}
		envService := NewEnvService()
		value := envService.GetBool(test.envName)

		if test.expected == nil {
			suite.Nil(value)
		} else {
			suite.Equal(test.expected, *value)
		}
	}
}

func (suite *EnvServiceSuite) TestGetBoolDefault() {
	tests := []struct {
		setEnv     bool
		envName    string
		envValue   string
		envDefault bool
		expected   any
	}{
		{
			setEnv:     true,
			envName:    "USE_SSL",
			envValue:   "true",
			envDefault: false,
			expected:   true,
		},
		{
			setEnv:     false,
			envName:    "USE_SSL",
			envValue:   "",
			envDefault: false,
			expected:   false,
		},
	}

	for _, test := range tests {
		if test.setEnv {
			suite.T().Setenv(test.envName, test.envValue)
		}
		envService := NewEnvService()
		value := envService.GetBoolDefault(test.envName, test.envDefault)
		suite.Equal(test.expected, value)

		if test.setEnv {
			os.Unsetenv(test.envName)
		}
	}
}
