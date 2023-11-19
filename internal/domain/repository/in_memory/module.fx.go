package in_memory

import "go.uber.org/fx"

var Module = fx.Module(
	"In-memory Repository",
	fx.Provide(
		NewTicketRepository,
		NewUserRepository,
		NewTrainRepository,
	),
)
