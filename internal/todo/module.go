package todo

import "go.uber.org/fx"

var Module = fx.Module("todo",
	fx.Provide(
		NewDB,
		NewServer,
	),
)
