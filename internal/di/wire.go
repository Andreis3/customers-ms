//go:build wireinject
// +build wireinject

package di

import (
	"github.com/andreis3/customers-ms/internal/infra/commons/logger"
	"github.com/andreis3/customers-ms/internal/infra/configs"
	"github.com/andreis3/customers-ms/internal/infra/server"
	"github.com/google/wire"
)

func InitializeServer(conf *configs.Configs) *server.Server {
	wire.Build(
		logger.NewLogger,
		server.NewServer,
	)

	return &server.Server{}
}
