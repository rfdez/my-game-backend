package fetcher

import (
	"context"
	"math/rand"
	"time"

	mygame "github.com/rfdez/my-game-backend/internal"
	"github.com/rfdez/my-game-backend/internal/errors"
)

// Service is the interface that provides the fetcher service.
type Service interface {
	RandomEvent(context.Context, string) (RandomEventResponse, error)
	EventQuestionsByRound(context.Context, string, int) ([]EventQuestionResponse, error)
	Question(context.Context, string) (QuestionResponse, error)
	EventQuestionAnswer(context.Context, string, string) (AnswerResponse, error)
}

type service struct {
	eventRepository         mygame.EventRepository
	eventQuestionRepository mygame.EventQuestionRepository
	questionRepository      mygame.QuestionRepository
	answerRepository        mygame.AnswerRepository
}

// NewService creates a new instance of Service.
func NewService(eventRepository mygame.EventRepository, eventQuestionRepository mygame.EventQuestionRepository, questionRepository mygame.QuestionRepository, answerRepository mygame.AnswerRepository) Service {
	return &service{
		eventRepository:         eventRepository,
		eventQuestionRepository: eventQuestionRepository,
		questionRepository:      questionRepository,
		answerRepository:        answerRepository,
	}
}

// RandomEvent returns a random event.
func (s *service) RandomEvent(ctx context.Context, date string) (RandomEventResponse, error) {
	var eventDate time.Time
	if date == "" {
		eventDate = time.Now()
	} else {
		var err error
		eventDate, err = time.Parse(mygame.EventDateRFC3339, date)
		if err != nil {
			return RandomEventResponse{}, errors.NewWrongInput("invalid %s date, the format should be %s", date, mygame.EventDateRFC3339)
		}
	}

	events, err := s.eventRepository.SearchAll(ctx)
	if err != nil {
		return RandomEventResponse{}, err
	}

	if len(events) == 0 {
		return RandomEventResponse{}, errors.NewNotFound("no events found")
	}

	var eventsDate []mygame.Event
	for _, event := range events {
		if event.Date().Value().Day() == eventDate.Day() && event.Date().Value().Month() == eventDate.Month() {
			eventsDate = append(eventsDate, event)
		}
	}

	if len(eventsDate) == 0 {
		return RandomEventResponse{}, errors.NewNotFound("no events found for date %s", date)
	}

	randomIndex := rand.Intn(len(eventsDate))
	evt := eventsDate[randomIndex]

	return NewRandomEventResponse(evt.ID().String(), evt.Name().String(), evt.Date().String(), evt.Keywords().Value()), nil
}

// EventQuestionsByRound returns the questions by round.
func (s *service) EventQuestionsByRound(ctx context.Context, id string, round int) ([]EventQuestionResponse, error) {
	if round < 1 {
		return nil, errors.NewWrongInput("round must be greater than 0")
	}

	eventID, err := mygame.NewEventID(id)
	if err != nil {
		return nil, err
	}

	eventQuestions, err := s.eventQuestionRepository.SearchByEventID(ctx, eventID)
	if err != nil {
		return nil, err
	}

	if len(eventQuestions) == 0 {
		return nil, nil
	}

	var eventQuestionsByRound []mygame.EventQuestion
	for _, v := range eventQuestions {
		if v.Round().Value() == round {
			eventQuestionsByRound = append(eventQuestionsByRound, v)
		}
	}

	if len(eventQuestionsByRound) == 0 {
		return nil, nil
	}

	eventQuestionResponse := make([]EventQuestionResponse, len(eventQuestionsByRound))
	for i, v := range eventQuestionsByRound {
		eventQuestionResponse[i] = NewEventQuestionResponse(v.EventID().String(), v.QuestionID().String(), v.Round().Value())
	}

	return eventQuestionResponse, nil
}

// Question returns the question.
func (s *service) Question(ctx context.Context, questionID string) (QuestionResponse, error) {
	questionIDVO, err := mygame.NewQuestionID(questionID)
	if err != nil {
		return QuestionResponse{}, err
	}

	question, err := s.questionRepository.Find(ctx, questionIDVO)
	if err != nil {
		return QuestionResponse{}, err
	}

	return NewQuestionResponse(question.ID().String(), question.Text().String()), nil
}

// EventQuestionAnswer returns the answer.
func (s *service) EventQuestionAnswer(ctx context.Context, eventID, questionID string) (AnswerResponse, error) {
	eventIDVO, err := mygame.NewEventID(eventID)
	if err != nil {
		return AnswerResponse{}, err
	}

	questionIDVO, err := mygame.NewQuestionID(questionID)
	if err != nil {
		return AnswerResponse{}, err
	}

	answer, err := s.answerRepository.FindByEventIDAndQuestionID(ctx, eventIDVO, questionIDVO)
	if err != nil {
		return AnswerResponse{}, err
	}

	return NewAnswerResponse(answer.EventID().String(), answer.QuestionID().String(), answer.Text().String()), nil
}
