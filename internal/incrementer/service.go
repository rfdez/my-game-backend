package incrementer

import (
	"context"

	mygame "github.com/rfdez/my-game-backend/internal"
)

// Service is the interface that provides the checking service.
type Service interface {
	IncreaseEventShown(context.Context, string) error
}

type service struct {
	eventRepository mygame.EventRepository
}

// NewService creates a new instance of Service.
func NewService(eventRepository mygame.EventRepository) Service {
	return &service{
		eventRepository: eventRepository,
	}
}

// IncreaseEventShown increases the number of times an event has been shown.
func (s *service) IncreaseEventShown(ctx context.Context, id string) error {
	evtID, err := mygame.NewEventID(id)
	if err != nil {
		return err
	}

	event, err := s.eventRepository.Find(ctx, evtID)
	if err != nil {
		return err
	}

	newEvent, err := mygame.NewEvent(
		event.ID().String(),
		event.Name().String(),
		event.Date().String(),
		event.Shown().Value()+1,
		event.Keywords().Value(),
	)
	if err != nil {
		return err
	}

	return s.eventRepository.Update(ctx, newEvent)
}
