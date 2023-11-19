package service

import (
	"go.uber.org/fx"
)

var Module = fx.Module(
	"Service",
	fx.Provide(
		NewTicketService,
		NewSectionService,
		NewTrainService,
	),
)
