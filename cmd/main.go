package main

import (
	"github.com/andreis3/customers-ms/internal/di"
	"github.com/andreis3/customers-ms/internal/infra/configs"
)

func main() {
	conf := configs.LoadConfig()

	if conf == nil {
		panic("Failed to load configuration")
	}

	srv := di.InitializeServer(conf)

	go srv.Start()

	srv.GracefulShutdown()
}
