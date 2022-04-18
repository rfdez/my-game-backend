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

// Question is the domain object for the question.
type Question struct {
	id   QuestionID
	text QuestionText
}

// QuestionRepository is the interface for the question repository.
type QuestionRepository interface {
	Find(context.Context, QuestionID) (Question, error)
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=QuestionRepository

// NewQuestion instantiate the entity for the question.
func NewQuestion(id, text string) (Question, error) {
	idVO, err := NewQuestionID(id)
	if err != nil {
		return Question{}, err
	}

	textVO, err := NewQuestionText(text)
	if err != nil {
		return Question{}, err
	}

	return Question{
		id:   idVO,
		text: textVO,
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
