package domain

import (
	"go.uber.org/fx"

	"github.com/mcampo84/cloudbees-assignment/internal/adapter/grpc"
	serviceAdapter "github.com/mcampo84/cloudbees-assignment/internal/domain/adapter/service"
	"github.com/mcampo84/cloudbees-assignment/internal/domain/repository/in_memory"
	"github.com/mcampo84/cloudbees-assignment/internal/domain/service"
)

var Module = fx.Module(
	"Domain",
	in_memory.Module,
	serviceAdapter.Module,
	service.Module,
	grpc.Module,
)
