package di

import (
	"github.com/andreis3/customers-ms/internal/infra/commons/logger"
	"github.com/andreis3/customers-ms/internal/infra/configs"
	"github.com/andreis3/customers-ms/internal/infra/server"
)

func InitializeServer(conf *configs.Configs, log logger.Logger) *server.Server {
	server := server.NewServer(conf, log)
	return server
}
