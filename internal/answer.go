package mygame

import (
	"context"

	"github.com/google/uuid"
	"github.com/rfdez/my-game-backend/internal/errors"
)

// AnswerID is the ID of the answer
type AnswerID struct {
	value string
}

// NewAnswerID instantiate the VO for AnswerID
func NewAnswerID(value string) (AnswerID, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return AnswerID{}, errors.NewWrongInput("invalid answer id %s", value)
	}

	return AnswerID{
		value: v.String(),
	}, nil
}

// String type converts the AnswerID into string.
func (id AnswerID) String() string {
	return id.value
}

// AnswerText is the text of the answer
type AnswerText struct {
	value string
}

// NewAnswerText instantiate the VO for AnswerText
func NewAnswerText(value string) (AnswerText, error) {
	if value == "" {
		return AnswerText{}, errors.NewWrongInput("the field Answer Text can not be empty")
	}

	return AnswerText{
		value: value,
	}, nil
}

// String type converts the AnswerText into string.
func (text AnswerText) String() string {
	return text.value
}

// Answer is the answer
type Answer struct {
	id         AnswerID
	text       AnswerText
	questionID QuestionID
}

// AnswerRepository is the interface for the answer repository
type AnswerRepository interface {
	FindByQuestionID(context.Context, QuestionID) (Answer, error)
}

// NewAnswer instantiate the VO for Answer
func NewAnswer(id, text, questionID string) (Answer, error) {
	idVO, err := NewAnswerID(id)
	if err != nil {
		return Answer{}, err
	}

	textVO, err := NewAnswerText(text)
	if err != nil {
		return Answer{}, err
	}

	questionIDVO, err := NewQuestionID(questionID)
	if err != nil {
		return Answer{}, err
	}

	return Answer{
		id:         idVO,
		text:       textVO,
		questionID: questionIDVO,
	}, nil
}

// ID type converts the Answer into string.
func (a Answer) ID() AnswerID {
	return a.id
}

// Text type converts the Answer into string.
func (a Answer) Text() AnswerText {
	return a.text
}

// QuestionID type converts the Answer into string.
func (a Answer) QuestionID() QuestionID {
	return a.questionID
}
