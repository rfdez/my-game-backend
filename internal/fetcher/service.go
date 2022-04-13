package fetcher

import (
	"context"
	"math/rand"

	mygame "github.com/rfdez/my-game-backend/internal"
	"github.com/rfdez/my-game-backend/internal/errors"
	"github.com/rfdez/my-game-backend/kit/event"
)

// Service is the interface that provides the fetcher service.
type Service interface {
	RandomEvent(ctx context.Context, date string) (RandomEventResponse, error)
}

type service struct {
	eventRepository mygame.EventRepository
	eventBus        event.Bus
}

// NewService creates a new instance of Service.
func NewService(eventRepository mygame.EventRepository, eventBus event.Bus) Service {
	return &service{
		eventRepository: eventRepository,
		eventBus:        eventBus,
	}
}

// RandomEvent returns a random event.
func (s *service) RandomEvent(ctx context.Context, date string) (RandomEventResponse, error) {
	events, err := s.eventRepository.SearchAll(ctx)
	if err != nil {
		return RandomEventResponse{}, err
	}

	if len(events) == 0 {
		return RandomEventResponse{}, errors.NewNotFound("no events found for date %s", date)
	}

	minShown := events[0].Shown().Value()
	for _, v := range events {
		if v.Shown().Value() < minShown {
			minShown = v.Shown().Value()
		}
	}

	var eventsWithMinShown []mygame.Event
	for _, v := range events {
		if v.Shown().Value() == minShown {
			eventsWithMinShown = append(eventsWithMinShown, v)
		}
	}

	randomIndex := rand.Intn(len(eventsWithMinShown))
	evt := eventsWithMinShown[randomIndex]

	newEvent := mygame.NewEventShownEvent(evt.ID().String())
	err = s.eventBus.Publish(ctx, []event.Event{newEvent})
	if err != nil {
		return RandomEventResponse{}, err
	}

	return NewRandomEventResponse(evt.ID().String(), evt.Name().String(), evt.Date().String(), evt.Keywords().Value()), nil
}
