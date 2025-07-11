package main

import (
	"os"

	"github.com/andreis3/customers-ms/internal/infra/adapters/logger"
	"github.com/andreis3/customers-ms/internal/infra/configs"
	"github.com/andreis3/customers-ms/internal/infra/server/web"
)

func main() {
	conf := configs.LoadConfig()
	log := logger.NewLogger()

	if conf == nil {
		log.CriticalText("Failed to load configuration")
		os.Exit(1)
	}

	serverWeb := web.NewServer(conf, *log)

	go serverWeb.Start()

	serverWeb.GracefulShutdown()
}
