package server

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/andreis3/users-ms/internal/infra/adapters/db/postegres"
	"github.com/andreis3/users-ms/internal/infra/commons/configs"
	"github.com/andreis3/users-ms/internal/infra/commons/logger"
	"github.com/andreis3/users-ms/internal/infra/setup"
	"github.com/andreis3/users-ms/internal/util"
	"github.com/go-chi/chi/v5"
)

func Start(conf *configs.Configs, log logger.Logger) {
	start := time.Now()

	mux := chi.NewRouter()
	server := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%s", conf.ServerPort),
		Handler: mux,
	}

	pool := postegres.NewPoolConnections(conf)

	go func() {
		setup.SetupRoutes(mux, pool, &log)
		end := time.Since(start)
		log.InfoText("[Server] ", "SERVER_STARTED", fmt.Sprintf("Server started in %s", end.String()))
		log.InfoText("[Server] ", "SERVER_STARTED", fmt.Sprintf("Server is listening on port %s", conf.ServerPort))

		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.CriticalText("[Server] ", "SERVER_ERROR", err.Error())
			os.Exit(util.ExitFailure)
		}
	}()

	gracefulShutdown(server, pool, log)
}
