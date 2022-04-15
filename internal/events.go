package mygame

import "github.com/rfdez/my-game-backend/kit/event"

const EventShownEventType event.Type = "events.event.shown"

type EventShownEvent struct {
	event.BaseEvent
	id string
}

func NewEventShownEvent(id string) EventShownEvent {
	return EventShownEvent{
		id: id,

		BaseEvent: event.NewBaseEvent(id),
	}
}

func (e EventShownEvent) Type() event.Type {
	return EventShownEventType
}

func (e EventShownEvent) EventID() string {
	return e.id
}
