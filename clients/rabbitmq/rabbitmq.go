package rabbitmq

import (
	"github.com/agilesoftgrowth/gommon/logger"
	"github.com/wagslane/go-rabbitmq"
)

func NewRabbitmq(
	logger logger.LoggerService,
	queue string,
	url string,
) (*rabbitmq.Conn, error) {
	return rabbitmq.NewConn(url, rabbitmq.WithConnectionOptionsLogger(logger))
}
