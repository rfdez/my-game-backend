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

	sb := eventSQLStruct.SelectFrom(sqlEventTable)

	query, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	rows, err := r.db.QueryContext(ctxTimeout, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to search events")
	}
	defer rows.Close()

	var events []mygame.Event
	for rows.Next() {
		var event sqlEvent
		if err := rows.Scan(eventSQLStruct.Addr(&event)...); err != nil {
			return nil, errors.Wrap(err, "failed to scan event")
		}

		evt, err := mygame.NewEvent(
			event.ID,
			event.Name,
			event.Date.Format(mygame.EventDateRFC3339),
			event.Keywords,
		)
		if err != nil {
			return nil, err
		}

		events = append(events, evt)
	}

	return events, nil
}
