package mygame

import (
	"context"

	"github.com/google/uuid"
	"github.com/rfdez/my-game-backend/internal/errors"
)

// QuestionID represents the question unique identifier.
type QuestionID struct {
	value string
}

// NewQuestionID instantiate the VO for QuestionID
func NewQuestionID(value string) (QuestionID, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return QuestionID{}, errors.NewWrongInput("invalid question id %s", value)
	}

	return QuestionID{
		value: v.String(),
	}, nil
}

// String type converts the QuestionID into string.
func (id QuestionID) String() string {
	return id.value
}

// QuestionText represents the question text.
type QuestionText struct {
	value string
}

// NewQuestionText instantiate VO for QuestionText
func NewQuestionText(value string) (QuestionText, error) {
	if value == "" {
		return QuestionText{}, errors.NewWrongInput("the field Question Text can not be empty")
	}

	return QuestionText{
		value: value,
	}, nil
}

// String type converts the QuestionText into string.
func (text QuestionText) String() string {
	return text.value
}

// QuestionRound represents the question round.
type QuestionRound struct {
	value int
}

// NewQuestionRound instantiate the VO for the question round.
func NewQuestionRound(value int) (QuestionRound, error) {
	if value < 1 {
		return QuestionRound{}, errors.NewWrongInput("the field Question Round can not be less than 1")
	}

	return QuestionRound{
		value: value,
	}, nil
}

// Value returns the question round value.
func (v QuestionRound) Value() int {
	return v.value
}

// Question is the domain object for the question.
type Question struct {
	id      QuestionID
	text    QuestionText
	round   QuestionRound
	eventID EventID
}

// QuestionRepository is the interface for the question repository.
type QuestionRepository interface {
	SearchByEventID(context.Context, EventID) ([]Question, error)
}

// NewQuestion instantiate the entity for the question.
func NewQuestion(id, text string, round int, eventID string) (Question, error) {
	idVO, err := NewQuestionID(id)
	if err != nil {
		return Question{}, err
	}

	textVO, err := NewQuestionText(text)
	if err != nil {
		return Question{}, err
	}

	roundVO, err := NewQuestionRound(round)
	if err != nil {
		return Question{}, err
	}

	eventIDVO, err := NewEventID(eventID)
	if err != nil {
		return Question{}, err
	}

	return Question{
		id:      idVO,
		text:    textVO,
		round:   roundVO,
		eventID: eventIDVO,
	}, nil
}

// ID returns the question unique identifier.
func (v Question) ID() QuestionID {
	return v.id
}

// Text returns the question text.
func (v Question) Text() QuestionText {
	return v.text
}

// Round returns the question round.
func (v Question) Round() QuestionRound {
	return v.round
}

// EventID returns the event unique identifier.
func (v Question) EventID() EventID {
	return v.eventID
}
