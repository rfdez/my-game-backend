package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	mygame "github.com/rfdez/my-game-backend/internal"
	"github.com/rfdez/my-game-backend/internal/checking"
	"github.com/rfdez/my-game-backend/internal/fetcher"
	"github.com/rfdez/my-game-backend/internal/incrementer"
	"github.com/rfdez/my-game-backend/internal/platform/bus/inmemory"
	"github.com/rfdez/my-game-backend/internal/platform/logger/zerolog"
	"github.com/rfdez/my-game-backend/internal/platform/server"
	"github.com/rfdez/my-game-backend/internal/platform/storage/postgresql"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/simukti/sqldb-logger/logadapter/zerologadapter"
)

func main() {
	// Initialize the logger.
	logger := zerolog.NewLogger()

	// Load environment variables.
	var cfg config
	err := envconfig.Process("", &cfg)
	if err != nil {
		logger.Fatal("failed to process config")
	}

	// Initialize the database connection.
	psqlURI := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?%s", cfg.DbUser, cfg.DbPass, cfg.DbHost, cfg.DbPort, cfg.DbName, cfg.DbParams)
	db, err := sql.Open("postgres", psqlURI)
	if err != nil {
		logger.Fatal("failed to open database")
	}

	// Add a logger to the database connection.
	loggerAdapter := zerologadapter.New(logger.Logger())
	db = sqldblogger.OpenDriver(psqlURI, db.Driver(), loggerAdapter)

	// Bus
	var (
		commandBus = inmemory.NewCommandBus()
		queryBus   = inmemory.NewQueryBus()
		eventBus   = inmemory.NewEventBus(logger)
	)

	// Repositories
	var (
		checkRepository = postgresql.NewCheckRepository(db, cfg.DbTimeout)
		eventRepository = postgresql.NewEventRepository(db, cfg.DbTimeout)
	)

	// Services
	var (
		checkService       = checking.NewService(checkRepository)
		fetcherService     = fetcher.NewService(eventRepository, eventBus)
		incrementerService = incrementer.NewService(eventRepository)
	)

	// Command Handlers
	var (
		checkingCheckCommandHandler = checking.NewCheckCommandHandler(checkService)
	)

	// Register Command Handlers
	commandBus.Register(checking.CheckCommandType, checkingCheckCommandHandler)

	// Query Handlers
	var (
		fetcherRandomEventQueryHandler = fetcher.NewRandomEventQueryHandler(fetcherService)
	)

	// Register Query Handlers
	queryBus.Register(fetcher.RandomEventQueryType, fetcherRandomEventQueryHandler)

	// Event Subscribers
	eventBus.Subscribe(mygame.EventShownEventType, fetcher.NewIncreaseEventShownOnEventShown(incrementerService))

	ctx, srv := server.New(context.Background(), cfg.Host, cfg.Port, cfg.ShutdownTimeout, commandBus, queryBus, logger)
	if err := srv.Run(ctx); err != nil {
		logger.Fatal("failed to run server")
	}
}

type config struct {
	// Server configuration
	Host            string        `default:""`
	Port            uint          `default:"8080"`
	ShutdownTimeout time.Duration `default:"10s"`
	// Database configuration
	DbUser    string        `envconfig:"DB_USER" required:"true"`
	DbPass    string        `envconfig:"DB_PASS" required:"true"`
	DbHost    string        `envconfig:"DB_HOST" required:"true"`
	DbPort    uint          `envconfig:"DB_PORT" required:"true"`
	DbName    string        `envconfig:"DB_NAME" required:"true"`
	DbParams  string        `envconfig:"DB_PARAMS" default:""`
	DbTimeout time.Duration `default:"5s"`
}
