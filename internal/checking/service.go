package checking

import (
	"context"

	mygame "github.com/rfdez/my-game-backend/internal"
)

type Service interface {
	Status(context.Context) error
}

type service struct {
	checkRepository mygame.CheckRepository
}

func NewService(checkRepository mygame.CheckRepository) Service {
	return &service{
		checkRepository: checkRepository,
	}
}

func (s *service) Status(ctx context.Context) error {
	return s.checkRepository.Status(ctx)
}
