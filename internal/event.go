package mygame

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/rfdez/my-game-backend/internal/errors"
)

var ErrInvalidEventID = errors.NewWrongInput("invalid event id")

// EventID represents the event unique identifier.
type EventID struct {
	value string
}

// NewEventID instantiate the VO for EventID
func NewEventID(value string) (EventID, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return EventID{}, errors.NewWrongInput("invalid event id %s", value)
	}

	return EventID{
		value: v.String(),
	}, nil
}

// String type converts the EventID into string.
func (id EventID) String() string {
	return id.value
}

// EventName represents the event name.
type EventName struct {
	value string
}

// NewEventName instantiate VO for EventName
func NewEventName(value string) (EventName, error) {
	if value == "" {
		return EventName{}, errors.NewWrongInput("the field Event Name can not be empty")
	}

	return EventName{
		value: value,
	}, nil
}

// String type converts the EventName into string.
func (name EventName) String() string {
	return name.value
}

const RFC3339FullDate = "2006-01-02"

// EventDate represents the event date.
type EventDate struct {
	value string
}

// NewEventDate instantiate VO for EventDate
func NewEventDate(value string) (EventDate, error) {
	date, err := time.Parse(RFC3339FullDate, value)
	if err != nil {
		return EventDate{}, errors.NewWrongInput("invalid event date %s", value)
	}

	return EventDate{
		value: date.Format(RFC3339FullDate),
	}, nil
}

// String type converts the EventDate into string.
func (date EventDate) String() string {
	return date.value
}

// EventShown represents the event shown.
type EventShown struct {
	value int
}

// NewEventShown instantiate VO for EventShown
func NewEventShown(value int) EventShown {
	if value <= 0 {
		return EventShown{
			value: 0,
		}
	}

	return EventShown{
		value: value,
	}
}

// Value returns the event shown value.
func (e EventShown) Value() int {
	return e.value
}

// EventKeywords returns the event keywords.
type EventKeywords struct {
	value []string
}

// NewEventKeywords instantiate VO for EventKeywords
func NewEventKeywords(value []string) (EventKeywords, error) {
	if len(value) == 0 {
		return EventKeywords{}, errors.NewWrongInput("the field Event Keywords must have at least one value")
	}

	return EventKeywords{
		value: value,
	}, nil
}

// Value returns the event keywords value.
func (e EventKeywords) Value() []string {
	return e.value
}

// Event is the data structure that represents a event.
type Event struct {
	id       EventID
	name     EventName
	date     EventDate
	shown    EventShown
	keywords EventKeywords
}

// EventRepository defines the expected behaviour from a course storage.
type EventRepository interface {
	SearchAll(ctx context.Context) ([]Event, error)
	Find(ctx context.Context, id EventID) (Event, error)
	Update(ctx context.Context, event Event) error
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=EventRepository

// NewEvent creates a new event.
func NewEvent(id, name, date string, shown int, keywords []string) (Event, error) {
	idVO, err := NewEventID(id)
	if err != nil {
		return Event{}, err
	}

	nameVO, err := NewEventName(name)
	if err != nil {
		return Event{}, err
	}

	dateVO, err := NewEventDate(date)
	if err != nil {
		return Event{}, err
	}

	shownVO := NewEventShown(shown)

	keywordsVO, err := NewEventKeywords(keywords)
	if err != nil {
		return Event{}, err
	}

	event := Event{
		id:       idVO,
		name:     nameVO,
		date:     dateVO,
		shown:    shownVO,
		keywords: keywordsVO,
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

// Date returns the event date.
func (e Event) Date() EventDate {
	return e.date
}

// Shown returns the event shown.
func (e Event) Shown() EventShown {
	return e.shown
}

// Keywords returns the event keywords.
func (e Event) Keywords() EventKeywords {
	return e.keywords
}
