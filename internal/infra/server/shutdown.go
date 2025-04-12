package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/andreis3/users-ms/internal/domain/interfaces"
	"github.com/andreis3/users-ms/internal/infra/adapters/db/postegres"
	"github.com/andreis3/users-ms/internal/infra/commons/logger"
	"github.com/andreis3/users-ms/internal/util"
)

func gracefulShutdown(
	mux *http.Server,
	pool *postegres.Postgres,
	log logger.Logger,
	prometheus interfaces.Prometheus,
) {
	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-shutdownSignal
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	log.InfoText("[Server] ", "SERVER_SHUTDOWN", "Server is shutting down...")
	if err := mux.Shutdown(ctx); err != nil {
		log.ErrorText("[Server] ", "SERVER_SHUTDOWN", err.Error())
	}
	log.InfoText("Closing postgres connection...")
	pool.Close()
	log.InfoText("Closing prometheus...")
	prometheus.Close()
	log.InfoText("Shutdown complete exit code 0...")
	os.Exit(util.ExitSuccess)
}
