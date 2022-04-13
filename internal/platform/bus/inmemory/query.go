package inmemory

import (
	"context"

	"github.com/rfdez/my-game-backend/kit/query"
)

// QueryBus is an in-memory implementation of the query.Bus.
type QueryBus struct {
	handlers map[query.Type]query.Handler
}

// NewQueryBus initializes a new instance of QueryBus.
func NewQueryBus() *QueryBus {
	return &QueryBus{
		handlers: make(map[query.Type]query.Handler),
	}
}

// Ask implements the query.Bus interface.
func (b *QueryBus) Ask(ctx context.Context, q query.Query) (query.Response, error) {
	handler, ok := b.handlers[q.Type()]
	if !ok {
		return nil, nil
	}

	return handler.Handle(ctx, q)
}

// Register implements the query.Bus interface.
func (b *QueryBus) Register(qType query.Type, handler query.Handler) {
	b.handlers[qType] = handler
}
