package fetcher

import (
	"context"

	"github.com/rfdez/my-game-backend/internal/errors"
	"github.com/rfdez/my-game-backend/kit/query"
)

const (
	RandomEventQueryType query.Type = "command.fetcher.random_event"
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

// Type returns the type of the query.
func (q RandomEventQuery) Type() query.Type {
	return RandomEventQueryType
}

// Date returns the date of the random event.
func (q RandomEventQuery) Date() string {
	return q.date
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
	queryR, ok := q.(RandomEventQuery)
	if !ok {
		return nil, errors.New("invalid query type")
	}

	return h.service.RandomEvent(ctx, queryR.Date())
}
