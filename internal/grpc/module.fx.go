package grpc

import "go.uber.org/fx"

var Module = fx.Module(
	"GRPC",
	fx.Provide(
		NewServer,
	),
)
