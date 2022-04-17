package fetcher

import (
	"context"
	"math/rand"
	"time"

	mygame "github.com/rfdez/my-game-backend/internal"
	"github.com/rfdez/my-game-backend/internal/errors"
	"github.com/rfdez/my-game-backend/kit/event"
)

// Service is the interface that provides the fetcher service.
type Service interface {
	RandomEvent(context.Context, string) (RandomEventResponse, error)
	EventQuestionsByRound(context.Context, string, int) (EventQuestionsByRoundResponse, error)
	QuestionAnswer(context.Context, string) (AnswerResponse, error)
}

type service struct {
	eventRepository    mygame.EventRepository
	questionRepository mygame.QuestionRepository
	answerRepository   mygame.AnswerRepository
	eventBus           event.Bus
}

// NewService creates a new instance of Service.
func NewService(eventRepository mygame.EventRepository, questionRepository mygame.QuestionRepository, answerRepository mygame.AnswerRepository, eventBus event.Bus) Service {
	return &service{
		eventRepository:    eventRepository,
		questionRepository: questionRepository,
		answerRepository:   answerRepository,
		eventBus:           eventBus,
	}
}

// RandomEvent returns a random event.
func (s *service) RandomEvent(ctx context.Context, date string) (RandomEventResponse, error) {
	if date == "" {
		date = time.Now().Format(mygame.EventDateRFC3339)
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
		if event.Date().String() == date {
			eventsDate = append(eventsDate, event)
		}
	}

	if len(eventsDate) == 0 {
		return RandomEventResponse{}, errors.NewNotFound("no events found for date %s", date)
	}

	minShown := eventsDate[0].Shown().Value()
	for _, v := range eventsDate {
		if v.Shown().Value() < minShown {
			minShown = v.Shown().Value()
		}
	}

	var eventsWithMinShown []mygame.Event
	for _, v := range eventsDate {
		if v.Shown().Value() == minShown {
			eventsWithMinShown = append(eventsWithMinShown, v)
		}
	}

	randomIndex := rand.Intn(len(eventsWithMinShown))
	evt := eventsWithMinShown[randomIndex]

	newEvent := mygame.NewEventShownEvent(evt.ID().String())
	err = s.eventBus.Publish(ctx, []event.Event{newEvent})
	if err != nil {
		return RandomEventResponse{}, err
	}

	return NewRandomEventResponse(evt.ID().String(), evt.Name().String(), evt.Date().String(), evt.Keywords().Value()), nil
}

// EventQuestionsByRound returns the questions by round.
func (s *service) EventQuestionsByRound(ctx context.Context, id string, round int) (EventQuestionsByRoundResponse, error) {
	eventID, err := mygame.NewEventID(id)
	if err != nil {
		return EventQuestionsByRoundResponse{}, err
	}

	questions, err := s.questionRepository.SearchByEventID(ctx, eventID)
	if err != nil {
		return EventQuestionsByRoundResponse{}, err
	}

	var questionsByRound []mygame.Question
	for _, v := range questions {
		if v.Round().Value() == round {
			questionsByRound = append(questionsByRound, v)
		}
	}

	if len(questionsByRound) == 0 {
		return EventQuestionsByRoundResponse{}, errors.NewNotFound("no questions found for event %s and round %d", id, round)
	}

	questionsResponse := make([]QuestionResponse, len(questionsByRound))
	for i, v := range questionsByRound {
		questionsResponse[i] = NewQuestionResponse(v.ID().String(), v.Text().String(), v.EventID().String())
	}

	return NewEventQuestionsByRoundResponse(questionsResponse), nil
}

// QuestionAnswer returns the answer.
func (s *service) QuestionAnswer(ctx context.Context, questionID string) (AnswerResponse, error) {
	questionIDVO, err := mygame.NewQuestionID(questionID)
	if err != nil {
		return AnswerResponse{}, err
	}

	answer, err := s.answerRepository.FindByQuestionID(ctx, questionIDVO)
	if err != nil {
		return AnswerResponse{}, err
	}

	return NewAnswerResponse(answer.ID().String(), answer.Text().String(), answer.QuestionID().String()), nil
}
