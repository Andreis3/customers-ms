package main

import (
	"os"

	"github.com/andreis3/customers-ms/internal/infra/commons/logger"
	"github.com/andreis3/customers-ms/internal/infra/configs"
	"github.com/andreis3/customers-ms/internal/infra/server"
)

func main() {
	conf := configs.LoadConfig()
	log := logger.NewLogger()

	if conf == nil {
		log.CriticalText("Failed to load configuration")
		os.Exit(1)
	}

	server := server.NewServer(conf, *log)

	go server.Start()

	server.GracefulShutdown()
}
