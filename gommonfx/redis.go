package gommonfx

import (
	gommonRedis "github.com/agilesoftgrowth/gommon/clients/redis"
	"github.com/agilesoftgrowth/gommon/logger"
	"github.com/go-redis/redis"
	"go.uber.org/fx"
)

type RedisParams struct {
	Logger logger.LoggerService
	URL    string
}

type RedisResult struct {
	fx.Out
	Client *redis.Client
}

func NewRedis(params RedisParams) (RedisResult, error) {
	client, err := gommonRedis.NewRedis(
		params.Logger,
		params.URL,
	)
	return RedisResult{Client: client}, err
}
