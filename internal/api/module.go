package api

import "go.uber.org/fx"

var Module = fx.Module("api",
	fx.Provide(New),
	fx.Invoke(invoke),
)
