package mygame

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

var ErrInvalidEventID = errors.New("invalid Event ID")

// EventID represents the event unique identifier.
type EventID struct {
	value string
}

// NewEventID instantiate the VO for EventID
func NewEventID(value string) (EventID, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return EventID{}, fmt.Errorf("%w: %s", ErrInvalidEventID, value)
	}

	return EventID{
		value: v.String(),
	}, nil
}

// String type converts the EventID into string.
func (id EventID) String() string {
	return id.value
}

var ErrEmptyEventName = errors.New("the field Event Name can not be empty")

// EventName represents the event name.
type EventName struct {
	value string
}

// NewEventName instantiate VO for EventName
func NewEventName(value string) (EventName, error) {
	if value == "" {
		return EventName{}, ErrEmptyEventName
	}

	return EventName{
		value: value,
	}, nil
}

// String type converts the EventName into string.
func (name EventName) String() string {
	return name.value
}

// Event is the data structure that represents a event.
type Event struct {
	id   EventID
	name EventName
}

// EventRepository defines the expected behaviour from a course storage.
type EventRepository interface {
	SearchAll(ctx context.Context) ([]Event, error)
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=EventRepository

// NewEvent creates a new event.
func NewEvent(id, name string) (Event, error) {
	idVO, err := NewEventID(id)
	if err != nil {
		return Event{}, err
	}

	nameVO, err := NewEventName(name)
	if err != nil {
		return Event{}, err
	}

	event := Event{
		id:   idVO,
		name: nameVO,
	}
	return event, nil
}

// ID returns the event unique identifier.
func (e Event) ID() EventID {
	return e.id
}

// Name returns the event name.
func (e Event) Name() EventName {
	return e.name
}
