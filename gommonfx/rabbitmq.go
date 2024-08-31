package gommonfx

import (
	"github.com/agilesoftgrowth/gommon/clients/rabbitmq"
	"github.com/agilesoftgrowth/gommon/logger"
	gorabbit "github.com/wagslane/go-rabbitmq"
	"go.uber.org/fx"
)

type RabbitmqParams struct {
	Logger logger.LoggerService
	Queue  string
	URL    string
}

type RabbitmqResult struct {
	fx.Out
	Conn *gorabbit.Conn
}

func NewRabbitmq(params RabbitmqParams) (RabbitmqResult, error) {
	conn, err := rabbitmq.NewRabbitmq(
		params.Logger,
		params.Queue,
		params.URL,
	)
	return RabbitmqResult{Conn: conn}, err
}
