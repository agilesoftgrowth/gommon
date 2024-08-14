package gommonfx

import "go.uber.org/fx"

type Controller interface {
	Routes() []Route
}

func AsController(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(Controller)),
		fx.ResultTags(`group:"controllers"`),
	)
}
