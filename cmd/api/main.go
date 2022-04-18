package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"github.com/rfdez/my-game-backend/internal/checking"
	"github.com/rfdez/my-game-backend/internal/fetcher"
	"github.com/rfdez/my-game-backend/internal/platform/bus/inmemory"
	"github.com/rfdez/my-game-backend/internal/platform/server"
	"github.com/rfdez/my-game-backend/internal/platform/storage/postgresql"
	"github.com/rs/zerolog"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/simukti/sqldb-logger/logadapter/zerologadapter"
)

func main() {
	// Load environment variables.
	var cfg config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the database connection.
	psqlURI := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?%s", cfg.DbUser, cfg.DbPass, cfg.DbHost, cfg.DbPort, cfg.DbName, cfg.DbParams)
	db, err := sql.Open("postgres", psqlURI)
	if err != nil {
		log.Fatal(err)
	}

	// Add a logger to the database connection.
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}

	loggerAdapter := zerologadapter.New(zerolog.New(consoleWriter).With().Timestamp().Logger())
	db = sqldblogger.OpenDriver(psqlURI, db.Driver(), loggerAdapter, sqldblogger.WithMinimumLevel(sqldblogger.LevelInfo))

	// Initialize the logger.
	// logger := zerologlogger.NewLogger()

	// Bus
	var (
		commandBus = inmemory.NewCommandBus()
		queryBus   = inmemory.NewQueryBus()
	)

	// Repositories
	var (
		checkRepository         = postgresql.NewCheckRepository(db, cfg.DbTimeout)
		eventRepository         = postgresql.NewEventRepository(db, cfg.DbTimeout)
		eventQuestionRepository = postgresql.NewEventQuestionRepository(db, cfg.DbTimeout)
		questionRepository      = postgresql.NewQuestionRepository(db, cfg.DbTimeout)
		answerRepository        = postgresql.NewAnswerRepository(db, cfg.DbTimeout)
	)

	// Services
	var (
		checkService   = checking.NewService(checkRepository)
		fetcherService = fetcher.NewService(eventRepository, eventQuestionRepository, questionRepository, answerRepository)
	)

	// Command Handlers
	var (
		checkingCheckCommandHandler = checking.NewCheckCommandHandler(checkService)
	)

	// Register Command Handlers
	commandBus.Register(checking.CheckCommandType, checkingCheckCommandHandler)

	// Query Handlers
	var (
		fetcherRandomEventQueryHandler           = fetcher.NewRandomEventQueryHandler(fetcherService)
		fetcherEventQuestionsByRoundQueryHandler = fetcher.NewEventQuestionsByRoundQueryHandler(fetcherService)
		fetcherQuestionQueryHandler              = fetcher.NewQuestionQueryHandler(fetcherService)
		fetcherEventQuestionAnswerQueryHandler   = fetcher.NewEventQuestionAnswerQueryHandler(fetcherService)
	)

	// Register Query Handlers
	queryBus.Register(fetcher.RandomEventQueryType, fetcherRandomEventQueryHandler)
	queryBus.Register(fetcher.EventQuestionsByRoundQueryType, fetcherEventQuestionsByRoundQueryHandler)
	queryBus.Register(fetcher.QuestionQueryType, fetcherQuestionQueryHandler)
	queryBus.Register(fetcher.EventQuestionAnswerQueryType, fetcherEventQuestionAnswerQueryHandler)

	ctx, srv := server.New(context.Background(), cfg.Host, cfg.Port, cfg.ShutdownTimeout, commandBus, queryBus)
	if err := srv.Run(ctx); err != nil {
		log.Fatal(err)
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
