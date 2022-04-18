package mygame

import (
	"context"

	"github.com/rfdez/my-game-backend/internal/errors"
)

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
	eventID    EventID
	questionID QuestionID
	text       AnswerText
}

// AnswerRepository is the interface for the answer repository
type AnswerRepository interface {
	FindByEventIDAndQuestionID(context.Context, EventID, QuestionID) (Answer, error)
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=AnswerRepository

// NewAnswer instantiate the VO for Answer
func NewAnswer(eventID, questionID, text string) (Answer, error) {
	eventIDVO, err := NewEventID(eventID)
	if err != nil {
		return Answer{}, err
	}

	questionIDVO, err := NewQuestionID(questionID)
	if err != nil {
		return Answer{}, err
	}

	textVO, err := NewAnswerText(text)
	if err != nil {
		return Answer{}, err
	}

	return Answer{
		eventID:    eventIDVO,
		questionID: questionIDVO,
		text:       textVO,
	}, nil
}

// EventID type converts the Answer into string.
func (a Answer) EventID() EventID {
	return a.eventID
}

// QuestionID type converts the Answer into string.
func (a Answer) QuestionID() QuestionID {
	return a.questionID
}

// Text type converts the Answer into string.
func (a Answer) Text() AnswerText {
	return a.text
}
