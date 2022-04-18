package mygame

import (
	"context"

	"github.com/rfdez/my-game-backend/internal/errors"
)

// EventQuestionRound represents the event for the question round.
type EventQuestionRound struct {
	round int
}

// NewEventQuestionRound instantiate the event for the question round.
func NewEventQuestionRound(round int) (EventQuestionRound, error) {
	if round <= 0 {
		return EventQuestionRound{}, errors.NewWrongInput("the field Round can not be less than 1")
	}

	return EventQuestionRound{
		round: round,
	}, nil
}

// Value returns the question round value.
func (v EventQuestionRound) Value() int {
	return v.round
}

// EventQuestion represents the event for the question.
type EventQuestion struct {
	eventID    EventID
	questionID QuestionID
	round      EventQuestionRound
}

// EventQuestionRepository is the repository for the event question.
type EventQuestionRepository interface {
	SearchByEventID(context.Context, EventID) ([]EventQuestion, error)
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=EventQuestionRepository

// NewEventQuestion creates a new event question.
func NewEventQuestion(eventID, questionID string, round int) (EventQuestion, error) {
	eventIDVO, err := NewEventID(eventID)
	if err != nil {
		return EventQuestion{}, err
	}

	questionIDVO, err := NewQuestionID(questionID)
	if err != nil {
		return EventQuestion{}, err
	}

	roundVO, err := NewEventQuestionRound(round)
	if err != nil {
		return EventQuestion{}, err
	}

	return EventQuestion{
		eventID:    eventIDVO,
		questionID: questionIDVO,
		round:      roundVO,
	}, nil
}

// EventID returns the event unique identifier.
func (v EventQuestion) EventID() EventID {
	return v.eventID
}

// QuestionID returns the question unique identifier.
func (v EventQuestion) QuestionID() QuestionID {
	return v.questionID
}

// Round returns the question round.
func (v EventQuestion) Round() EventQuestionRound {
	return v.round
}
