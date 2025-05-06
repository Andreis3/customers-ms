package postegres

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	dbtracer "github.com/amirsalarsafaei/sqlc-pgx-monitoring/dbtracer"
	"github.com/andreis3/customers-ms/internal/domain/interfaces"
	"github.com/andreis3/customers-ms/internal/infra/commons/logger"
	"github.com/andreis3/customers-ms/internal/infra/configs"
	"github.com/andreis3/customers-ms/internal/util"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.opentelemetry.io/otel"
)

var (
	singleton sync.Once
	pool      *pgxpool.Pool
)

type Postgres struct {
	pool *pgxpool.Pool
}

func NewPoolConnections(conf *configs.Configs, metrics interfaces.Prometheus) *Postgres {
	log := logger.NewLogger()
	singleton.Do(func() {
		connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			conf.PostgresHost, conf.PostgresPort, conf.PostgresUser, conf.PostgresPassword, conf.PostgresDBName)

		connConfig, err := pgxpool.ParseConfig(connStr)
		if err != nil {
			log.CriticalText(fmt.Sprintf("NotificationErrors parsing connection string: %v", err))
		}

		slogLogger := log.SlogJSON()

		// integration opentelemetry
		tracer, err := dbtracer.NewDBTracer(
			conf.PostgresDBName,
			dbtracer.WithLogger(slogLogger),
			dbtracer.WithTraceProvider(otel.GetTracerProvider()),
			dbtracer.WithMeterProvider(metrics.MeterProvider()),
			dbtracer.WithLogArgs(false),
			dbtracer.WithIncludeSQLText(false),
			dbtracer.WithLogArgsLenLimit(1000),
		)
		if err != nil {
			log.ErrorText(fmt.Sprintf("NotificationsErrors creating connection poll: %v", err))
			os.Exit(util.ExitFailure)
		}

		connConfig.ConnConfig.Tracer = tracer

		// connConfig.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeExec
		// connConfig.ConnConfig.StatementCacheCapacity = 200

		connConfig.MinConns = conf.PostgresMinConnections
		connConfig.MaxConns = conf.PostgresMaxConnections
		connConfig.MaxConnIdleTime = conf.PostgresMaxConnLifetime
		connConfig.MaxConnIdleTime = conf.PostgresMaxConnIdleTime
		connConfig.HealthCheckPeriod = 15 * time.Second
		connConfig.ConnConfig.RuntimeParams["application_name"] = conf.ApplicationName

		pool, err = pgxpool.NewWithConfig(context.Background(), connConfig)
		if err != nil {
			log.ErrorText(fmt.Sprintf("NotificationsErrors creating connection poll: %v", err))
			os.Exit(util.ExitFailure)
		}
	})

	return &Postgres{pool: pool}
}

func (p *Postgres) Instance() any {
	return p.pool
}

func (p *Postgres) Close() {
	p.pool.Close()
}

func (p *Postgres) Exec(ctx context.Context, sql string, arguments ...any) (commandtag pgconn.CommandTag, err error) {
	return p.pool.Exec(ctx, sql, arguments...)
}

func (p *Postgres) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	return p.pool.Query(ctx, sql, args...)
}

func (p *Postgres) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	return p.pool.QueryRow(ctx, sql, args...)
}

func (p *Postgres) SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults {
	return p.pool.SendBatch(ctx, b)
}

type Queries struct {
	interfaces.InstructionPostgres
}

func New(db interfaces.InstructionPostgres) *Queries {
	return &Queries{db}
}
