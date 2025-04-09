package main

import (
	"fmt"
	"os"

	"github.com/andreis3/users-ms/internal/infra/commons/logger"
	"github.com/andreis3/users-ms/internal/infra/configs"
	"github.com/andreis3/users-ms/internal/infra/server"
	"github.com/andreis3/users-ms/internal/util"
)

func main() {
	log := logger.NewLogger()
	conf, err := configs.LoadConfig()
	if err != nil {
		log.ErrorText(fmt.Sprintf("Notification Errors loading config: %s", err.Error()))
		os.Exit(util.ExitFailure)
	}
	server.Start(conf, *log)
}
