package service

import "go.uber.org/fx"

var Module = fx.Module(
	"Service Adapters",
	fx.Provide(
		NewTicketRepositoryAdapter,
	),
)
