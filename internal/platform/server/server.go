package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rfdez/my-game-backend/internal/platform/server/handler/checks"
	"github.com/rfdez/my-game-backend/internal/platform/server/handler/events"
	"github.com/rfdez/my-game-backend/internal/platform/server/middleware/logging"
	"github.com/rfdez/my-game-backend/internal/platform/server/middleware/recovery"
	"github.com/rfdez/my-game-backend/kit/command"
	"github.com/rfdez/my-game-backend/kit/logger"
	"github.com/rfdez/my-game-backend/kit/query"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	shutdownTimeout time.Duration

	// deps
	commandBus command.Bus
	queryBus   query.Bus
	logger     logger.Logger
}

func New(ctx context.Context, host string, port uint, shutdownTimeout time.Duration, commandBus command.Bus, queryBus query.Bus, logger logger.Logger) (context.Context, Server) {
	srv := Server{
		engine:   gin.New(),
		httpAddr: fmt.Sprintf("%s:%d", host, port),

		shutdownTimeout: shutdownTimeout,

		// deps
		commandBus: commandBus,
		queryBus:   queryBus,
		logger:     logger,
	}

	srv.registerRoutes()
	return serverContext(ctx), srv
}

func (s *Server) registerRoutes() {
	s.engine.Use(recovery.Middleware(), logging.Middleware(s.logger))

	// Health checks
	s.engine.GET("/ping", checks.PingHandler())
	s.engine.GET("/health", checks.HealthHandler(s.commandBus))

	// Events
	s.engine.GET("/events/random", events.RandomHandler(s.queryBus))
}

func (s *Server) Run(ctx context.Context) error {
	s.logger.Info(fmt.Sprintf("Server running on %s", s.httpAddr))

	srv := &http.Server{
		Addr:    s.httpAddr,
		Handler: s.engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Fatal(fmt.Sprintf("Server error: %s", err))
		}
	}()

	<-ctx.Done()
	ctxShutDown, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	s.logger.Info("Server shutting down...")

	return srv.Shutdown(ctxShutDown)
}

func serverContext(ctx context.Context) context.Context {
	c := make(chan os.Signal, 1)
	// interrupt signal sent from terminal
	signal.Notify(c, os.Interrupt)
	// sigterm signal sent from kubernetes
	signal.Notify(c, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		<-c
		cancel()
	}()

	return ctx
}
