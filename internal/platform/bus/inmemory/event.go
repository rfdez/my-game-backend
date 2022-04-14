package inmemory

import (
	"context"

	"github.com/rfdez/my-game-backend/kit/event"
	"github.com/rfdez/my-game-backend/kit/logger"
)

// EventBus is an in-memory implementation of the event.Bus.
type EventBus struct {
	handlers map[event.Type][]event.Handler
	logger   logger.Logger
}

// NewEventBus initializes a new EventBus.
func NewEventBus(logger logger.Logger) *EventBus {
	return &EventBus{
		handlers: make(map[event.Type][]event.Handler),
		logger:   logger,
	}
}

// Publish implements the event.Bus interface.
func (b *EventBus) Publish(ctx context.Context, events []event.Event) error {
	for _, evt := range events {
		handlers, ok := b.handlers[evt.Type()]
		if !ok {
			return nil
		}

		b.logger.Debug("Publishing event")

		for _, handler := range handlers {
			err := handler.Handle(ctx, evt)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// Subscribe implements the event.Bus interface.
func (b *EventBus) Subscribe(evtType event.Type, handler event.Handler) {
	subscribersForType, ok := b.handlers[evtType]
	if !ok {
		b.handlers[evtType] = []event.Handler{handler}
	}

	subscribersForType = append(subscribersForType, handler)

	b.handlers[evtType] = subscribersForType
}
