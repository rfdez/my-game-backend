package checking

import (
	"context"

	mygame "github.com/rfdez/my-game-backend/internal"
)

// Service is the interface that provides the checking service.
type Service interface {
	Status(context.Context) error
}

type service struct {
	checkRepository mygame.CheckRepository
}

// NewService creates a new instance of Service.
func NewService(checkRepository mygame.CheckRepository) Service {
	return &service{
		checkRepository: checkRepository,
	}
}

// Status checks the status of the application.
func (s *service) Status(ctx context.Context) error {
	return s.checkRepository.Status(ctx)
}
