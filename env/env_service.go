package env

import (
	"os"
	"strconv"
)

type EnvService interface {
	GetString(key string) *string
	GetStringDefault(key string, fallback string) string
	GetInt(key string) *int
	GetIntDefault(key string, fallback int) int
	GetBool(key string) *bool
	GetBoolDefault(key string, fallback bool) bool
}

func NewEnvService() EnvService {
	return &envService{}
}

type envService struct{}

func (s *envService) GetString(key string) *string {
	if value, ok := os.LookupEnv(key); ok {
		return &value
	}
	return nil
}

func (s *envService) GetStringDefault(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func (s *envService) GetInt(key string) *int {
	if value, ok := os.LookupEnv(key); ok {
		if number, err := strconv.ParseInt(value, 10, 64); err == nil {
			intNumber := int(number)
			return &intNumber
		}
	}
	return nil
}

func (s *envService) GetIntDefault(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		if number, err := strconv.ParseInt(value, 10, 64); err == nil {
			return int(number)
		}
	}
	return fallback
}

func (s *envService) GetBool(key string) *bool {
	if value, ok := os.LookupEnv(key); ok {
		if boolean, err := strconv.ParseBool(value); err == nil {
			return &boolean
		}
	}
	return nil
}

func (s *envService) GetBoolDefault(key string, fallback bool) bool {
	if value, ok := os.LookupEnv(key); ok {
		if boolean, err := strconv.ParseBool(value); err == nil {
			return boolean
		}
	}
	return fallback
}
