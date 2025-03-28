package postegres

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/andreis3/users-ms/internal/app/interfaces"
	"github.com/andreis3/users-ms/internal/infra/commons/configs"
	"github.com/andreis3/users-ms/internal/infra/commons/logger"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	singleton sync.Once
	pool      *pgxpool.Pool
)

type Postgres struct {
	pool *pgxpool.Pool
}

func NewPoolConnections(conf *configs.Configs) *Postgres {
	log := logger.NewLogger()
	singleton.Do(func() {
		connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			conf.PostgresHost, conf.PostgresPort, conf.PostgresUser, conf.PostgresPassword, conf.PostgresDBName)

		connConfig, err := pgxpool.ParseConfig(connStr)
		if err != nil {
			log.CriticalText(fmt.Sprintf("NotificationErrors parsing connection string: %v", err))
		}

		connConfig.MinConns = conf.PostgresMinConnections
		connConfig.MaxConns = conf.PostgresMaxConnections
		connConfig.MaxConnIdleTime = conf.PostgresMaxConnLifetime
		connConfig.MaxConnIdleTime = conf.PostgresMaxConnIdleTime
		connConfig.HealthCheckPeriod = 1 * time.Minute
		connConfig.ConnConfig.RuntimeParams["application_name"] = conf.ApplicationName

		pool, err = pgxpool.NewWithConfig(context.Background(), connConfig)
		if err != nil {
			log.ErrorText(fmt.Sprintf("NotificationsErrors creating connection poll: %v", err))
			os.Exit(1)
		}
	})

	return &Postgres{pool: pool}
}

func (p *Postgres) InstanceDB() any {
	return p.pool
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

type Queries struct {
	interfaces.InstructionDB
}

func (p *Postgres) Close() {
	pool.Close()
}

func New(db interfaces.InstructionDB) *Queries {
	return &Queries{db}
}
