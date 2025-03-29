package main

import (
	"fmt"
	"os"

	// _ "net/http/pprof"

	"github.com/andreis3/users-ms/internal/infra/commons/configs"
	"github.com/andreis3/users-ms/internal/infra/commons/logger"
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
	// go func() { http.ListenAndServe("localhost:6061", nil) }()
	server.Start(conf, *log)
}
