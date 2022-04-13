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

// Find implements the mygame.EventRepository repository.
func (r *eventRepository) Find(ctx context.Context, id mygame.EventID) (mygame.Event, error) {
	eventSQLStruct := sqlbuilder.NewStruct(new(sqlEvent))

	query, args := eventSQLStruct.SelectFrom(sqlEventTable).Where(sqlEvent.ID.Eq(id.String())).Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)

	defer cancel()

	var event sqlEvent

	if err := r.db.QueryRowContext(ctxTimeout, query, args...).Scan(eventSQLStruct.Addr(&event)...); err != nil {
		if err == sql.ErrNoRows {
			return mygame.Event{}, errors.NewNotFound("event %s not found", id.String())
		}

		return mygame.Event{}, errors.WrapInternal(err, "failed to search event")
	}

	evt, err := mygame.NewEvent(
		event.ID,
		event.Name,
		event.Date.Format(time.RFC3339),
		event.Shown,
		event.Keywords,
	)
	if err != nil {
		return mygame.Event{}, errors.WrapInternal(err, "failed to create event")
	}

	return evt, nil
}

// Update implements the mygame.EventRepository repository.
func (r *eventRepository) Update(ctx context.Context, event mygame.Event) error {
	eventSQLStruct := sqlbuilder.NewStruct(new(sqlEvent))

	query, args := eventSQLStruct.Update(sqlEventTable).
		Set(sqlEvent.Name, event.Name).
		Set(sqlEvent.Date, event.Date).
		Set(sqlEvent.Shown, event.Shown).
		Set(sqlEvent.Keywords, event.Keywords).
		Where(sqlEvent.ID.Eq(event.ID.String())).
		Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	if _, err := r.db.ExecContext(ctxTimeout, query, args...); err != nil {
		return errors.WrapInternal(err, "failed to update event")
	}

	return nil
}
