package server

import "go.uber.org/fx"

var Module = fx.Module(
	"Server Adapter",
	fx.Provide(
		NewServerAdapter,
	),
)
