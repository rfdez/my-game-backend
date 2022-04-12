package fetcher

import (
	"context"

	"github.com/rfdez/my-game-backend/internal/errors"
	"github.com/rfdez/my-game-backend/kit/query"
)

const (
	RandomEventQueryType = "command.fetcher.random_event"
)

// RandomEventQuery is a query to get a random event.
type RandomEventQuery struct {
	date string
}

// NewRandomEventQuery creates a new random event query.
func NewRandomEventQuery(date string) RandomEventQuery {
	return RandomEventQuery{
		date: date,
	}
}

func (q RandomEventQuery) Type() query.Type {
	return RandomEventQueryType
}

// RandomEventQueryHandler is a query handler for RandomEventQuery.
type RandomEventQueryHandler struct {
	service Service
}

// NewRandomEventQueryHandler creates a new instance of RandomEventQueryHandler.
func NewRandomEventQueryHandler(service Service) RandomEventQueryHandler {
	return RandomEventQueryHandler{
		service: service,
	}
}

// Handle implements the query.Handler interface.
func (h RandomEventQueryHandler) Handle(ctx context.Context, q query.Query) (query.Response, error) {
	randomEventQuery, ok := q.(RandomEventQuery)
	if !ok {
		return nil, errors.NewWrongInput("unexpected query")
	}

	return h.service.RandomEvent(ctx, randomEventQuery.date)
}
