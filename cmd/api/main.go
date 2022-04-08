package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"github.com/rfdez/my-game-backend/internal/checking"
	"github.com/rfdez/my-game-backend/internal/errors"
	"github.com/rfdez/my-game-backend/internal/platform/bus/inmemory"
	"github.com/rfdez/my-game-backend/internal/platform/server"
	"github.com/rfdez/my-game-backend/internal/platform/storage/postgresql"
)

func main() {
	var cfg config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(errors.WrapInternal(err, "failed to process config"))
	}

	psqlURI := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?%s", cfg.DbUser, cfg.DbPass, cfg.DbHost, cfg.DbPort, cfg.DbName, cfg.DbParams)
	db, err := sql.Open("postgres", psqlURI)
	if err != nil {
		log.Fatal(errors.WrapInternal(err, "failed to open database"))
	}

	// Bus
	var (
		commandBus = inmemory.NewCommandBus()
	)

	// Repositories
	var (
		checkRepository = postgresql.NewCheckRepository(db, cfg.DbTimeout)
	)

	// Services
	var (
		checkService = checking.NewService(checkRepository)
	)

	// Command Handlers
	var (
		checkingCheckCommandHandler = checking.NewCheckCommandHandler(checkService)
	)

	// Register Command Handlers
	commandBus.Register(checking.CheckCommandType, checkingCheckCommandHandler)

	ctx, srv := server.New(context.Background(), cfg.Host, cfg.Port, cfg.ShutdownTimeout, commandBus)
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
