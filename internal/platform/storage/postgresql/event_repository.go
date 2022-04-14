package postgresql

import (
	"context"
	"database/sql"
	"time"

	"github.com/huandu/go-sqlbuilder"
	"github.com/lib/pq"
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
			event.Date.Format(mygame.RFC3339FullDate),
			event.Shown,
			event.Keywords,
		)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create event")
		}

		events = append(events, evt)
	}

	return events, nil
}

// Find implements the mygame.EventRepository repository.
func (r *eventRepository) Find(ctx context.Context, id mygame.EventID) (mygame.Event, error) {
	eventSQLStruct := sqlbuilder.NewStruct(new(sqlEvent))

	sb := eventSQLStruct.SelectFrom(sqlEventTable)
	sb.Where(sb.Equal("id", id.String()))

	query, args := sb.Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	var event sqlEvent
	if err := r.db.QueryRowContext(ctxTimeout, query, args...).Scan(eventSQLStruct.Addr(&event)...); err != nil {
		if err == sql.ErrNoRows {
			return mygame.Event{}, errors.NewNotFound("event %s not found", id.String())
		}

		return mygame.Event{}, errors.Wrap(err, "failed to search event")
	}

	evt, err := mygame.NewEvent(
		event.ID,
		event.Name,
		event.Date.Format(mygame.RFC3339FullDate),
		event.Shown,
		event.Keywords,
	)
	if err != nil {
		return mygame.Event{}, errors.Wrap(err, "failed to create event")
	}

	return evt, nil
}

// Update implements the mygame.EventRepository repository.
func (r *eventRepository) Update(ctx context.Context, event mygame.Event) error {
	eventSQLStruct := sqlbuilder.NewStruct(new(sqlEvent))
	ub := eventSQLStruct.Flavor.NewUpdateBuilder()
	ub.Update(sqlEventTable)
	ub.Set(
		ub.Assign("name", event.Name().String()),
		ub.Assign("date", event.Date().String()),
		ub.Assign("shown", event.Shown().Value()),
		ub.Assign("keywords", pq.Array(event.Keywords().Value())),
	)
	ub.Where(ub.Equal("id", event.ID().String()))

	query, args := ub.Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	if _, err := r.db.ExecContext(ctxTimeout, query, args...); err != nil {
		return errors.Wrap(err, "failed to update event")
	}

	return nil
}
