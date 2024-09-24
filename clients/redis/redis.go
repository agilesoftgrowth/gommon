package redis

import (
	"github.com/agilesoftgrowth/gommon/logger"
	"github.com/go-redis/redis"
)

func NewRedis(
	logger logger.LoggerService,
	url string,
) (*redis.Client, error) {
	opts, err := redis.ParseURL(url)
	if err != nil {
		return nil, err
	}

	return redis.NewClient(opts), nil
}
