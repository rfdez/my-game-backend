package postgresql

import (
	"context"
	"database/sql"
	"time"

	"github.com/huandu/go-sqlbuilder"
	mygame "github.com/rfdez/my-game-backend/internal"
	"github.com/rfdez/my-game-backend/internal/errors"
)

type eventRepository struct {
	db        *sql.DB
	dbTimeout time.Duration
}

// NewEventRepository initializes a PostgreSQL implementation of mygame.EventRepository.
func NewEventRepository(db *sql.DB, dbTimeout time.Duration) *eventRepository {
	return &eventRepository{
		db:        db,
		dbTimeout: dbTimeout,
	}
}

// SearchAll implements the mygame.EventRepository repository.
func (r *eventRepository) SearchAll(ctx context.Context) ([]mygame.Event, error) {
	eventSQLStruct := sqlbuilder.NewStruct(new(sqlEvent))

	query, args := eventSQLStruct.SelectFrom(sqlEventTable).Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	rows, err := r.db.QueryContext(ctxTimeout, query, args...)
	if err != nil {
		return nil, errors.WrapInternal(err, "failed to search events")
	}

	var events []sqlEvent
	rows.Scan(eventSQLStruct.Addr(&events)...)

	var eventsDomain = make([]mygame.Event, len(events))
	for i, event := range events {
		evt, err := mygame.NewEvent(
			event.ID,
			event.Name,
		)
		if err != nil {
			return nil, errors.WrapInternal(err, "failed to create event")
		}

		eventsDomain[i] = evt
	}

	return eventsDomain, nil
}
