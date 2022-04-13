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

	defer rows.Close()

	var events []mygame.Event
	for rows.Next() {
		var event sqlEvent
		if err := rows.Scan(eventSQLStruct.Addr(&event)...); err != nil {
			return nil, errors.WrapInternal(err, "failed to scan event")
		}

		evt, err := mygame.NewEvent(
			event.ID,
			event.Name,
			event.Date.Format(time.RFC3339),
			event.Shown,
			event.Keywords,
		)
		if err != nil {
			return nil, errors.WrapInternal(err, "failed to create event")
		}

		events = append(events, evt)
	}

	return events, nil
}
