package main

import (
	"os"

	"github.com/andreis3/customers-ms/internal/di"
	"github.com/andreis3/customers-ms/internal/infra/commons/logger"
	"github.com/andreis3/customers-ms/internal/infra/configs"
)

func main() {
	conf := configs.LoadConfig()
	log := logger.NewLogger()

	if conf == nil {
		log.CriticalText("Failed to load configuration")
		os.Exit(1)
	}

	srv := di.InitializeServer(conf, *log)

	go srv.Start()

	srv.GracefulShutdown()
}
