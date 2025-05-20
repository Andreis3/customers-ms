package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/andreis3/customers-ms/internal/infra/adapters/db/postegres"
	"github.com/andreis3/customers-ms/internal/infra/adapters/observability"
	"github.com/andreis3/customers-ms/internal/infra/commons/logger"
	"github.com/andreis3/customers-ms/internal/infra/configs"
	"github.com/andreis3/customers-ms/internal/infra/routes"
	"github.com/andreis3/customers-ms/internal/util"
	"github.com/go-chi/chi/v5"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type Server struct {
	HTTPServer *http.Server
	Postgres   *postegres.Postgres
	Log        logger.Logger
	Prometheus *observability.Prometheus
}

func NewServer(conf *configs.Configs, log *logger.Logger) *Server {
	start := time.Now()

	observability.InitTracer()
	prometheus := observability.NewPrometheus()
	pool := postegres.NewPoolConnections(conf, prometheus)

	mux := chi.NewRouter()

	// OpenTelemetry Middleware
	mux.Use(func(next http.Handler) http.Handler {
		return otelhttp.NewHandler(next, "customers-ms")
	})

	routes.SetupRoutes(mux, pool, log, prometheus, conf)

	server := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%s", conf.ServerPort),
		Handler: mux,
	}

	log.InfoText("[Server] ", "SERVER_STARTED", fmt.Sprintf("Server started in %s", time.Since(start)))
	log.InfoText("[Server] ", "SERVER_STARTED", fmt.Sprintf("Server address http://localhost:%s", conf.ServerPort))

	return &Server{
		HTTPServer: server,
		Postgres:   pool,
		Log:        *log,
		Prometheus: prometheus,
	}
}

func (s *Server) Start() {
	if err := s.HTTPServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		s.Log.CriticalText("[Server] ", "SERVER_ERROR", err.Error())
		os.Exit(util.ExitFailure)
	}
}

func (s *Server) GracefulShutdown() {
	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-shutdownSignal
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	s.Log.InfoText("[Server] ", "SERVER_SHUTDOWN", "Server is shutting down...")

	if err := s.HTTPServer.Shutdown(ctx); err != nil {
		s.Log.ErrorText("[Server] ", "SERVER_SHUTDOWN", err.Error())
	}
	s.Log.InfoText("Closing postgres connection...")
	s.Postgres.Close()
	s.Log.InfoText("Closing prometheus...")
	s.Prometheus.Close()
	s.Log.InfoText("Shutdown complete exit code 0...")

	os.Exit(util.ExitSuccess)
}
