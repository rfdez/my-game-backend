package server

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rfdez/my-game-backend/internal/platform/server/handler/checks"
	"github.com/rfdez/my-game-backend/internal/platform/server/handler/events"
	"github.com/rfdez/my-game-backend/kit/command"
	"github.com/rfdez/my-game-backend/kit/query"
	"github.com/rs/zerolog"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	shutdownTimeout time.Duration

	// deps
	commandBus command.Bus
	queryBus   query.Bus
}

func New(ctx context.Context, host string, port uint, shutdownTimeout time.Duration, commandBus command.Bus, queryBus query.Bus) (context.Context, Server) {
	srv := Server{
		engine:   gin.New(),
		httpAddr: fmt.Sprintf("%s:%d", host, port),

		shutdownTimeout: shutdownTimeout,

		// deps
		commandBus: commandBus,
		queryBus:   queryBus,
	}

	srv.registerRoutes()
	return serverContext(ctx), srv
}

func (s *Server) registerRoutes() {
	s.engine.Use(gin.Recovery())
	s.engine.Use(logger.SetLogger(
		logger.WithLogger(func(c *gin.Context, out io.Writer, latency time.Duration) zerolog.Logger {
			consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}
			return zerolog.New(consoleWriter).With().
				Timestamp().
				Int("statusCode", c.Writer.Status()).
				Dur("latency", latency).
				Str("clientIP", c.ClientIP()).
				Str("method", c.Request.Method).
				Str("path", c.Request.URL.Path).
				Logger()
		}),
	))

	// Health checks
	s.engine.GET("/ping", checks.PingHandler())
	s.engine.GET("/health", checks.HealthHandler(s.commandBus))

	// Events
	s.engine.GET("/events/random", events.RandomHandler(s.queryBus))
}

func (s *Server) Run(ctx context.Context) error {
	log.Printf("Server running on %s", s.httpAddr)

	srv := &http.Server{
		Addr:    s.httpAddr,
		Handler: s.engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()
	ctxShutDown, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	log.Println("Server shutting down...")

	return srv.Shutdown(ctxShutDown)
}

func serverContext(ctx context.Context) context.Context {
	c := make(chan os.Signal, 1)
	// interrupt signal sent from terminal
	signal.Notify(c, os.Interrupt)
	// sigterm signal sent from kubernetes
	signal.Notify(c, syscall.SIGTERM)
	// context that is canceled when interrupt signal is sent
	signal.Notify(c, syscall.SIGINT)
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		<-c
		cancel()
	}()

	return ctx
}
