package fetcher

import (
	"context"

	"github.com/rfdez/my-game-backend/internal/errors"
	"github.com/rfdez/my-game-backend/kit/query"
)

const (
	RandomEventQueryType           query.Type = "command.fetcher.random_event"
	EventQuestionsByRoundQueryType query.Type = "command.fetcher.question_by_event_id"
)

// RandomEventQuery is a query to get a random event.
type RandomEventQuery struct {
	date string
}

// NewRandomEventQuery creates a new random event query.
func NewRandomEventQuery(date string) RandomEventQuery {
	return RandomEventQuery{
		date: date,
	}
}

// Type returns the type of the query.
func (q RandomEventQuery) Type() query.Type {
	return RandomEventQueryType
}

// Date returns the date of the random event.
func (q RandomEventQuery) Date() string {
	return q.date
}

// RandomEventQueryHandler is a query handler for RandomEventQuery.
type RandomEventQueryHandler struct {
	service Service
}

// NewRandomEventQueryHandler creates a new instance of RandomEventQueryHandler.
func NewRandomEventQueryHandler(service Service) RandomEventQueryHandler {
	return RandomEventQueryHandler{
		service: service,
	}
}

// Handle implements the query.Handler interface.
func (h RandomEventQueryHandler) Handle(ctx context.Context, q query.Query) (query.Response, error) {
	queryR, ok := q.(RandomEventQuery)
	if !ok {
		return nil, errors.New("invalid query type")
	}

	return h.service.RandomEvent(ctx, queryR.Date())
}

// EventQuestionsByRoundQuery is a query to get the questions by round.
type EventQuestionsByRoundQuery struct {
	eventID string
	round   int
}

// NewQuestionByEventIDQuery creates a new question by event id query.
func NewEventQuestionsByRoundQuery(eventID string, round int) EventQuestionsByRoundQuery {
	return EventQuestionsByRoundQuery{
		eventID: eventID,
		round:   round,
	}
}

func (q EventQuestionsByRoundQuery) Type() query.Type {
	return EventQuestionsByRoundQueryType
}

// EventID returns the event id.
func (q EventQuestionsByRoundQuery) EventID() string {
	return q.eventID
}

// Round is the round of the question.
func (q EventQuestionsByRoundQuery) Round() int {
	return q.round
}

// EventQuestionsByRoundQueryHandler is a query handler for EventQuestionsByRoundQuery.
type EventQuestionsByRoundQueryHandler struct {
	service Service
}

// NewEventQuestionsByRoundQueryHandler creates a new instance of QuestionByEventIDQueryHandler.
func NewEventQuestionsByRoundQueryHandler(service Service) EventQuestionsByRoundQueryHandler {
	return EventQuestionsByRoundQueryHandler{
		service: service,
	}
}

// Handle implements the query.Handler interface.
func (h EventQuestionsByRoundQueryHandler) Handle(ctx context.Context, q query.Query) (query.Response, error) {
	queryR, ok := q.(EventQuestionsByRoundQuery)
	if !ok {
		return nil, errors.New("invalid query type")
	}

	return h.service.EventQuestionsByRound(ctx, queryR.EventID(), queryR.Round())
}
