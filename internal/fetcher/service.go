package fetcher

import (
	"context"
	"log"

	mygame "github.com/rfdez/my-game-backend/internal"
	"github.com/rfdez/my-game-backend/kit/query"
)

type Service interface {
	RandomEvent(ctx context.Context, date string) (query.Response, error)
}

type service struct {
	eventRepository mygame.EventRepository
}

func NewService(eventRepository mygame.EventRepository) Service {
	return &service{
		eventRepository: eventRepository,
	}
}

func (s *service) RandomEvent(ctx context.Context, date string) (query.Response, error) {
	log.Println(date)
	event, err := s.eventRepository.SearchAll(ctx)
	if err != nil {
		return RandomEventResponse{}, err
	}

	return NewRandomEventResponse(event[0].ID().String(), event[0].Name().String()), nil
}
