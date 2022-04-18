package fetcher

import (
	"context"

	"github.com/rfdez/my-game-backend/internal/errors"
	"github.com/rfdez/my-game-backend/kit/query"
)

const (
	RandomEventQueryType           query.Type = "command.fetcher.random_event"
	EventQuestionsByRoundQueryType query.Type = "command.fetcher.event_questions_by_round"
	QuestionQueryType              query.Type = "command.fetcher.question"
	EventQuestionAnswerQueryType   query.Type = "command.fetcher.event_question_answer"
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

// QuestionQuery is a query to get a question.
type QuestionQuery struct {
	questionID string
}

// NewQuestionQuery creates a new question query.
func NewQuestionQuery(questionID string) QuestionQuery {
	return QuestionQuery{
		questionID: questionID,
	}
}

func (q QuestionQuery) Type() query.Type {
	return QuestionQueryType
}

// QuestionID returns the question id.
func (q QuestionQuery) QuestionID() string {
	return q.questionID
}

// QuestionQueryHandler is a query handler for QuestionQuery.
type QuestionQueryHandler struct {
	service Service
}

// NewQuestionQueryHandler creates a new instance of QuestionQueryHandler.
func NewQuestionQueryHandler(service Service) QuestionQueryHandler {
	return QuestionQueryHandler{
		service: service,
	}
}

// Handle implements the query.Handler interface.
func (h QuestionQueryHandler) Handle(ctx context.Context, q query.Query) (query.Response, error) {
	queryR, ok := q.(QuestionQuery)
	if !ok {
		return nil, errors.New("invalid query type")
	}

	return h.service.Question(ctx, queryR.QuestionID())
}

// EventQuestionAnswerQuery is a query to answer a question.
type EventQuestionAnswerQuery struct {
	eventID    string
	questionID string
}

// NewQuestionAnswerQuery creates a new question answer query.
func NewEventQuestionAnswerQuery(eventID, questionID string) EventQuestionAnswerQuery {
	return EventQuestionAnswerQuery{
		eventID:    eventID,
		questionID: questionID,
	}
}

// Type returns the type of the query.
func (q EventQuestionAnswerQuery) Type() query.Type {
	return EventQuestionAnswerQueryType
}

// EventID returns the event id.
func (q EventQuestionAnswerQuery) EventID() string {
	return q.eventID
}

// QuestionID returns the question id.
func (q EventQuestionAnswerQuery) QuestionID() string {
	return q.questionID
}

// EventQuestionAnswerQueryHandler is a query handler for EventQuestionAnswerQuery.
type EventQuestionAnswerQueryHandler struct {
	service Service
}

// NewEventQuestionAnswerQueryHandler creates a new instance of EventQuestionAnswerQueryHandler.
func NewEventQuestionAnswerQueryHandler(service Service) EventQuestionAnswerQueryHandler {
	return EventQuestionAnswerQueryHandler{
		service: service,
	}
}

// Handle implements the query.Handler interface.
func (h EventQuestionAnswerQueryHandler) Handle(ctx context.Context, q query.Query) (query.Response, error) {
	queryR, ok := q.(EventQuestionAnswerQuery)
	if !ok {
		return nil, errors.New("invalid query type")
	}

	return h.service.EventQuestionAnswer(ctx, queryR.EventID(), queryR.QuestionID())
}
