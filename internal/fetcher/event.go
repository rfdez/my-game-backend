package fetcher

import (
	"context"

	mygame "github.com/rfdez/my-game-backend/internal"
	"github.com/rfdez/my-game-backend/internal/errors"
	"github.com/rfdez/my-game-backend/internal/incrementer"
	"github.com/rfdez/my-game-backend/kit/event"
)

type IncreaseEventShownOnEventShown struct {
	incrementerService incrementer.Service
}

func NewIncreaseEventShownOnEventShown(incrementerService incrementer.Service) IncreaseEventShownOnEventShown {
	return IncreaseEventShownOnEventShown{
		incrementerService: incrementerService,
	}
}

func (e IncreaseEventShownOnEventShown) Handle(ctx context.Context, evt event.Event) error {
	eventShownEvt, ok := evt.(mygame.EventShownEvent)
	if !ok {
		return errors.New("unexpected event")
	}

	return e.incrementerService.IncreaseEventShown(ctx, eventShownEvt.ID())
}
